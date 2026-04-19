package simple_redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func GetSelectRedis(ctx context.Context, rdb *redis.Client, key string) (value string, err error) {
	value, err = rdb.Get(ctx, key).Result()
	return value, err
}
