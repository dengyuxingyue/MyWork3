package cache

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
)

var Rd *redis.Client

func InitRedis(RedisDbName, RedisAddr, RedisPw string) {
	redisDbName, _ := strconv.Atoi(RedisDbName)
	client := redis.NewClient(&redis.Options{

		Addr:     RedisAddr,
		DB:       redisDbName,
		Password: RedisPw,
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Fatalln(err)
	}
	Rd = client
	fmt.Println("redis链接成功")
	//utils.LogrusObj.Infoln("Redis init success!")
}
