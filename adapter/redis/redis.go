package redis

import "time"

type RedisInterface interface {
	GetRedisValue(redisKey string) (string, error)
	SetRedisValue(redisKey string, redisValue interface{}, timeOutPeriod time.Duration)
	DeleteRedisValue(redisKey string)
	GetTTLValue(redisKey string) int64
}
