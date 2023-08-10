package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func FormatTtl(ctx context.Context, client *redis.Client, key string) float32 {
	ttl := client.TTL(ctx, key).Val()
	if ttl <= -1 {
		return float32(ttl)
	}
	return float32(ttl / time.Second)
}
