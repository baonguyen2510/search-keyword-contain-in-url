package credis

import (
	"context"
	"search-keyword-service/pkg/cache"
	driverRedis "search-keyword-service/pkg/driver/redis"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	client *driverRedis.Client
}

func NewClient(client *driverRedis.Client) *Client {
	return &Client{
		client: client,
	}
}

func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	value, err := c.client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, cache.ErrNotFound
		}
		return nil, err
	}

	return value, nil
}

func (c *Client) Set(ctx context.Context, key string, value []byte) error {
	return c.client.Set(ctx, key, value, redis.KeepTTL).Err()
}

func (c *Client) SetEx(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *Client) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}
