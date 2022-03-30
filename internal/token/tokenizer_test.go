package token

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tyrm/megabot/internal/config"
	"testing"
)

var testTables = []struct {
	k Kind
	i int64
	t string
}{
	{KindUser, 123, "AMroVP59cwLPE5pb"},
	{KindGroupMembership, 84685, "wM9xPkbiDrkX7ve0"},
	{KindChatbotService, 1, "pMeLrPDzIxagV8KY"},
}

func TestNew(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.TokenSalt, "test1234")

	tokenizer, err := New()
	if err != nil {
		t.Errorf("got error: %s", err.Error())
		return
	}

	if tokenizer.h == nil {
		t.Errorf("hashid is nil")
		return
	}
}

func TestNew_SaltEmpty(t *testing.T) {
	viper.Reset()

	tokenizer, err := New()
	if err != ErrSaltEmpty {
		t.Errorf("unexpected error, got: '%s', want: '%s'", err, ErrSaltEmpty)
		return
	}

	if tokenizer != nil {
		t.Errorf("unexpected tokenizer, got: '%T', want: '%T'", tokenizer, nil)
		return
	}
}

func TestTokenizer_DecodeToken(t *testing.T) {
	tokenizer, err := testNewTestTokenizer()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	for i, table := range testTables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running EncodeToken %d(%s)", i, table.i, table.k.String())
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			kind, id, err := tokenizer.DecodeToken(table.t)
			if err != nil {
				t.Errorf("got error: %s", err.Error())
				return
			}
			if kind != table.k {
				t.Errorf("[%d] wrong kind: got '%s', want '%s'", i, kind, table.k)
			}
			if id != table.i {
				t.Errorf("[%d] wrong id: got '%d', want '%d'", i, id, table.i)
			}
		})
	}
}

func TestTokenizer_DecodeToken_InvalidLength(t *testing.T) {
	tokenizer, err := testNewTestTokenizer()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	_, _, err = tokenizer.DecodeToken("1vxqadgcYibQ2pOj")
	errText := "negative number not supported"
	if err != ErrInvalidLength {
		t.Errorf("unexpected error, got: '%s', want: '%s'", err, errText)
		return
	}
}

func TestTokenizer_EncodeToken(t *testing.T) {
	tokenizer, err := testNewTestTokenizer()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	for i, table := range testTables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running EncodeToken %d(%s)", i, table.i, table.k.String())
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			token, err := tokenizer.EncodeToken(table.k, table.i)
			if err != nil {
				t.Errorf("got error: %s", err.Error())
				return
			}
			if token != table.t {
				t.Errorf("[%d] wrong token: got '%s', want '%s'", i, token, table.t)
			}
		})
	}
}

func TestTokenizer_EncodeToken_Negative(t *testing.T) {
	tokenizer, err := testNewTestTokenizer()
	if err != nil {
		t.Errorf("init: %s", err.Error())
		return
	}

	_, err = tokenizer.EncodeToken(KindUser, -1)
	errText := "negative number not supported"
	if err == nil {
		t.Errorf("expected error, got: 'nil', want: '%s'", errText)
		return
	}
	if err.Error() != errText {
		t.Errorf("unexpected error, got: '%s', want: '%s'", err, errText)
		return
	}
}

func testNewTestTokenizer() (*Tokenizer, error) {
	viper.Reset()
	viper.Set(config.Keys.TokenSalt, "test1234")
	return New()
}
