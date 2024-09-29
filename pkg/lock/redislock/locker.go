package redislock

import (
	"context"
	"time"

	"search-keyword-service/pkg/driver/redis"
	"search-keyword-service/pkg/id"
	"search-keyword-service/pkg/lock"
)

type RedisLocker struct {
	client *redis.Client
}

func NewLocker(client *redis.Client) *RedisLocker {
	return &RedisLocker{
		client: client,
	}
}

func (r *RedisLocker) Obtain(ctx context.Context, key string, ttl time.Duration) (lock.Lock, error) {
	value := id.NewULID()
	ok, err := r.client.SetNX(ctx, key, value, ttl).Result()
	if ok {
		return &redisLock{r.client, key, value, ctx}, nil
	}

	if err == nil {
		return nil, lock.ErrLockObtained
	}
	return nil, err
}
