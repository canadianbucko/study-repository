package simple_redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func Connect(ctx context.Context) (connection *redis.Client, err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	_, err = rdb.Ping(ctx).Result()
	return rdb, err

}
