package redis

import (
	"errors"

	"github.com/redis/go-redis/v9"
)

// Connection --
type Config interface {
	BuildClient() (redis.UniversalClient, error)
}

var (
	// ErrorMissingRedisAddress --
	ErrorMissingRedisAddress = errors.New("missing redis address")

	// ErrorRedisClientNotSupported --
	ErrorRedisClientNotSupported = errors.New("redis client not supported")
)

const (
	DefaultDB         = 0
	DefaultPoolSize   = 100
	DefaultMaxRetries = 2
)
