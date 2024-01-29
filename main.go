package main

import (
	"redisgo/model"
	"redisgo/redis"
)

func main() {

	redis.InitRedis()

	redis.HsetData(model.Users, "users")
	redis.HgetData("user")

}
