package redis

import (
	"fmt"
	"time"

	"github.com/anditakaesar/uwa-back/internal/env"
	"github.com/go-redis/redis/v7"
)

type InternalRedis struct {
	Client *redis.Client
}

func NewInternalRedis() *InternalRedis {
	return &InternalRedis{}
}

func (r *InternalRedis) Connect() error {
	redisPassword := env.RedisPassword()
	redisOption := &redis.Options{
		Addr: fmt.Sprintf("%s:%s", env.RedisHost(), env.RedisPort()),
		DB:   env.RedisDB(),
	}

	if redisPassword != "" {
		redisOption.Password = redisPassword
	}

	redisClient := redis.NewClient(redisOption)
	_, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}

	r.Client = redisClient

	return nil
}

func (r *InternalRedis) GetRedisValue(redisKey string) (string, error) {
	redisValue, err := r.Client.Get(redisKey).Result()
	if err != nil {
		return "", err
	}
	return redisValue, nil
}

func (r *InternalRedis) SetRedisValue(redisKey string, redisValue interface{}, timeOutPeriod time.Duration) {
	r.Client.Set(redisKey, redisValue, timeOutPeriod)
}

func (r *InternalRedis) DeleteRedisValue(redisKey string) {
	r.Client.Del(redisKey)
}

func (r *InternalRedis) GetTTLValue(redisKey string) int64 {
	return r.Client.TTL(redisKey).Val().Microseconds()

}
