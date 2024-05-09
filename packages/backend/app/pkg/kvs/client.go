package kvs

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/stlatica/stlatica/packages/backend/app/pkg/logger"
)

// Client is an interface for key-value store.
type Client interface {
	// SetValue sets a key value pair in valkey.
	SetValue(ctx context.Context, key string, value string) error
	// GetValue retrieves a value for a given key from valkey.
	GetValue(ctx context.Context, key string) (string, error)
}

type client struct {
	client    *redis.Client
	appLogger *logger.AppLogger
}

// NewClient creates a new key-value store client.
func NewClient(ctx context.Context, appLogger *logger.AppLogger, addr string,
	password string, defaultDB int) Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,     // address of kvs server
		Password: password, // password of kvs
		DB:       defaultDB,
	})

	// pingを飛ばす
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Err: Could not connect to Redis: %v", err)
		return nil
	}

	return &client{
		client:    rdb,
		appLogger: appLogger,
	}
}

// SetValue sets a key value pair in key-value store
func (c *client) SetValue(ctx context.Context, key string, value string) error {
	err := c.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set value: %w", err)
	}
	return nil
}

// GetValue retrieves a value for a given key from key-value store.
func (c *client) GetValue(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get value: %w", err)
	}
	return val, nil
}
