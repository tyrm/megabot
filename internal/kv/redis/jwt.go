package redis

import (
	"context"
	"github.com/tyrm/megabot/internal/kv"
	"time"
)

// DeleteJWTAccessToken deletes an access token from redis.
func (c *client) DeleteJWTAccessToken(ctx context.Context, accessTokenID string) error {
	_, err := c.client.Del(ctx, kv.KeyJwtAccess(accessTokenID)).Result()
	if err != nil {
		return err
	}

	return nil
}

// DeleteJWTRefreshToken deletes a refresh token from redis.
func (c *client) DeleteJWTRefreshToken(ctx context.Context, refreshTokenID string) error {
	_, err := c.client.Del(ctx, kv.KeyJwtRefresh(refreshTokenID)).Result()
	if err != nil {
		return err
	}

	return nil
}

// GetJWTAccessToken retrieves an access token from redis.
func (c *client) GetJWTAccessToken(ctx context.Context, accessTokenID string) (string, error) {
	resp, err := c.client.Get(ctx, kv.KeyJwtAccess(accessTokenID)).Result()
	if err != nil {
		return "", err
	}

	return resp, nil
}

// GetJWTRefreshToken retrieves an refresh token from redis.
func (c *client) GetJWTRefreshToken(ctx context.Context, refreshTokenID string) (string, error) {
	resp, err := c.client.Get(ctx, kv.KeyJwtRefresh(refreshTokenID)).Result()
	if err != nil {
		return "", err
	}

	return resp, nil
}

// SetJWTAccessToken adds an access token to redis.
func (c *client) SetJWTAccessToken(ctx context.Context, accessTokenID, userID string, expire time.Duration) error {
	_, err := c.client.SetEX(ctx, kv.KeyJwtAccess(accessTokenID), userID, expire).Result()
	if err != nil {
		return err
	}

	return nil
}

// SetJWTRefreshToken adds a refresh token to redis.
func (c *client) SetJWTRefreshToken(ctx context.Context, refreshTokenID string, userID string, expire time.Duration) error {
	_, err := c.client.SetEX(ctx, kv.KeyJwtRefresh(refreshTokenID), userID, expire).Result()
	if err != nil {
		return err
	}

	return nil
}
