package simple_redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func DeleteByIDRedis(ctx context.Context, rdb *redis.Client, key string) (err error) {
	err = rdb.Del(ctx, key).Err()
	return err
}
