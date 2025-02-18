package inits

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func InitRedis() {
	cf := ViperData.Redis
	Client = redis.NewClient(&redis.Options{
		Addr:     cf.Addr,
		Password: cf.Passwd,
		DB:       cf.Data,
	})
	err := Client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("init redis success")

}
