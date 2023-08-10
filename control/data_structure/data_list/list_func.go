package data_list

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func (l *ListData) GetList(ctx context.Context, RedisDB *redis.Client) (data ListData, err error) {
	data.Value, err = RedisDB.LRange(ctx, l.Key, 0, -1).Result()
	if err != nil {
		return ListData{}, err
	}
	return
}

func (l *ListData) SetList(ctx context.Context, RedisDB *redis.Client) (err error) {

	if l.AddTab == "after" {
		err = RedisDB.RPush(ctx, l.Key, l.Value).Err()
	} else {
		err = RedisDB.LPush(ctx, l.Key, l.Value).Err()
	}
	//fmt.Println("执行2", result, s.Key, s.Value)
	return err
}

func (l *ListData) SetListSet(ctx context.Context, RedisDB *redis.Client) (err error) {
	ok, err := RedisDB.LSet(ctx, l.Key, l.Index, l.Value).Result()
	//fmt.Println("执行2", result, s.Key, s.Value)
	if ok == "ok" {
		return nil
	}
	return err
}

func (l *ListData) DelListOne(ctx context.Context, RedisDB *redis.Client) (err error) {
	err = RedisDB.LSet(ctx, l.Key, l.Index, "").Err()

	err = RedisDB.LRem(ctx, l.Key, 1, "").Err()

	return
}

func (l *ListData) SetListOne(ctx context.Context, RedisDB *redis.Client, oldValue any, newValue any) (err error) {

	_, err = RedisDB.LInsert(ctx, l.Key, "after", oldValue, newValue).Result()
	if err != nil {
		return err
	}
	return nil
}
