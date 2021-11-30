package main

import (
	"context"
	"fmt"
	//"github.com/go-redis/redis"

	//"github.com/go-redis/redis"
	redis8 "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	//c, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//	fmt.Println("Connect to redis error", err)
	//	return
	//}
	//defer c.Close()
	//
	//_, err = c.Do("SET", "mykey2", "superWang2")
	//if err != nil {
	//	fmt.Println("redis set failed:", err)
	//}
	//
	//username, err := redis.String(c.Do("GET", "mykey2"))
	//if err != nil {
	//	fmt.Println("redis get failed:", err)
	//} else {
	//	fmt.Printf("Get mykey: %v \n", username)
	//}

	//--------------------------------------------------------------------
	ExampleClient8()

}

func ExampleClient8() {
	rdb := redis8.NewClient(&redis8.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis8.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
