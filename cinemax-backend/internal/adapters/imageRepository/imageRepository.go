package imagerepository

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type imageRepository struct {
	path string
}

func NewImageRepository(path string) *imageRepository {
	return &imageRepository{path}
}

func (i *imageRepository) LoadImage(ctx context.Context, fileName string, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	path := fmt.Sprintf("%s/%s", i.path, fileName)
	fmt.Println(path)
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func (i *imageRepository) DownloadImage(ctx context.Context, fileName string) (*os.File, error) {
	path := fmt.Sprintf("%s/%s", i.path, fileName)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}
