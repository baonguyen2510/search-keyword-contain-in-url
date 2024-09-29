package redis

import (
	"context"
	"search-keyword-service/configs"
	"search-keyword-service/pkg/driver/redis"
	"strings"
)

var client *redis.Client

func New(ctx context.Context) (*redis.Client, error) {
	var conn redis.Config
	if configs.Config.RedisCluster {
		conn = &redis.ClusterConnection{
			ClusterAddresses: strings.Split(configs.Config.RedisAddr, ","),
			Password:         configs.Config.RedisPassword,
		}
	} else if configs.Config.RedisSingle {
		conn = &redis.SingleConnection{
			Address:  configs.Config.RedisAddr,
			Password: configs.Config.RedisPassword,
		}
	} else if configs.Config.RedisSentinel {
		conn = &redis.SentinelConnection{
			MasterGroup:       configs.Config.RedisSentinelMasterGroup,
			SentinelAddresses: strings.Split(configs.Config.RedisAddr, ","),
			Password:          configs.Config.RedisPassword,
		}
	}

	c, err := redis.NewConnection(ctx, conn)
	if err != nil {
		return nil, err
	}
	client = c
	return client, nil
}

func GetClient() *redis.Client {
	return client
}
