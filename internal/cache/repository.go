package cache

import "context"

type Repository interface {
	Set(ctx context.Context, key string, val any) error
	Get(ctx context.Context, key string, target any) error

	GetBytes(ctx context.Context, key string) ([]byte, error)

	Close() error
}
