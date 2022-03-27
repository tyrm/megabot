package models

import (
	"crypto/aes"
	gocipher "crypto/cipher"
	"crypto/rand"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
	"io"
	"strings"
)

type EncryptedString string

func (e *EncryptedString) Scan(src interface{}) error {
	l := logger.WithField("func", "Scan").WithField("type", "EncryptedString")

	var data []byte
	switch src := src.(type) {
	case nil:
		return nil
	case string:
		if src == "" {
			return nil
		}
		data = []byte(src)
	case []byte:
		if len(src) == 0 {
			return nil
		}
		data = src
	default:
		msg := fmt.Sprintf("unable to scan type %T into EncryptedString", src)
		l.Error(msg)
		return errors.New(msg)
	}

	gcm, err := getCrypto()
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		msg := "data too small"
		l.Error(msg)
		return errors.New(msg)
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		l.Errorf("decrypting: %s", err.Error())
		return err
	}

	*e = EncryptedString(plaintext)
	return nil
}

func (e *EncryptedString) Value() (driver.Value, error) {
	l := logger.WithField("func", "Value").WithField("type", "EncryptedString")

	gcm, err := getCrypto()
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		l.Errorf("reading nonce: %s", err.Error())
		return nil, err
	}

	return gcm.Seal(nonce, nonce, []byte(*e), nil), nil
}

func getCrypto() (gocipher.AEAD, error) {
	l := logger.WithField("func", "getCrypto").WithField("type", "EncryptedString")

	key := []byte(strings.ToLower(viper.GetString(config.Keys.DbEncryptionKey)))
	cipher, err := aes.NewCipher(key)
	if err != nil {
		l.Errorf("new cipher: %s", err.Error())
		return nil, err
	}

	gcm, err := gocipher.NewGCM(cipher)
	if err != nil {
		l.Errorf("new gcm: %s", err.Error())
		return nil, err
	}

	return gcm, nil
}
