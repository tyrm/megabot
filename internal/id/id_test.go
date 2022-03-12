package id

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestNewULID(t *testing.T) {
	for i := 1; i <= 5; i++ {
		i := i
		validate := validator.New()

		name := fmt.Sprintf("[%d] Starting", i)
		t.Run(name, func(t *testing.T) {
			id, err := NewULID()
			if err != nil {
				t.Errorf("[%d] got error generating id: %s", i, err.Error())
				return
			}

			err = validate.Var(id, "required,ulid")
			if err != nil {
				t.Errorf("[%d] %s is not a valid ulid: %s", i, id, err.Error())
				return
			}
		})
	}
}

func TestNewRandomULID(t *testing.T) {
	for i := 1; i <= 5; i++ {
		i := i
		validate := validator.New()

		name := fmt.Sprintf("[%d] Starting", i)
		t.Run(name, func(t *testing.T) {
			id, err := NewRandomULID()
			if err != nil {
				t.Errorf("[%d] got error generating id: %s", i, err.Error())
				return
			}

			err = validate.Var(id, "required,ulid")
			if err != nil {
				t.Errorf("[%d] %s is not a valid ulid: %s", i, id, err.Error())
				return
			}
		})
	}
}

func BenchmarkNewULID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := NewULID()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkNewRandomULID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := NewRandomULID()
		if err != nil {
			panic(err)
		}
	}
}
