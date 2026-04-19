package simple_redis

import (
	"context"

	"github.com/k0kubun/pp"
	"github.com/redis/go-redis/v9"
)

// добавить в конец списка
func RPush(ctx context.Context, rdb *redis.Client, key string, value any) (err error) {
	err = rdb.RPush(ctx, key, value).Err()
	return err
}

// FIFO first in first out (LPOP) т.е. слева берем
func HandleTask(ctx context.Context, rdb *redis.Client, key string) (task string, err error) {
	task, err = rdb.LPop(ctx, key).Result()
	if err == nil {
		pp.Println("we did something with the task + deleted it already! gut!")
	}
	return task, err
}
