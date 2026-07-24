package redis

import (
	"fmt"
	"gin-scaffold/internal/config"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	DB *redis.Client
}

func NewClient(cfg *config.RedisConfig) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConn,
	})
	return &Client{DB: client}
}

func (c *Client) Close() error {
	return c.DB.Close()
}