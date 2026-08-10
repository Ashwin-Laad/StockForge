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

	opt, err := redis.ParseURL("rediss://default:gQAAAAAAAlQUAAIgcDFlZWE5MWQzNjllMDY0ZTg2YjAxMDA1NTI2ZjZlYjk3ZQ@wanted-jay-152596.upstash.io:6379")
	if err != nil {
		panic(err)
	}

	// Create a new Redis client
	// RedisClient = redis.NewClient(opt)
	// //creates an empty context
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// _, err = RedisClient.Ping(ctx).Result()
	// if err == nil {
	// 	print("ping done")
	// } else {
	// 	print("No ping")
	// }

	// println("Redis is connected ")

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

	RedisClient = redis.NewClient(opt)
	fmt.Println("About to ping Redis...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = RedisClient.Ping(ctx).Result()

	fmt.Println("Ping returned")

	if err != nil {
		fmt.Println("Redis ping failed:", err)
		return
	}

	fmt.Println("Redis ping successful")

}

// package database

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// var RedisURL = "https://wanted-jay-152596.upstash.io"
// var RedisToken = "gQAAAAAAAlQUAAIgcDFlZWE5MWQzNjllMDY0ZTg2YjAxMDA1NTI2ZjZlYjk3ZQ"

// func ConnectRedis() {
// 	fmt.Println("Redis REST configured")
// }

// func RedisGet(key string) (string, error) {

// 	url := RedisURL + "/get/" + key

// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return "", err
// 	}

// 	req.Header.Set("Authorization", "Bearer "+RedisToken)

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return "", err
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return "", fmt.Errorf("Redis GET returned status %d", resp.StatusCode)
// 	}

// 	var result struct {
// 		Result *string `json:"result"`
// 	}

// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	if err != nil {
// 		return "", err
// 	}

// 	if result.Result == nil {
// 		return "", nil
// 	}

// 	return *result.Result, nil
// }

// func RedisSet(key string, value string, expiry int) error {

// 	payload := map[string]interface{}{
// 		"command": []string{
// 			"SET",
// 			key,
// 			value,
// 			"EX",
// 			fmt.Sprint(expiry),
// 		},
// 	}

// 	body, err := json.Marshal(payload)
// 	if err != nil {
// 		return err
// 	}

// 	req, err := http.NewRequest(
// 		"POST",
// 		RedisURL,
// 		bytes.NewBuffer(body),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Set("Authorization", "Bearer "+RedisToken)
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return err
// 	}

// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("Redis SET returned status %d", resp.StatusCode)
// 	}

// 	return nil
// }
