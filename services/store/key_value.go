package store

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
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

func init() {
	services.Register(func(builder *di.Builder, cfg config.Configuration) error {
		redisAddr := fmt.Sprintf("%s:%d", cfg.KeyValueStore.Host, cfg.KeyValueStore.Port)
		redisPass := cfg.KeyValueStore.Password
		redisDb := cfg.KeyValueStore.Db

		return builder.Add(di.Def{
			Name: services.KeyValue,
			Build: func(ctn di.Container) (interface{}, error) {
				redisClient := redis.NewClient(&redis.Options{
					Addr:     redisAddr,
					Password: redisPass,
					DB:       redisDb,
				})

				return &RedisStore{redisClient}, nil
			},
		})
	})
}
