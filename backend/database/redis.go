package database

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client // to export

func ConnectRedis() {
	fmt.Println("entered the connect ")

	opt, err := redis.ParseURL("redis://default:gQAAAAAAAlQUAAIgcDFlZWE5MWQzNjllMDY0ZTg2YjAxMDA1NTI2ZjZlYjk3ZQ@wanted-jay-152596.upstash.io:6379")
	if err != nil {
		panic(err)
	}

	// Create a new Redis client
	RedisClient = redis.NewClient(opt)
	//creates an empty context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = RedisClient.Ping(ctx).Result()
	if err == nil {
		print("ping done")
	} else {
		print("No ping")
	}

	println("Redis is connected ")

	// ctx := context.Background() //creates an empty context
	// err = client.Set(ctx, "name", "Test", 5*time.Minute).Err()
	// if err != nil {
	// 	println("could not set", err)
	// }

	// test, errr := client.Get(ctx, "name").Result()
	// if errr != nil {
	// 	println("could not set", err)
	// }

	// println(test)

}
