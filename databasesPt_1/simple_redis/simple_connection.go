package simple_redis

import (
	"context"
	"os"
	"strconv"

	"github.com/k0kubun/pp"
	"github.com/redis/go-redis/v9"
)

func Connect(ctx context.Context) (connection *redis.Client, err error) {
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASS")
	redisDefaultDB := os.Getenv("REDIS_DEFAULT_DB")
	redisProtocol := os.Getenv("REDIS_PROTOCOL")

	redisDefaultDBint, err := strconv.Atoi(redisDefaultDB)
	if err != nil || redisDefaultDB == "" {
		pp.Println("either default DB is empty or it's in retarded language, here's the error that triggered it anyways", err)
	}

	redisProtocolint, err := strconv.Atoi(redisProtocol)
	if err != nil || redisProtocol == "" {
		pp.Println("either protocol is empty or it's in retarded language, here's the error that triggered it anyways", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
		DB:       redisDefaultDBint,
		Protocol: redisProtocolint,
	})

	_, err = rdb.Ping(ctx).Result()
	return rdb, err

}
