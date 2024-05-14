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
		panic(err2)
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

	// add key with prefix
	prefix := "user:"
	identifier := 1
	keyWithPrefix := fmt.Sprintf("%s%d", prefix, identifier)
	ttl3 := time.Duration(1) * time.Minute
	err3 := rdb.Set(ctx, keyWithPrefix, "value1", ttl3).Err()
	if err3 != nil {
		panic(err3)
	}

	keys3 := rdb.Keys(ctx, prefix+"*").Val()
	if len(keys3) == 0 {
		fmt.Println("no keys found with prefix")
		return
	}
	for _, key3 := range keys3 {
		val3, err3a := rdb.Get(ctx, key3).Result()
		if err3a != nil {
			fmt.Println("error get value for key", key3)
		} else {
			fmt.Println("key", key3, "value", val3)
		}
	}

}
