package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
)

// Client represents a redis client
type Client struct {
	redis *redis.Client
}

// New creates a new redis client.
func New(ctx context.Context) (*Client, error) {
	l := logger.WithField("func", "New")

	c := Client{
		redis: redis.NewClient(&redis.Options{
			Addr:     viper.GetString(config.Keys.RedisAddress),
			Password: viper.GetString(config.Keys.RedisPassword),
			DB:       viper.GetInt(config.Keys.RedisDB),
		}),
	}

	resp := c.redis.Ping(ctx)
	l.Debugf("%s", resp.String())

	return &c, nil
}

// Close closes the redis pool
func (c *Client) Close(ctx context.Context) error {
	return c.redis.Close()
}

// RedisClient returns the redis client
func (c *Client) RedisClient() *redis.Client {
	return c.redis
}
