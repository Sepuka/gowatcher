package config

import (
	"github.com/go-redis/redis"
)

type RedisStore struct {
	client *redis.Client
}

func (redis RedisStore) Push(key string, value interface{}) error {
	status := redis.client.LPush(key, value)

	return status.Err()
}

func (redis RedisStore) Trim(key string, cnt int) {
	redis.client.LTrim(key, 0, int64(cnt))
}

func (redis RedisStore) List(key string) []string {
	length := redis.client.LLen(key)
	cmd := redis.client.LRange(key, 0, length.Val())

	return cmd.Val()
}
