/*
 * @FilePath: \driver\redis.go
 * @Descripttion:
 *
 * @Date: 2020-07-29 13:59:01
 * @Author: yuanhao
 *
 */
package driver

import (
	"context"
	"nssa/config"
	"nssa/log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb       *redis.Client
	redisOnce sync.Once
	ctx       = context.Background()
)

// 获取redis操作对象
func GetRedisDB() *redis.Client {

	redisOnce.Do(func() {
		conf := config.GetConfigure()
		rdb = redis.NewClient(&redis.Options{
			Addr:     conf.Redis.Host,
			Password: conf.Redis.Password,
			DB:       0,
		})

		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			log.Fatal("fail to connect redis", err)
		}
	})

	return rdb
}

// 存储redis数据
func SetRedisValue(key, value string, exp time.Duration) error {
	return rdb.Set(ctx, key, value, exp).Err()
}

// 获取redis数据
func GetRedisValue(key string) (string, error) {

	return rdb.Get(ctx, key).Result()
}

func DelRedisValue(key string) {
	rdb.Del(ctx, key)
}
