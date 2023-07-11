package redis

import "fmt"

type cacheOptions struct {
	NoEncrypt bool
}

func ConnectToRedis() error {
	options := cacheOptions{NoEncrypt: true}

	fmt.Println("connecting:", options)

	return nil
}
