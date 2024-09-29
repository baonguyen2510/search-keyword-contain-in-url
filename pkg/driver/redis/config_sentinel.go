package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

// SentinelConnection -- redis connection
type SentinelConnection struct {
	MasterGroup       string
	SentinelAddresses []string
	Password          string
	PoolSize          int
	MaxRetries        int
}

// BuildClient --
func (conn *SentinelConnection) BuildClient() (redis.UniversalClient, error) {
	if len(conn.SentinelAddresses) == 0 {
		return nil, ErrorMissingRedisAddress
	}

	if conn.PoolSize <= 0 {
		conn.PoolSize = DefaultPoolSize
	}

	if conn.MaxRetries <= 0 {
		conn.MaxRetries = DefaultMaxRetries
	}

	masterGroup := conn.MasterGroup
	if masterGroup == "" {
		masterGroup = "master"
	}

	redisClient := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    masterGroup,
		SentinelAddrs: conn.SentinelAddresses,
		Password:      conn.Password,
		DB:            DefaultDB,
		PoolSize:      conn.PoolSize,
		PoolTimeout:   time.Second * 4,
		MaxRetries:    conn.MaxRetries,
	})

	return redisClient, nil
}
