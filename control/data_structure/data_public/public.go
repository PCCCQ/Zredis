package data_public

import (
	"changeme/utils"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func Del(ctx context.Context, client *redis.Client, key string) (isTrue bool) {
	tmp := client.Del(ctx, key).Val()
	if tmp == 1 {
		isTrue = true
	} else {
		isTrue = false
	}
	return
}
func SetKey(ctx context.Context, client *redis.Client, oldKey string, newKey string) (result string, err error) {
	result, err = client.Rename(ctx, oldKey, newKey).Result()
	return

}

func SetTTL(ctx context.Context, client *redis.Client, key string, TTL float32) (isTrue bool) {
	return client.Expire(ctx, key, time.Duration(TTL)*time.Second).Val()
}

func GetTTL(ctx context.Context, client *redis.Client, key string) float32 {
	return utils.FormatTtl(ctx, client, key)
}
