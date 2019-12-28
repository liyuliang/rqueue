package system

import (
	"github.com/go-redis/redis/v7"
)

func init(){
	redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
