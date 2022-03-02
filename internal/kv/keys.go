package kv

var (
	keyBase = "megabot:"

	keyJwt = keyBase + "jwt:"

	keyJwtAccesses  = keyJwt + "a:"
	keyJwtRefreshes = keyJwt + "r:"
)

// KeyJwtAccess returns the kv key which holds a JWT access token
func KeyJwtAccess(d string) string { return keyJwtAccesses + d }

// KeyJwtRefresh returns the kv key which holds a JWT refresh token
func KeyJwtRefresh(d string) string { return keyJwtRefreshes + d }
