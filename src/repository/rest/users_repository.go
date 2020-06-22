package rest

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/sampado/bookstore_oauth-api/src/domain/users"
	"github.com/sampado/bookstore_utils-go/rest_errors"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestError)
}

type usersRepository struct{}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestError) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, rest_errors.NewInternalServerError("invalid rest client response when trying to login user", errors.New("null response"))
	}

	if response.StatusCode > 299 {
		var restErr rest_errors.RestError
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user", err)
		}
		return nil, restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("invalid user interface when trying to login user", err)
	}

	return &user, nil
}
