package redis

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

// 全局变量
var (
	rdb *redis.Client
)

// 连接redis
func ConnectionRedis() (*redis.Client, error) {
	dsn := os.Getenv("REDIS")
	if dsn == "" {
		return nil, fmt.Errorf("MYSQL: 环境变量未设置")
	}
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}

	rdb = redis.NewClient(opt)

	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Redis 连接失败: %v", err)
	} else {
		fmt.Println("Redis 连接成功")
	}

	return rdb, nil
}

func GetDB() *redis.Client {
	return rdb
}
