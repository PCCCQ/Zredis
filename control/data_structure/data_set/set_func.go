package data_set

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func (s *SetData) SetSet(ctx context.Context, RedisDB *redis.Client) (err error) {
	result, err := RedisDB.SAdd(ctx, s.Key, s.Value).Result()
	//fmt.Println("执行2", result, s.Key, s.Value)
	if result == 1 {
		return nil
	}
	return err
}

func (s *SetData) GetSet(ctx context.Context, RedisDB *redis.Client) (data SetData, err error) {
	data.Key = s.Key
	data.Value, err = RedisDB.SMembers(ctx, s.Key).Result()
	if err != nil {
		return
	}

	return
}

func (s *SetData) DelSetOne(ctx context.Context, RedisDB *redis.Client) (data any, err error) {
	data, err = RedisDB.SRem(ctx, s.Key, s.Value).Result()
	return
}
