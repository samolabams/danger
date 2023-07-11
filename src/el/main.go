package main

import (
	"fmt"
	"samolabams/redis"
)

func main() {
	err := redis.ConnectToRedis()
	if err != nil {
		panic(err)
	}

	fmt.Print("Connected to redis")
}
