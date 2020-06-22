package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/sampado/bookstore_utils-go/rest_errors"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// Used for password grand type
	UserName string `json:"username"`
	Password string `json:"password"`

	// Used for client_credentials grand type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() rest_errors.RestError {
	switch at.GrantType {
	case grantTypePassword:
		if at.UserName == "" {
			return rest_errors.NewBadRequestError("invalid UserName id")
		}
		if at.Password == "" {
			return rest_errors.NewBadRequestError("invalid Password id")
		}
	case grantTypeClientCredentials:
		if at.ClientId == "" {
			return rest_errors.NewBadRequestError("invalid ClientId id")
		}
		if at.ClientSecret == "" {
			return rest_errors.NewBadRequestError("invalid ClientSecret id")
		}
	default:
		return rest_errors.NewBadRequestError("invalid access token grant_type")
	}

	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() rest_errors.RestError {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return rest_errors.NewBadRequestError("invalid access token id")
	}
	if at.ClientId <= 0 {
		return rest_errors.NewBadRequestError("invalid client id")
	}
	if at.UserId <= 0 {
		return rest_errors.NewBadRequestError("invalid user id")
	}
	if at.Expires <= 0 {
		return rest_errors.NewBadRequestError("invalid expiration time id")
	}

	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires)
}
