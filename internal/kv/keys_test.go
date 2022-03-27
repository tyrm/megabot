package kv

import "testing"

func TestKeyJwtAccess(t *testing.T) {
	v := KeyJwtAccess("test123")
	if v != "megabot:jwt:a:test123" {
		t.Errorf("enexpected value for KeyDomains, got: '%s', want: 'megabot:jwt:a:test123'.", v)
	}
}

func TestKeyJwtRefresh(t *testing.T) {
	v := KeyJwtRefresh("test123")
	if v != "megabot:jwt:r:test123" {
		t.Errorf("enexpected value for KeyDomains, got: '%s', want: 'megabot:jwt:r:test123'.", v)
	}
}

func TestKeySession(t *testing.T) {
	v := KeySession()
	if v != "megabot:session:" {
		t.Errorf("enexpected value for KeyDomains, got: '%s', want: 'megabot:session:'.", v)
	}
}
