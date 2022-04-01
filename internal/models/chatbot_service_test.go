package models

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tmthrgd/go-hex"
	"github.com/tyrm/megabot/internal/chatbot"
	"github.com/tyrm/megabot/internal/config"
	"github.com/uptrace/bun"
	"testing"
	"time"
)

func TestChatbotService_BeforeAppendModel_Insert(t *testing.T) {
	obj := &ChatbotService{
		Description: "test 1",
		ServiceType: chatbot.ServiceTelegram,
	}

	err := obj.BeforeAppendModel(context.Background(), &bun.InsertQuery{})
	if err != nil {
		t.Errorf("got error: %s", err.Error())
		return
	}

	emptyTime := time.Time{}
	if obj.CreatedAt == emptyTime {
		t.Errorf("invalid created at time: %s", obj.CreatedAt.String())
	}
	if obj.UpdatedAt == emptyTime {
		t.Errorf("invalid updated at time: %s", obj.UpdatedAt.String())
	}
}

func TestChatbotService_BeforeAppendModel_Update(t *testing.T) {
	obj := &ChatbotService{
		Description: "test 1",
		ServiceType: chatbot.ServiceTelegram,
	}

	err := obj.BeforeAppendModel(context.Background(), &bun.UpdateQuery{})
	if err != nil {
		t.Errorf("got error: %s", err.Error())
		return
	}

	emptyTime := time.Time{}
	if obj.UpdatedAt == emptyTime {
		t.Errorf("invalid updated at time: %s", obj.UpdatedAt.String())
	}
}

func TestChatbotService_GetConfig(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	obj := ChatbotService{
		Config: hex.MustDecodeString("43dc49ab017fbde685011bc75e7aeecf46e2e6ca2d960681ebca6b7d9b5a74ad0348cfcadbdb71bebb"),
	}

	val, err := obj.GetConfig()
	if err != nil {
		t.Errorf("unexpected error getting scanning, got: '%s', want: 'nil", err)
		return
	}
	if string(val) != "test string 1" {
		t.Errorf("unexpected value, got: '%s', want: '%s'", val, "test string 1")
	}
}

func TestChatbotService_SetConfig(t *testing.T) {

	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	tables := []struct {
		n string
	}{
		{"test string 1"},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Getting id", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var obj ChatbotService

			err := obj.SetConfig(table.n)
			if err != nil {
				t.Errorf("unexpected error getting value: %s", err.Error())
				return
			}

			gcm, err := getCrypto()
			if err != nil {
				t.Errorf("getting crypto: %s", err.Error())
				return
			}

			nonceSize := gcm.NonceSize()
			if len(obj.Config) < nonceSize {
				t.Errorf("value too small")
				return
			}

			nonce, ciphertext := obj.Config[:nonceSize], obj.Config[nonceSize:]
			plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
			if err != nil {
				t.Errorf("decrypting: %s", err.Error())
				return
			}
			if string(plaintext) != table.n {
				t.Errorf("unexpected value, got: '%s', want: '%s'", plaintext, table.n)
			}
		})
	}
}
