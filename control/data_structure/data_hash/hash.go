package data_hash

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type HashValue struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type HashData struct {
	Key       string      `json:"key,omitempty"`
	ValueList []HashValue `json:"valueList,omitempty"`
}

func (h *HashData) GetHash(ctx context.Context, RedisDB *redis.Client) (err error) {
	tmp, err := RedisDB.HGetAll(ctx, h.Key).Result()
	if err != nil {
		return
	}

	for i, v := range tmp {
		h.ValueList = append(h.ValueList, HashValue{
			Key:   i,
			Value: v,
		})
	}
	fmt.Println(h)
	return err
}

//func (h *HashData) SetHash(ctx context.Context, RedisDB *redis.Client) (err error) {
//	h.Value, err = RedisDB.HGetAll(ctx, h.Key).Result()
//	return
//}
//
//
//func (h *HashData) DelHashOne(ctx context.Context, RedisDB *redis.Client) (err error) {
//	h.Value, err = RedisDB.HDel(ctx, h.Key,).Result()
//	return
//}
