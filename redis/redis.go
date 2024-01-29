package redis

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	ctx    = context.Background()
)

func InitRedis() {
	// connect to redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// ping the server
	pong, err := client.Ping(ctx).Result()

	log.Println(pong, err)

}

func HsetData(data interface{}, key string) {
	// marshalling the users list
	marshall_data, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	//adding the users into redis
	err = client.HSet(ctx, key, marshall_data).Err()
	if err != nil {
		panic(err)
	}

}

func HgetData(key string) {
	//getting the users from redis
	datas := client.HGetAll(ctx, key).Val()
	log.Println(datas)

}
