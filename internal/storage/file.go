package storage

import "context"

type fileStorage struct {
	fileName string
}

func NewFileStorage(fileName string) Storage {
	return &fileStorage{fileName: fileName}
}

func (s *fileStorage) SaveImage(ctx context.Context, img string, token string) error {
	return nil
}

func (s *fileStorage) GetImage(ctx context.Context, token string) (img string, err error) {
	return "", nil
}
