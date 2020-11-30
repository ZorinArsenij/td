// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// PollAnswerVoters represents TL type `pollAnswerVoters#3b6ddad2`.
type PollAnswerVoters struct {
	// Flags field of PollAnswerVoters.
	Flags bin.Fields
	// Chosen field of PollAnswerVoters.
	Chosen bool
	// Correct field of PollAnswerVoters.
	Correct bool
	// Option field of PollAnswerVoters.
	Option []byte
	// Voters field of PollAnswerVoters.
	Voters int
}

// PollAnswerVotersTypeID is TL type id of PollAnswerVoters.
const PollAnswerVotersTypeID = 0x3b6ddad2

// Encode implements bin.Encoder.
func (p *PollAnswerVoters) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pollAnswerVoters#3b6ddad2 as nil")
	}
	b.PutID(PollAnswerVotersTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pollAnswerVoters#3b6ddad2: field flags: %w", err)
	}
	b.PutBytes(p.Option)
	b.PutInt(p.Voters)
	return nil
}

// SetChosen sets value of Chosen conditional field.
func (p *PollAnswerVoters) SetChosen(value bool) {
	if value {
		p.Flags.Set(0)
	} else {
		p.Flags.Unset(0)
	}
}

// SetCorrect sets value of Correct conditional field.
func (p *PollAnswerVoters) SetCorrect(value bool) {
	if value {
		p.Flags.Set(1)
	} else {
		p.Flags.Unset(1)
	}
}

// Decode implements bin.Decoder.
func (p *PollAnswerVoters) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pollAnswerVoters#3b6ddad2 to nil")
	}
	if err := b.ConsumeID(PollAnswerVotersTypeID); err != nil {
		return fmt.Errorf("unable to decode pollAnswerVoters#3b6ddad2: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode pollAnswerVoters#3b6ddad2: field flags: %w", err)
		}
	}
	p.Chosen = p.Flags.Has(0)
	p.Correct = p.Flags.Has(1)
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode pollAnswerVoters#3b6ddad2: field option: %w", err)
		}
		p.Option = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode pollAnswerVoters#3b6ddad2: field voters: %w", err)
		}
		p.Voters = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PollAnswerVoters.
var (
	_ bin.Encoder = &PollAnswerVoters{}
	_ bin.Decoder = &PollAnswerVoters{}
)
