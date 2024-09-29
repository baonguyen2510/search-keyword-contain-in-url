package redislock

import (
	"context"

	"search-keyword-service/pkg/driver/redis"
	"search-keyword-service/pkg/lock"

	goredis "github.com/redis/go-redis/v9"
)

var (
	luaRelease = goredis.NewScript(`if redis.call("get", KEYS[1]) == ARGV[1] then` +
		` return redis.call("del", KEYS[1]) else return 0 end`)
)

type redisLock struct {
	client *redis.Client
	key    string
	value  string
	ctx    context.Context
}

func (r *redisLock) Release() error {
	if r.key == "" {
		return nil
	}
	res, err := luaRelease.Run(r.ctx, r.client, []string{r.key}, r.value).Result()
	if err == goredis.Nil {
		return nil
	} else if err != nil {
		return err
	}

	if i, ok := res.(int64); !ok || i != 1 {
		return lock.ErrLockNotHeld
	}
	return nil
}
