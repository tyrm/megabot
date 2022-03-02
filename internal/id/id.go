package id

import (
	"crypto/rand"
	"github.com/oklog/ulid"
	"math/big"
	"time"
)

const randomRange = 630720000

// NewULID returns a new ULID string.
func NewULID() (string, error) {
	newUlid, err := ulid.New(ulid.Timestamp(time.Now()), rand.Reader)
	if err != nil {
		return "", err
	}
	return newUlid.String(), nil
}

// NewRandomULID returns a new ULID string using a random time.
func NewRandomULID() (string, error) {
	b1, err := rand.Int(rand.Reader, big.NewInt(randomRange))
	if err != nil {
		return "", err
	}
	r1 := time.Duration(int(b1.Int64()))

	b2, err := rand.Int(rand.Reader, big.NewInt(randomRange))
	if err != nil {
		return "", err
	}
	r2 := -time.Duration(int(b2.Int64()))

	arbitraryTime := time.Now().Add(r1 * time.Second).Add(r2 * time.Second)
	newUlid, err := ulid.New(ulid.Timestamp(arbitraryTime), rand.Reader)
	if err != nil {
		return "", err
	}
	return newUlid.String(), nil
}
