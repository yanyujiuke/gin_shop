package common

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
}

var redisCtx = context.Background()
var RedisDb *redis.Client

func init() {
	// config, err1 := ini.LooseLoad("./config/app.ini")
	// if err1 != nil {
	// 	fmt.Println(err1)
	// }
	//
	// host := config.Section("redis").Key("host").String()
	// port := config.Section("redis").Key("host").String()
	// println(host + ":" + port)
	// RedisDb = redis.NewClient(&redis.Options{
	// 	Addr:     host + ":" + port,
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })

	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "192.168.31.100:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err2 := RedisDb.Ping(redisCtx).Result()
	if err2 != nil {
		println(err2)
	}
}

func (RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return RedisDb.Set(redisCtx, key, value, expiration).Err()
}

func (RedisClient) Get(key string) (string, bool) {
	val, err := RedisDb.Get(redisCtx, key).Result()
	if err == redis.Nil {
		return "key does not exist", false
	} else if err != nil {
		return "", false
	} else {
		return val, true
	}
}

func (RedisClient) Del(key string) error {
	return RedisDb.Del(redisCtx, key).Err()
}
