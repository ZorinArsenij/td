package proto

import (
	"fmt"
	"sync"
	"time"
)

// Message identifiers are coupled to message creation time.
//
// https://core.telegram.org/mtproto/description#message-identifier-msg-id

const (
	yieldClient         = 0
	yieldServerResponse = 1
	yieldFromServer     = 3

	messageIDModulo = 4
)

func newMessageID(now time.Time, yield int) int64 {
	const nano = 1e9
	// Must approximately equal unixtime*2^32.
	nowNano := now.UnixNano()

	// Important: to counter replay-attacks the lower 32 bits of msg_id
	// passed by the client must not be empty and must present a
	// fractional part of the time point when the message was created.
	intPart := nowNano / nano
	fracPart := nowNano % nano

	// Ensure that fracPart % 4 == 0.
	fracPart &= -messageIDModulo
	// Adding modulo 4 yield to ensure message type.
	fracPart += int64(yield)

	return (intPart << 32) | fracPart
}

// MessageID represents 64-bit message id.
type MessageID int64

func (id MessageID) String() string {
	return fmt.Sprintf("%x (%s, %s)",
		int64(id), id.Type(), id.Time().Format(time.RFC3339),
	)
}

// MessageType is type of message determined by message id.
//
// A message is rejected over 300 seconds after it is created or
// 30 seconds before it is created (this is needed to protect from replay attacks).
//
// The identifier of a message container must be strictly greater than those of
// its nested messages.
type MessageType byte

const (
	// MessageUnknown reports that message id has unknown time and probably
	// should be ignored.
	MessageUnknown MessageType = iota
	// MessageFromClient is client message identifiers.
	MessageFromClient
	// MessageServerResponse is a response to a client message.
	MessageServerResponse
	// MessageFromServer is a message from the server.
	MessageFromServer
)

func (m MessageType) String() string {
	switch m {
	case MessageFromClient:
		return "FromClient"
	case MessageServerResponse:
		return "ServerResponse"
	case MessageFromServer:
		return "FromServer"
	default:
		return "Unknown"
	}
}

// Time returns approximate time when MessageID were generated.
func (id MessageID) Time() time.Time {
	intPart := int64(id) >> 32
	fracPart := int64(int32(id))
	return time.Unix(intPart, fracPart).UTC()
}

// Type returns message type.
func (id MessageID) Type() MessageType {
	switch id % messageIDModulo {
	case yieldClient:
		return MessageFromClient
	case yieldServerResponse:
		return MessageServerResponse
	case yieldFromServer:
		return MessageFromServer
	default:
		return MessageUnknown
	}
}

// NewMessageID returns new message id for provided time and type.
func NewMessageID(now time.Time, typ MessageType) MessageID {
	var yield int
	switch typ {
	case MessageFromClient:
		yield = yieldClient
	case MessageFromServer:
		yield = yieldFromServer
	case MessageServerResponse:
		yield = yieldServerResponse
	default:
		yield = yieldClient
	}
	return MessageID(newMessageID(now, yield))
}

// MessageIDGen is message id generator that provides collision prevention.
//
// The main reason of such structure is that now() can return same time during
// multiple calls and that leads to duplicate message id.
type MessageIDGen struct {
	// mux protects buf and cursor, which is ring buffer for
	// recent generated message ids.
	mux    sync.Mutex
	buf    []int64
	cursor int

	now func() time.Time
}

// contains returns true if id was generated previously or false
// if (probably) not generated.
//
// Assumes mux locked.
func (g *MessageIDGen) contains(id int64) bool {
	for i := range g.buf {
		if g.buf[i] == id {
			return true
		}
	}
	return false
}

// add saves id as already generated.
//
// Assumes mux locked.
func (g *MessageIDGen) add(id int64) {
	if g.cursor >= len(g.buf) {
		g.cursor = 0
	}
	g.buf[g.cursor] = id
	g.cursor++
}

// New generates new message id for provided type, protecting from collisions
// that are caused by low system time resolution.
func (g *MessageIDGen) New(t MessageType) int64 {
	g.mux.Lock()
	defer g.mux.Unlock()

	now := g.now()
	id := int64(NewMessageID(now, t))

	const resolution = time.Nanosecond * 5

	for g.contains(id) {
		now = now.Add(resolution)
		id = int64(NewMessageID(now, t))
	}

	g.add(id)

	return id
}

// NewMessageIDGen creates new message id generator.
//
// Current time will be provided by now() function.
// The n parameter configures capacity of message id history, bigger n results
// in bigger memory consumption.
func NewMessageIDGen(now func() time.Time, n int) *MessageIDGen {
	return &MessageIDGen{
		buf: make([]int64, n),
		now: now,
	}
}

// MessageIDBuf stores last N message ids and is used in replay attack mitigation.
type MessageIDBuf struct {
	mux sync.Mutex
	buf []int64
}

// NewMessageIDBuf initializes new message id buffer for last N stored values.
func NewMessageIDBuf(n int) *MessageIDBuf {
	return &MessageIDBuf{
		buf: make([]int64, n),
	}
}

// Consume returns false if message should be discarded.
func (b *MessageIDBuf) Consume(newID int64) bool {
	// In addition, the identifiers (msg_id) of the last N messages received
	// from the other side must be stored, and if a message comes in with an
	// msg_id lower than all or equal to any of the stored values, that message
	// is to be ignored. Otherwise, the new message msg_id is added to the set,
	// and, if the number of stored msg_id values is greater than N, the oldest
	// (i. e. the lowest) is discarded.
	//
	// https://core.telegram.org/mtproto/security_guidelines#checking-msg-id

	b.mux.Lock()
	defer b.mux.Unlock()

	var (
		minIDx int
		minID  int64
	)
	for i, id := range b.buf {
		if id == newID {
			// Equal to stored value.
			return false
		}
		// Searching for minimum value.
		if id < minID {
			minIDx = i
			minID = id
		}
	}
	if newID < minID {
		// Lower than all stored values.
		return false
	}

	// Message is accepted. Replacing lowest message id with new id.
	b.buf[minIDx] = newID
	return true
}
