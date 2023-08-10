package data_string

import (
	"changeme/utils"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

// SetString 给string类型赋值
func (s *StringData) SetString(ctx context.Context, RedisDB *redis.Client) (data any, err error) {
	data, err = RedisDB.Set(ctx, s.Key, s.Value, time.Duration(s.TTL)*time.Second).Result()
	return

}

// GetString 得到string类型的值
func (s *StringData) GetString(ctx context.Context, RedisDB *redis.Client) (data StringData, err error) {
	data.Key = s.Key
	data.Value, err = RedisDB.Get(ctx, s.Key).Result()
	data.TTL = utils.FormatTtl(ctx, RedisDB, s.Key)
	return
}

// DelString 删除String类型的值
func (s *StringData) DelString(ctx context.Context, RedisDB *redis.Client) (data any, err error) {
	data, err = RedisDB.Del(ctx, s.Key).Result()
	return
}
