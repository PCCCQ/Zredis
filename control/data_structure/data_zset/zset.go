package data_zset

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type ZsetData struct {
	Key       string    `json:"key,omitempty"`
	ValueList []redis.Z `json:"valueList,omitempty"`
	Value     redis.Z   `json:"value,omitempty"`
}

func (z *ZsetData) GetZset(ctx context.Context, RedisDB *redis.Client) (err error) {
	z.ValueList, err = RedisDB.ZRangeWithScores(ctx, z.Key, 0, -1).Result()
	return err
}

func (z *ZsetData) AddZset(ctx context.Context, RedisDB *redis.Client) (err error) {
	result, err := RedisDB.ZAdd(ctx, z.Key, z.Value).Result()
	fmt.Println("zet添加", z.Value)
	if result != 1 {
		return errors.New("失败")
	}
	return err
}

func (z *ZsetData) DelZsetOne(ctx context.Context, RedisDB *redis.Client) (err error) {
	return RedisDB.ZRem(ctx, z.Key, z.Value.Member.(string)).Err()
}

func (z *ZsetData) SetZsetScore(ctx context.Context, RedisDB *redis.Client) (err error) {

	return RedisDB.ZIncrBy(ctx, z.Key, z.Value.Score, z.Value.Member.(string)).Err()
}
