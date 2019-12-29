package system

import (
	"github.com/go-redis/redis/v7"
)

var client *redis.Client

func conn(redisUri string) *redis.Client {

	if client == nil {
		opt, err := redis.ParseURL(redisUri)
		if err != nil {
			panic(err)
		}

		client = redis.NewClient(opt)
	}
	return client
}

func Client() *redis.Client{
	redisUri := Config()["redisUri"]
	if redisUri == "" {
		panic("redis connect address is required")
	}

	return conn(redisUri)
}