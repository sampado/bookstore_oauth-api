package db

import (
	"github.com/gocql/gocql"
	"github.com/sampado/bookstore_oauth-api/src/clients/cassandra"
	"github.com/sampado/bookstore_oauth-api/src/domain/access_token"
	"github.com/sampado/bookstore_utils-go/rest_errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, client_id, expires, user_id FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token, client_id, expires, user_id)  VALUES(?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, rest_errors.RestError)
	Create(access_token.AccessToken) rest_errors.RestError
	UpdateExpirationTime(access_token.AccessToken) rest_errors.RestError
}

type dbRepository struct {
}

func (r *dbRepository) GetById(ID string) (*access_token.AccessToken, rest_errors.RestError) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, ID).Scan(
		&result.AccessToken,
		&result.ClientId,
		&result.Expires,
		&result.UserId,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError(err.Error())
		}
		return nil, rest_errors.NewInternalServerError(err.Error(), err)
	}
	return &result, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) rest_errors.RestError {
	if err := cassandra.GetSession().Query(
		queryCreateAccessToken,
		token.AccessToken,
		token.ClientId,
		token.Expires,
		token.UserId,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(token access_token.AccessToken) rest_errors.RestError {
	if err := cassandra.GetSession().Query(
		queryUpdateExpires,
		token.Expires,
		token.AccessToken,
	).Exec(); err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}

	return nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
