package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
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
	redisAddr := fmt.Sprintf("%s:%s", loadConfig.Redis.Host, loadConfig.Redis.Port)
	Client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Username: loadConfig.Redis.Username,
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

func AcquireLock(key string, time int64, expiration time.Duration) bool {
	result, err := Client.SetNX(ctx, key, time, expiration).Result()
	if err != nil {
		log.Println("Error acquiring lock:", err)
		return false
	}
	return result
}

func ReleaseLock(key string) {
	_, err := Client.Del(ctx, key).Result()
	if err != nil {
		log.Println("Error releasing lock:", err)
	}
}
