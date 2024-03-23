package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := connRdb()
	ctx := context.Background()
	err := rdb.Set(ctx, "session_id:admin", "testvalue", time.Second*5).Err()
	if err != nil {
		panic(err)
	}
	sessionID, err := rdb.Get(ctx, "session_id:admin").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(sessionID)
}

func connRdb() *redis.Client {
	// redis-cli
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
