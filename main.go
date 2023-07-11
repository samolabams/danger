package main

import "fmt"

type CacheOptions struct {
	Mock      bool
	NoEncrypt bool
}

func main() {
	options := CacheOptions{NoEncrypt: true}

	fmt.Println("Intialize cache:", options)
}
