package graphql

import "errors"

var (
	errBadLogin = errors.New("email/password combo invalid")
)
