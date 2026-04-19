package simple_redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func IncreaseRedisCount(rdb *redis.Client, ctx context.Context, key string) (result int64, err error) {
	result, err = rdb.Incr(ctx, key).Result()
	return result, err
}

func IncreaseByRedisCount(rdb *redis.Client, ctx context.Context, key string, increaseBy int) (result int64, err error) {
	result, err = rdb.IncrBy(ctx, key, int64(increaseBy)).Result()
	return result, err
}
