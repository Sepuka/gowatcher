package config

import (
	"github.com/go-redis/redis"
	"log"
	"fmt"
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
	fmt.Println()
	redis.client.LTrim(key, 0, int64(cnt))
}
