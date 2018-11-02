package config

import (
	"github.com/go-redis/redis"
	"log"
)

type RedisWriter struct {
	client *redis.Client
}

func (redis RedisWriter) Push(key string, value interface{}) error {
	status := redis.client.LPush(key, value)

	return status.Err()
}

func (redis RedisWriter) Trim(key string, cnt int) {
	log.Println(cnt)
	redis.client.LTrim(key, 0, int64(cnt))
}
