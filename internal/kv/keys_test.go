package kv

import "testing"

func TestKeyJwtAccess(t *testing.T) {
	v := KeyJwtAccess("test123")
	if v != "megabot:jwt:a:test123" {
		t.Errorf("enexpected value for KeyDomains, got: '%s', want: 'sr:jwt:a:test123'.", v)
	}
}

func TestKeyJwtRefresh(t *testing.T) {
	v := KeyJwtRefresh("test123")
	if v != "megabot:jwt:r:test123" {
		t.Errorf("enexpected value for KeyDomains, got: '%s', want: 'sr:jwt:r:test123'.", v)
	}
}
