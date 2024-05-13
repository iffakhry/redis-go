package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "qwertypoiu123", // write "" if no password set
		DB:       0,               // use default DB. Default=0
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	// err2 := rdb.Set(ctx, "key2", "value2", 0).Err()
	// if err2 != nil {
	// 	panic(err)
	// }

	ttl := time.Duration(3) * time.Second
	err2 := rdb.Set(ctx, "key2", "value2", ttl).Err()
	if err2 != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
