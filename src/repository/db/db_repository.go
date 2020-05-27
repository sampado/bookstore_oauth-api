package db

import (
	"github.com/gocql/gocql"
	"github.com/sampado/bookstore_oauth-api/src/clients/cassandra"
	"github.com/sampado/bookstore_oauth-api/src/domain/access_token"
	"github.com/sampado/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, client_id, expires, user_id FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token, client_id, expires, user_id)  VALUES(?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type dbRepository struct {
}

func (r *dbRepository) GetById(ID string) (*access_token.AccessToken, *errors.RestError) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, ID).Scan(
		&result.AccessToken,
		&result.ClientId,
		&result.Expires,
		&result.UserId,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError(err.Error())
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestError {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer session.Close()

	if err := session.Query(
		queryCreateAccessToken,
		token.AccessToken,
		token.ClientId,
		token.Expires,
		token.UserId,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
func (r *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestError {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer session.Close()

	if err := session.Query(
		queryUpdateExpires,
		token.Expires,
		token.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
func NewRepository() DbRepository {
	return &dbRepository{}
}
