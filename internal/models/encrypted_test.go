package models

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tmthrgd/go-hex"
	"github.com/tyrm/megabot/internal/config"
	"testing"
)

func TestEncryptedString_Scan_Bytes(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	var estring EncryptedString

	byts := hex.MustDecodeString("43dc49ab017fbde685011bc75e7aeecf46e2e6ca2d960681ebca6b7d9b5a74ad0348cfcadbdb71bebb")

	err := estring.Scan(byts)
	if err != nil {
		t.Errorf("unexpected error getting scanning, got: '%s', want: 'nil", err)
		return
	}
	if estring != "test string 1" {
		t.Errorf("unexpected value, got: '%s', want: '%s'", estring, "test string 1")
	}
}

func TestEncryptedString_Scan_BytesEmpty(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	var estring EncryptedString

	var byts []byte

	err := estring.Scan(byts)
	if err != nil {
		t.Errorf("unexpected error getting scanning, got: '%s', want: 'nil", err)
		return
	}
	if estring != "" {
		t.Errorf("unexpected value, got: '%s', want: '%s'", estring, "")
	}
}

func TestEncryptedString_Scan_Int(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	var estring EncryptedString

	errMsg := "unable to scan type int into EncryptedString"
	err := estring.Scan(1)
	if err.Error() != errMsg {
		t.Errorf("unexpected error getting scanning, got: '%s', want: '%s", err, errMsg)
		return
	}
	if estring != "" {
		t.Errorf("unexpected value, got: '%s', want: '%s'", estring, "")
	}
}

func TestEncryptedString_Scan_Nil(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	var estring EncryptedString

	err := estring.Scan(nil)
	if err != nil {
		t.Errorf("unexpected error getting scanning, got: '%s', want: 'nil", err)
		return
	}
	if estring != "" {
		t.Errorf("unexpected value, got: '%s', want: '%s'", estring, "")
	}
}

func TestEncryptedString_Scan_String(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	var estring EncryptedString

	str := string(hex.MustDecodeString("43dc49ab017fbde685011bc75e7aeecf46e2e6ca2d960681ebca6b7d9b5a74ad0348cfcadbdb71bebb"))

	err := estring.Scan(str)
	if err != nil {
		t.Errorf("unexpected error getting scanning, got: '%s', want: 'nil", err)
		return
	}
	if estring != "test string 1" {
		t.Errorf("unexpected value, got: '%s', want: '%s'", estring, "test string 1")
	}
}

func TestEncryptedString_Scan_StringEmpty(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	var estring EncryptedString

	err := estring.Scan("")
	if err != nil {
		t.Errorf("unexpected error getting scanning, got: '%s', want: 'nil", err)
		return
	}
	if estring != "" {
		t.Errorf("unexpected value, got: '%s', want: '%s'", estring, "")
	}
}

func TestEncryptedString_String(t *testing.T) {
	var e EncryptedString = "test string 1"

	if e.String() != "test string 1" {
		t.Errorf("unexpected value, got: '%s', want: '%s'", e.String(), "test string 1")
	}
}

func TestEncryptedString_Value(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbEncryptionKey, "0123456789012345")

	tables := []struct {
		n EncryptedString
	}{
		{"test string 1"},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Getting id", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			val, err := table.n.Value()
			if err != nil {
				t.Errorf("unexpected error getting value: %s", err.Error())
				return
			}

			gcm, err := getCrypto()
			if err != nil {
				t.Errorf("getting crypto: %s", err.Error())
				return
			}

			data := val.([]byte)
			nonceSize := gcm.NonceSize()
			if len(data) < nonceSize {
				t.Errorf("value too small")
				return
			}

			nonce, ciphertext := data[:nonceSize], data[nonceSize:]
			plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
			if err != nil {
				t.Errorf("decrypting: %s", err.Error())
				return
			}
			if string(plaintext) != string(table.n) {
				t.Errorf("unexpected value, got: '%s', want: '%s'", plaintext, table.n)
			}
		})
	}
}
