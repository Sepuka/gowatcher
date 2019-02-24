package store

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sarulabs/di"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/services/store"
	"io"
)

const DefStoreRedis = "definition.store.redis"

func init() {
	services.Register(func(builder *di.Builder, cfg config.Configuration) error {
		return builder.Add(di.Def{
			Name: DefStoreRedis,
			Build: func(ctn di.Container) (interface{}, error) {
				redisAddr := fmt.Sprintf("%v:%d", cfg.KeyValueStore.Host, cfg.KeyValueStore.Port)
				redisPass := cfg.KeyValueStore.Password
				redisDb := cfg.KeyValueStore.Db

				redisClient := redis.NewClient(&redis.Options{
					Addr:     redisAddr,
					Password: redisPass,
					DB:       redisDb,
				})

				return &store.RedisStore{Client: redisClient}, nil
			},
			Close: func(obj interface{}) error {
				return obj.(io.Closer).Close()
			},
		})
	})
}
