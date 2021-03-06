package downloader

import (
	"context"

	"github.com/gotd/td/tg"
)

// CDN represents Telegram RPC client to CDN server.
type CDN interface {
	UploadGetCdnFile(ctx context.Context, request *tg.UploadGetCdnFileRequest) (tg.UploadCdnFileClass, error)
}

// Client represents Telegram RPC client.
type Client interface {
	UploadGetFile(ctx context.Context, request *tg.UploadGetFileRequest) (tg.UploadFileClass, error)
	UploadGetFileHashes(ctx context.Context, request *tg.UploadGetFileHashesRequest) ([]tg.FileHash, error)

	UploadReuploadCdnFile(ctx context.Context, request *tg.UploadReuploadCdnFileRequest) ([]tg.FileHash, error)
	UploadGetCdnFileHashes(ctx context.Context, request *tg.UploadGetCdnFileHashesRequest) ([]tg.FileHash, error)

	UploadGetWebFile(ctx context.Context, request *tg.UploadGetWebFileRequest) (*tg.UploadWebFile, error)
}

type chunk struct {
	data []byte
	tag  tg.StorageFileTypeClass
}

// schema is simple interface for different download schemas.
type schema interface {
	Chunk(ctx context.Context, offset, limit int) (chunk, error)
	Hashes(ctx context.Context, offset int) ([]tg.FileHash, error)
}
