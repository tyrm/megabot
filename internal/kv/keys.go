package kv

const (
	keyBase = "megabot:"

	keyJwt          = keyBase + "jwt:"
	keyJwtAccesses  = keyJwt + "a:"
	keyJwtRefreshes = keyJwt + "r:"

	keySession = keyBase + "session:"
)

// KeySession returns the base kv key prefix
func KeySession() string { return keySession }

// KeyJwtAccess returns the kv key which holds a JWT access token
func KeyJwtAccess(d string) string { return keyJwtAccesses + d }

// KeyJwtRefresh returns the kv key which holds a JWT refresh token
func KeyJwtRefresh(d string) string { return keyJwtRefreshes + d }
