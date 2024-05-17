package middleware

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
	"trade/config"
)

var (
	ctx    = context.Background()
	Client *redis.Client
)

func RedisConnect() {
	loadConfig, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}
	redisAddr := fmt.Sprintf("%s:%s", loadConfig.GormConfig.Redis.Host, loadConfig.GormConfig.Redis.Port)
	Client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Username: loadConfig.GormConfig.Redis.Username,
		Password: loadConfig.GormConfig.Redis.Password,
		DB:       loadConfig.GormConfig.Redis.DB,
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
