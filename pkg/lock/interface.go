package lock

import (
	"context"
	"errors"
	"time"
)

var (
	ErrLockObtained = errors.New("already obtained")
	ErrLockNotHeld  = errors.New("lock not held")
)

type Lock interface {
	Release() error
}

type Locker interface {
	Obtain(ctx context.Context, key string, ttl time.Duration) (Lock, error)
}
