package system

import (
	"github.com/go-redis/redis/v7"
)

var client *redis.Client

func Conn(redisUri string) *redis.Client {

	if client == nil {
		opt, err := redis.ParseURL(redisUri)
		if err != nil {
			panic(err)
		}

		client = redis.NewClient(opt)

	}
	return client
}
