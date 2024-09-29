package redis

import (
	"search-keyword-service/pkg/log"
	"time"

	"github.com/redis/go-redis/v9"
)

// SingleConnection -- redis connection
type SingleConnection struct {
	Address    string
	Password   string
	MaxRetries int
	PoolSize   int
}

// BuildClient -- build single redis client
func (conn *SingleConnection) BuildClient() (redis.UniversalClient, error) {
	if conn.Address == "" {
		return nil, ErrorMissingRedisAddress
	}

	if conn.PoolSize <= 0 {
		conn.PoolSize = DefaultPoolSize
	}

	if conn.MaxRetries <= 0 {
		conn.MaxRetries = DefaultMaxRetries
	}

	log.Infof("[redis] single - address: %v, pass: %v, db: %v, pollSize: %v",
		conn.Address, "***", DefaultDB, conn.PoolSize)

	return redis.NewClient(
		&redis.Options{
			Addr:        conn.Address,
			Password:    conn.Password,
			DB:          DefaultDB,
			PoolSize:    conn.PoolSize,
			PoolTimeout: time.Second * 4,
			MaxRetries:  conn.MaxRetries,
		},
	), nil
}
