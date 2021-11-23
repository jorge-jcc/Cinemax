package ports

import (
	"context"
	"mime/multipart"
	"os"
)

type ImageRepository interface {
	LoadImage(ctx context.Context, peliculaId string, file *multipart.FileHeader) error
	DownloadImage(ctx context.Context, fileName string) (*os.File, error)
}
