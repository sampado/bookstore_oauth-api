package access_token

import (
	"github.com/sampado/bookstore_oauth-api/src/repository/rest"
	"github.com/sampado/bookstore_utils-go/rest_errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *rest_errors.RestError)
	Create(AccessToken) *rest_errors.RestError
	UpdateExpirationTime(AccessToken) *rest_errors.RestError
}

type Service interface {
	GetById(string) (*AccessToken, *rest_errors.RestError)
	Create(AccessTokenRequest) (*AccessToken, *rest_errors.RestError)
	UpdateExpirationTime(AccessToken) *rest_errors.RestError
}

type service struct {
	repository    Repository
	restUsersRepo rest.RestUsersRepository
}

func NewService(repo Repository, restRepo rest.RestUsersRepository) Service {
	return &service{
		repository:    repo,
		restUsersRepo: restRepo,
	}
}

func (s *service) GetById(ID string) (*AccessToken, *rest_errors.RestError) {
	accessToken, err := s.repository.GetById(ID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request AccessTokenRequest) (*AccessToken, *rest_errors.RestError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// Authenticate the user against the users API:
	user, err := s.restUsersRepo.LoginUser(request.UserName, request.Password)
	if err != nil {
		return nil, err
	}

	at := GetNewAccessToken(user.Id)
	at.Generate()
	if err := s.repository.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *rest_errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.UpdateExpirationTime(at)
}
