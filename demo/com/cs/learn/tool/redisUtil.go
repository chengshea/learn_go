package tool

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client
var ctx = context.Background()

func initRedis() *redis.Client {
	conf := getConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.GetString("redis.addr"),
		Username: conf.GetString("redis.username"),
		Password: conf.GetString("redis.pwd"), // no password set
		DB:       conf.GetInt("redis.db"),     // use default DB
	})
	return rdb

}

func OpenConRedis() *redis.Client {
	log.Println("OpenCon======", client)
	if client == nil {
		log.Println("===init===")
		client = initRedis()
	}
	return client
}

func IsExist(err error) bool {
	if err == redis.Nil {
		return false
	}
	return true
}

func P() {
	val, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("ping err:", err)
	}
	log.Println("redis>", val)
}
