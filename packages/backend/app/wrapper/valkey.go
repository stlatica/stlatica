package wrapper

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

// RedisClient is a client of Redis.
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient creates a new instance of RedisClient.
func NewRedisClient(ctx context.Context) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("STLATICA_KVS_ADDRESS"),  // address of kvs server
		Password: os.Getenv("STLATICA_KVS_PASSWORD"), // password of kvs
		DB:       0,
	})

	// pingを飛ばす
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Err: Could not connect to Redis: %v", err)
		return nil
	}

	return &RedisClient{
		client: rdb,
	}
}

// SetValue sets a key value pair in Redis.
func (rc *RedisClient) SetValue(ctx context.Context, key string, value string) error {
	err := rc.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set value: %w", err)
	}
	return nil
}

// GetValue retrieves a value for a given key from Redis.
func (rc *RedisClient) GetValue(ctx context.Context, key string) (string, error) {
	val, err := rc.client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get value: %w", err)
	}
	return val, nil
}
