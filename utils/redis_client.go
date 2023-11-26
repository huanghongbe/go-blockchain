package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址和端口
		Password: "",               // Redis 服务器密码（如果有的话）
		DB:       0,                // Redis 数据库索引
	})

	return &RedisClient{
		client: rdb,
	}
}

func (rc *RedisClient) GetValueByKey(key string) (string, error) {
	value, err := rc.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

func (rc *RedisClient) GetIntValueByKey(key string) (int, error) {
	value, err := rc.client.Get(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("转换失败:", err)
		return 0, err
	}

	return parsedValue, nil
}

func (rc *RedisClient) SetValue(key string, value string) error {
	err := rc.client.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rc *RedisClient) SetIntValue(key string, value int) error {
	err := rc.client.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
