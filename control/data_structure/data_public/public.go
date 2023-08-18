package data_public

import (
	"changeme/utils"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
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

func AnyCmd(ctx context.Context, client *redis.Client, cmd string) (r any, err error) {
	tmpMap := make(map[int]string, 10)
	tmpList := strings.Split(cmd, " ")
	for i, v := range tmpList {
		tmpMap[i] = v
	}
	tmpLength := len(tmpList)

	switch tmpLength {
	case 1:
		r, err = client.Do(ctx, tmpMap[0]).Result()
	case 2:
		r, err = client.Do(ctx, tmpMap[0], tmpMap[1]).Result()
	case 3:
		r, err = client.Do(ctx, tmpMap[0], tmpMap[1], tmpMap[2]).Result()
	case 4:
		r, err = client.Do(ctx, tmpMap[0], tmpMap[1], tmpMap[2], tmpMap[3]).Result()
	case 5:
		r, err = client.Do(ctx, tmpMap[0], tmpMap[1], tmpMap[2], tmpMap[3], tmpMap[4]).Result()
	case 6:
		r, err = client.Do(ctx, tmpMap[0], tmpMap[1], tmpMap[2], tmpMap[3], tmpMap[4], tmpMap[5]).Result()
	case 7:
		r, err = client.Do(ctx, tmpMap[0], tmpMap[1], tmpMap[2], tmpMap[3], tmpMap[4], tmpMap[5], tmpMap[6]).Result()
	case 8:
		r, err = client.Do(ctx, tmpMap[0], tmpMap[1], tmpMap[2], tmpMap[3], tmpMap[4], tmpMap[5], tmpMap[6], tmpMap[7]).Result()
	case 9:
		r, err = client.Do(ctx, tmpMap[0], tmpMap[1], tmpMap[2], tmpMap[3], tmpMap[4], tmpMap[5], tmpMap[6], tmpMap[7], tmpMap[8]).Result()
	case 10:
		r, err = client.Do(ctx, tmpMap[0], tmpMap[1], tmpMap[2], tmpMap[3], tmpMap[4], tmpMap[5], tmpMap[6], tmpMap[7], tmpMap[8], tmpMap[9]).Result()
	}
	fmt.Println("r", r, "err", err)
	return
}

func GetDBList(ctx context.Context, client *redis.Client) (any, error) {
	//r, err := client.Do(ctx, "CONFIG", "GET", "databases").Result()
	result, err := client.ConfigGet(ctx, "databases").Result()
	if err != nil {
		return nil, err
	}
	tmp, _ := strconv.Atoi(result["databases"])
	type tmpStruct struct {
		DBName   string
		DBNumber int
	}
	tmpMap := make([]tmpStruct, tmp)
	for i := range tmpMap {
		tmpMap[i].DBName = "db" + strconv.Itoa(i)
		tmpMap[i].DBNumber = i
	}
	return tmpMap, err
}

func ChangeDB(ctx context.Context, client *redis.Client, DB int) error {
	return client.Do(ctx, "SELECT", DB).Err()
}
