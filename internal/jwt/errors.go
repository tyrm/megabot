package jwt

import "errors"

var (
	errRefreshExpired      = errors.New("refresh expired")
	errUnauthorized        = errors.New("unauthorized")
	errUnprocessableEntity = errors.New("unprocessable entity")
)
