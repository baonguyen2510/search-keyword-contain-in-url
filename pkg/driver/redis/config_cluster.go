package redis

import (
	"search-keyword-service/pkg/log"
	"time"

	"github.com/redis/go-redis/v9"
)

// ClusterConnection -- redis connection
type ClusterConnection struct {
	ClusterAddresses []string
	Password         string
	PoolSize         int
	MaxRetries       int
}

// BuildClient --
func (conn *ClusterConnection) BuildClient() (redis.UniversalClient, error) {
	if len(conn.ClusterAddresses) == 0 {
		return nil, ErrorMissingRedisAddress
	}

	if conn.PoolSize <= 0 {
		conn.PoolSize = DefaultPoolSize
	}

	if conn.MaxRetries <= 0 {
		conn.MaxRetries = DefaultMaxRetries
	}

	log.Infof("[redis] Create cluster client to %v", conn.ClusterAddresses)

	redisDB := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       conn.ClusterAddresses,
		Password:    conn.Password,
		PoolSize:    conn.PoolSize,
		PoolTimeout: time.Second * 4,
		MaxRetries:  conn.MaxRetries,
	})

	return redisDB, nil
}
