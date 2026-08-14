package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/renji18/rms/utils"
)

var ctx = context.Background()
var RedisClient *redis.Client

func InitializeRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     utils.Config.REDIS_ADDR + ":" + utils.Config.REDIS_PORT,
		Password: "",
		DB:       0,
	})

	err := RedisClient.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		utils.Fatal(fmt.Errorf("Error initializing redis: %v", err), err)
	}

	fmt.Println("Redis up!!")
}

func SetRedis(key string, value any, expiration time.Duration) error {
	err := RedisClient.Set(ctx, key, value, expiration).Err()

	if err != nil {
		return fmt.Errorf("Error setting value in redis: %v", err)
	}

	return nil
}

func GetRedis(key string) (value any, err error) {
	value, err = RedisClient.Get(ctx, key).Result()

	if err == redis.Nil {
		return nil, fmt.Errorf("Key %s does not exist", key)
	} else if err != nil {
		return nil, fmt.Errorf("Error getting value from redis: %v", err)
	} else {
		return value, nil
	}
}

func DelRedis(key string) error {
	err := RedisClient.Del(ctx, key).Err()

	if err != nil {
		return fmt.Errorf("Error deleting key from redis: %v", err)
	}

	return nil
}
