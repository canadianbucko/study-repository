package main

import (
	"context"
	"someproject/simple_redis"

	"github.com/k0kubun/pp"
)

func main() {
	ctx := context.Background()
	rdb, err := simple_redis.Connect(ctx)
	if err != nil {
		panic(err)
	}

	if err := simple_redis.SetInsertRedis(rdb, ctx, "kiriki", "hello i am kirik"); err != nil {
		panic(err)
	}

	value, err := simple_redis.GetSelectRedis(ctx, rdb, "kiriki")
	if err != nil {
		panic(err)
	}
	pp.Println(value)

	result, err := simple_redis.IncreaseByRedisCount(rdb, ctx, "zalupa", 20)
	if err != nil {
		panic(err)
	}
	pp.Println("counter", result)

	if err := simple_redis.DeleteByIDRedis(ctx, rdb, "zalupa"); err != nil {
		panic(err)
	}

	value, err = simple_redis.GetSelectRedis(ctx, rdb, "zalupa")
	pp.Println(value, err) // should be empty

	key := "kirik"
	value = "hi i am kirik"
	if err := simple_redis.RPush(ctx, rdb, key, value); err != nil {
		panic(err)
	}

	task, err := simple_redis.HandleTask(ctx, rdb, key)
	if err != nil {
		panic(err)
	}

	pp.Println(task)
}
