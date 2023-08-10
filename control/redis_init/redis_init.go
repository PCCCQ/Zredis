package redis_init

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	// redis服务器地址，ip:port格式，比如：192.168.1.100:6379
	// 默认为 :6379
	IP   string `json:"IP,omitempty"`
	Port string `json:"port,omitempty"`

	// 当redis服务器版本在6.0以上时，作为ACL认证信息配合密码一起使用，
	// ACL是redis 6.0以上版本提供的认证功能，6.0以下版本仅支持密码认证。
	// 默认为空，不进行认证。
	Username string `json:"username,omitempty"`

	// 当redis服务器版本在6.0以上时，作为ACL认证信息配合密码一起使用，
	// 当redis服务器版本在6.0以下时，仅作为密码认证。
	// ACL是redis 6.0以上版本提供的认证功能，6.0以下版本仅支持密码认证。
	// 默认为空，不进行认证。
	Password string `json:"password,omitempty"`

	// redis DB 数据库，默认为0
	DB int `json:"DB,omitempty"`
}

func (r *RedisClient) RedisInit() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.IP + ":" + r.Port,
		Username: r.Username,
		Password: r.Password,
		DB:       r.DB,
	})
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		return rdb, err
	}
	return rdb, nil
}

func AllKeyTypes(ctx context.Context, RedisDB *redis.Client) (data any, err error) {
	type p struct {
		Type string `json:"type,omitempty"`
		Key  string `json:"key,omitempty"`
	}
	// 获取所有匹配模式"*"的键
	keys, err := RedisDB.Keys(ctx, "*").Result()
	keyTypes := make([]p, len(keys))
	for i, key := range keys {
		// 获取每个键的数据类型
		keyTypes[i].Key = key
		keyTypes[i].Type = RedisDB.Type(ctx, key).Val()
	}
	return keyTypes, err
}
