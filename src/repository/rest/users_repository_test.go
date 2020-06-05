package rest

import (
	"fmt"
	"net/http"

	"github.com/stretchr/testify/assert"

	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com", "password":"password"}`,
		RespHTTPCode: 504,
		RespBody:     `{}`,
	},
	)
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com", "password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "an error message", "status": "404", "error": "not_found" }`,
	},
	)
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com", "password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "an error message", "status": 404, "error": "not_found" }`,
	},
	)
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com", "password":"password"}`,
		RespHTTPCode: http.StatusOK,
		// passing id as string
		RespBody: `{"id":"2","firstName":"Santiago","lastName":"Correa","email":"sampado@gmail.com","dataCreated":"2020-05-08:53:38Z","status":"active"}`,
	},
	)
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestLoginUserOK(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@gmail.com", "password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id":2,"firstName":"Santiago","lastName":"Correa","email":"sampado@gmail.com","dataCreated":"2020-05-08:53:38Z","status":"active"}`,
	},
	)
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 2, user.Id)
}
