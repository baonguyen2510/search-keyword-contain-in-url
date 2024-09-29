package cache

import (
	"context"
	"encoding/json"
	"time"
)

var c Cache

func Init(client Cache) error {
	c = client
	return nil
}

func Client() Cache {
	return c
}

func Get[B any](ctx context.Context, key string) (B, error) {
	var res B
	value, err := Client().Get(ctx, key)
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(value, &res); err != nil {
		return res, err
	}
	return res, nil
}

func Set(ctx context.Context, key string, value any) error {
	entry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return Client().Set(ctx, key, entry)
}

func SetEx(ctx context.Context, key string, value any, ttl time.Duration) error {
	entry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return Client().SetEx(ctx, key, entry, ttl)
}

func Delete(ctx context.Context, key string) error {
	return Client().Delete(ctx, key)
}
