package kv

import (
	"context"
	"time"
)

type KV interface {
	JWT
	Close(ctx context.Context) error
}

type JWT interface {
	DeleteJWTAccessToken(ctx context.Context, accessTokenID string) error
	DeleteJWTRefreshToken(ctx context.Context, refreshTokenID string) error
	GetJWTAccessToken(ctx context.Context, accessTokenID string) (string, error)
	GetJWTRefreshToken(ctx context.Context, refreshTokenID string) (string, error)
	SetJWTAccessToken(ctx context.Context, accessTokenID, userID string, expire time.Duration) error
	SetJWTRefreshToken(ctx context.Context, refreshTokenID, userID string, expire time.Duration) error
}
