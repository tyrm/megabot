package graphql

import "errors"

var (
	errBadLogin       = errors.New("email/password combo invalid")
	errRefreshExpired = errors.New("refresh expired")
	errUnauthorized   = errors.New("unauthorized")
)
