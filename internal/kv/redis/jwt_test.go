//go:build redis

package redis

import (
	"context"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
	"github.com/tyrm/megabot/internal/kv"
	"strconv"
	"testing"
	"time"
)

func TestClient_DeleteJWTAccessToken(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.RedisAddress, "redis:6379")
	viper.Set(config.Keys.RedisPassword, "test")

	client, err := New(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing redis connection: %s", err.Error())
		return
	}

	_, err = client.redis.Set(context.Background(), kv.KeyJwtAccess("test"), "01FY8C6NW8BJDX2YMNJ8FBFCD2", 0).Result()
	if err != nil {
		t.Logf("error preping test: %s", err.Error())
		return
	}

	err = client.DeleteJWTAccessToken(context.Background(), "test")
	if err != nil {
		t.Logf("unexpected error running DeleteJWTAccessToken: %s", err.Error())
		return
	}

	resp, err := client.redis.Get(context.Background(), kv.KeyJwtAccess("test")).Result()
	if err != nil {
		t.Logf("error checking test: %s", err.Error())
		return
	}
	if resp != "" {
		t.Logf("unexpected data in redis: %s", resp)
		return
	}
}

func TestClient_DeleteJWTRefreshToken(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.RedisAddress, "redis:6379")
	viper.Set(config.Keys.RedisPassword, "test")

	client, err := New(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing redis connection: %s", err.Error())
		return
	}

	_, err = client.redis.Set(context.Background(), kv.KeyJwtRefresh("test"), "01FY8C6NW8BJDX2YMNJ8FBFCD2", 0).Result()
	if err != nil {
		t.Logf("error preping test: %s", err.Error())
		return
	}

	err = client.DeleteJWTRefreshToken(context.Background(), "test")
	if err != nil {
		t.Logf("unexpected error running DeleteJWTAccessToken: %s", err.Error())
		return
	}

	resp, err := client.redis.Get(context.Background(), kv.KeyJwtRefresh("test")).Result()
	if err != nil {
		t.Logf("error checking test: %s", err.Error())
		return
	}
	if resp != "" {
		t.Logf("unexpected data in redis: %s", resp)
		return
	}
}

func TestClient_GetJWTAccessToken(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.RedisAddress, "redis:6379")
	viper.Set(config.Keys.RedisPassword, "test")

	client, err := New(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing redis connection: %s", err.Error())
		return
	}

	testID := "01FY8C6NW8BJDX2YMNJ8FBFCD2"

	_, err = client.redis.Set(context.Background(), kv.KeyJwtAccess("test"), testID, 0).Result()
	if err != nil {
		t.Logf("error preping test: %s", err.Error())
		return
	}

	resp, err := client.GetJWTAccessToken(context.Background(), "test")
	if err != nil {
		t.Logf("unexpected error running GetJWTAccessToken: %s", err.Error())
		return
	}
	if resp != testID {
		t.Logf("unexpected data in redis, got: %s, want: %s", resp, testID)
		return
	}
}

func TestClient_GetJWTRefreshToken(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.RedisAddress, "redis:6379")
	viper.Set(config.Keys.RedisPassword, "test")

	client, err := New(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing redis connection: %s", err.Error())
		return
	}

	testID := "01FY8C6NW8BJDX2YMNJ8FBFCD2"

	_, err = client.redis.Set(context.Background(), kv.KeyJwtRefresh("test"), testID, 0).Result()
	if err != nil {
		t.Logf("error preping test: %s", err.Error())
		return
	}

	resp, err := client.GetJWTRefreshToken(context.Background(), "test")
	if err != nil {
		t.Logf("unexpected error running GetJWTRefreshToken: %s", err.Error())
		return
	}
	if resp != testID {
		t.Logf("unexpected data in redis, got: %s, want: %s", resp, testID)
		return
	}
}

func TestClient_SetJWTAccessToken(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.RedisAddress, "redis:6379")
	viper.Set(config.Keys.RedisPassword, "test")

	client, err := New(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing redis connection: %s", err.Error())
		return
	}

	var testID int64 = 9999

	err = client.SetJWTAccessToken(context.Background(), "test", testID, 1*time.Second)
	if err != nil {
		t.Logf("unexpected error running SetJWTAccessToken: %s", err.Error())
		return
	}

	resp, err := client.redis.Get(context.Background(), kv.KeyJwtAccess("test")).Result()
	if err != nil {
		t.Logf("error checking test: %s", err.Error())
		return
	}
	respInt, err := strconv.ParseInt(resp, 10, 64)
	if err != nil {
		t.Logf("error converting to int64: %s", err.Error())
		return
	}
	if respInt != testID {
		t.Logf("unexpected data in redis, got: %d, want: %d", respInt, testID)
		return
	}

	time.Sleep(2 * time.Second)

	resp, err = client.redis.Get(context.Background(), kv.KeyJwtAccess("test")).Result()
	if err != nil {
		t.Logf("error checking test: %s", err.Error())
		return
	}
	if resp != "" {
		t.Logf("unexpected data in redis, got: %s, want: nil", resp)
		return
	}
}

func TestClient_SetJWTRefreshToken(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.RedisAddress, "redis:6379")
	viper.Set(config.Keys.RedisPassword, "test")

	client, err := New(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing redis connection: %s", err.Error())
		return
	}

	var testID int64 = 9999

	err = client.SetJWTRefreshToken(context.Background(), "test", testID, 1*time.Second)
	if err != nil {
		t.Logf("unexpected error running SetJWTRefreshToken: %s", err.Error())
		return
	}

	resp, err := client.redis.Get(context.Background(), kv.KeyJwtAccess("test")).Result()
	if err != nil {
		t.Logf("error checking test: %s", err.Error())
		return
	}
	respInt, err := strconv.ParseInt(resp, 10, 64)
	if err != nil {
		t.Logf("error converting to int64: %s", err.Error())
		return
	}
	if respInt != testID {
		t.Logf("unexpected data in redis, got: %d, want: %d", respInt, testID)
		return
	}

	time.Sleep(2 * time.Second)

	resp, err = client.redis.Get(context.Background(), kv.KeyJwtAccess("test")).Result()
	if err != nil {
		t.Logf("error checking test: %s", err.Error())
		return
	}
	if resp != "" {
		t.Logf("unexpected data in redis, got: %s, want: nil", resp)
		return
	}
}
