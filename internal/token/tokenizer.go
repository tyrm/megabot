package token

import (
	"github.com/speps/go-hashids/v2"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
)

// Tokenizer generates public tokens for database IDs to obfuscate the database IDs
type Tokenizer struct {
	h *hashids.HashID
}

// DecodeToken returns the kind and id number of a provided token
func (t *Tokenizer) DecodeToken(token string) (Kind, int64, error) {
	d, err := t.h.DecodeInt64WithError(token)
	if err != nil {
		return 0, 0, err
	}
	if len(d) != 2 {
		return 0, 0, ErrInvalidLength
	}
	return Kind(d[0]), d[1], nil
}

// EncodeToken turns a model kind and id into a token
func (t *Tokenizer) EncodeToken(kind Kind, id int64) (string, error) {
	return t.h.EncodeInt64([]int64{int64(kind), id})
}

// New returns a new tokenizer
func New() (*Tokenizer, error) {
	// set config
	hd := hashids.NewData()
	salt := viper.GetString(config.Keys.TokenSalt)
	if salt == "" {
		return nil, ErrSaltEmpty
	}
	hd.Salt = salt
	hd.MinLength = 16

	// create hashid
	hid, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}

	return &Tokenizer{
		h: hid,
	}, nil
}
