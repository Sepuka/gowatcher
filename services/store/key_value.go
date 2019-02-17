package store

import (
	"github.com/go-redis/redis"
)

type RedisStore struct {
	Client *redis.Client
}

func (redis RedisStore) Push(key string, value interface{}) error {
	status := redis.Client.LPush(key, value)

	return status.Err()
}

func (redis RedisStore) Trim(key string, cnt int) {
	redis.Client.LTrim(key, 0, int64(cnt))
}

func (redis RedisStore) List(key string) []string {
	length := redis.Client.LLen(key)
	cmd := redis.Client.LRange(key, 0, length.Val())

	return cmd.Val()
}
