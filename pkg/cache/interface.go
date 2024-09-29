package cache

import (
	"context"
	"errors"
	"time"
)

// Errors
var (
	ErrNotFound = errors.New("key not found")
)

type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte) error
	SetEx(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
}
