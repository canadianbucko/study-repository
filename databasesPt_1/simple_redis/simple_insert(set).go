package simple_redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func SetInsertRedis(rdb *redis.Client, ctx context.Context, key string, value any) error {
	err := rdb.Set(ctx, key, value, 5*time.Minute).Err()
	return err
}

// err := rdb.Set(ctx, "grok", "username", 5*time.Minute).Err()
