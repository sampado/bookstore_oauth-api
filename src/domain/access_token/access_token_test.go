package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Error("expiration time should be 24 hours")
	}
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken(1)

	assert.False(t, at.IsExpired(), "brand new access token should not be expired")

	assert.Equal(t, "", at.AccessToken, "new access token should not have defined access token id")

	assert.EqualValues(t, 0, at.UserId, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	if !at.IsExpired() {
		t.Error("emptty access tiken should be expired by default")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("access token expiring 3 hours from now should NOT be expired")
	}

	at.Expires = time.Now().UTC().Add(-3 * time.Hour).Unix()
	if !at.IsExpired() {
		t.Error("access token expired 3 hours ago should be expired")
	}
}
