package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type RedisCacheInterface interface {
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Keys(ctx context.Context, pattern string) *redis.StringSliceCmd
}

type RedisCache struct {
	ctx	   context.Context
	logger *logrus.Logger
	client *redis.Client
}

func NewRedisCache(ctx context.Context, logger *logrus.Logger, client *redis.Client) *RedisCache {
	return &RedisCache{
		ctx:	ctx,
		logger: logger,
		client: client,
	}
}

func (c *RedisCache) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.client.Set(c.ctx, key, value, expiration)
}

func (c *RedisCache) Get(ctx context.Context, key string) *redis.StringCmd {
	return c.client.Get(c.ctx, key)
}

func (c *RedisCache) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return c.client.Del(c.ctx, keys...)
}

func (c *RedisCache) Keys(ctx context.Context, pattern string) *redis.StringSliceCmd {
	return c.client.Keys(c.ctx, pattern)
}
