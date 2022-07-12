package data

import (
	"context"

	"github.com/Aj002Th/LittlePrince/pkg/setting"
	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func init() {
	var ctx = context.Background()
	Redis = redis.NewClient(&redis.Options{
		Addr:     setting.Redis.Host,
		Password: setting.Redis.Password,
		DB:       setting.Redis.DB,
		PoolSize: setting.Redis.PoolSize,
	})

	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

func GetRedisConn() *redis.Client {
	return Redis
}
