package redis

import (
	"context"
	"fmt"
	"search-keyword-service/pkg/log"
	"sync"

	"github.com/redis/go-redis/v9"
)

// Client --
type Client struct {
	redis.UniversalClient
	Slots []redis.ClusterSlot
	lock  sync.Once
}

// NewConnection -- open connection to db
func NewConnection(ctx context.Context, conn Config, hooks ...redis.Hook) (*Client, error) {
	var err error

	c, err := conn.BuildClient()
	if err != nil {
		return nil, err
	}

	pong, err := c.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	log.Info("[redis] Ping to redis: ", pong)

	for _, h := range hooks {
		c.AddHook(h)
	}

	cs, err := getClusterInfo(ctx, c)
	if err != nil {
		return nil, err
	}
	return &Client{UniversalClient: c, Slots: cs}, nil
}

func getClusterInfo(ctx context.Context, c redis.UniversalClient) ([]redis.ClusterSlot, error) {
	var cs = make([]redis.ClusterSlot, 0)
	if ci := c.ClusterInfo(ctx); ci.Err() == nil {
		csr := c.ClusterSlots(ctx)
		var err error
		cs, err = csr.Result()
		if err != nil {
			return nil, err
		}
	}
	return cs, nil
}

func (r *Client) Name() string {
	return "Redis"
}

// Close -- close connection
func (r *Client) Close() error {
	var err error
	r.lock.Do(func() {
		err = r.UniversalClient.Close()
	})
	return err
}

// GetClient --
func (r *Client) GetClient() redis.UniversalClient {
	return r.UniversalClient
}

func (r *Client) IsSingle() bool {
	switch r.UniversalClient.(type) {
	case *redis.Client:
		return true
	default:
		return false
	}
}

// GetClusterSlots -
func (r *Client) GetClusterSlots(ctx context.Context) ([]redis.ClusterSlot, error) {
	res := r.ClusterSlots(ctx)
	return res.Result()
}

// GetRedisSlot -
func (r *Client) GetRedisSlot(key string) int {
	return Slot(key)
}

// GetRedisSlotID -
func (r *Client) GetRedisSlotID(key string) string {
	return GetSlotID(key, r.Slots)
}

// GetSlotID -
func GetSlotID(key string, slots []redis.ClusterSlot) string {
	s := Slot(key)
	for k := range slots {
		slot := slots[k]
		if slot.Start <= s && s <= slot.End {
			return fmt.Sprintf("%v-%v", slot.Start, slot.End)
		}
	}
	return ""
}
