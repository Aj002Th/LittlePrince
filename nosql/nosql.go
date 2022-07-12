package nosql

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Aj002Th/LittlePrince/data"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
	rdb = data.GetRedisConn()
}

func Set(key string, data interface{}, expire time.Duration) error {
	ctx := context.Background()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = rdb.Set(ctx, key, value, expire).Err()
	if err != nil {
		return err
	}

	return nil
}

func Exists(key string) bool {
	ctx := context.Background()

	err := rdb.Exists(ctx, key).Err()
	if err != nil {
		return false
	}

	return true
}

func Get(key string) (string, error) {
	ctx := context.Background()

	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

func Delete(key string) error {
	ctx := context.Background()

	err := rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

func LikeDeletes(key string) error {
	return nil
}
