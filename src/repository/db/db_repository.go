package db

import (
	"github.com/sampado/bookstore_oauth-api/src/clients/cassandra"
	"github.com/sampado/bookstore_oauth-api/src/domain/access_token"
	"github.com/sampado/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(ID string) (*access_token.AccessToken, *errors.RestError) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.Query("SELECT access_token, client_id, expires, user_id FROM access_tokens WHERE access_token='ssd' ; ")
	return nil, errors.NewInternalServerError("database hasn't been implemented yet")
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
