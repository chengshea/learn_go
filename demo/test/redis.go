package test

import (
	"context"
	"demo/com/cs/learn/tool"
	"log"
)

var ctx = context.Background()

func redisGet(name string) {
	client := tool.OpenConRedis()
	val, err := client.Get(ctx, name).Result()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("key:%s get val:%s", name, val)
}
