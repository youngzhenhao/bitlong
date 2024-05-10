package middleware

import (
	"AssetsTrade/config"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()
var Client *redis.Client

func RedisConnect() {
	loadConfig, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	Client = redis.NewClient(&redis.Options{
		Addr:     loadConfig.Redis.Addr,
		Password: loadConfig.Redis.Password,
		DB:       loadConfig.Redis.DB,
	})

	_, err = Client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
func RedisSet(key string, value interface{}, expiration time.Duration) error {
	return Client.Set(ctx, key, value, expiration).Err()
}

func RedisGet(key string) (string, error) {
	return Client.Get(ctx, key).Result()
}

func RedisDel(key string) error {
	return Client.Del(ctx, key).Err()
}
