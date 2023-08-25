package storage

import "context"

type Storage interface {
	SaveImage(ctx context.Context, img string, token string) error
	GetImage(ctx context.Context, token string) (img string, err error)
}
