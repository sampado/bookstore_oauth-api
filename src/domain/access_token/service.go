package access_token

import "github.com/sampado/bookstore_oauth-api/src/utils/errors"

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(ID string) (*AccessToken, *errors.RestError) {
	accessToken, err := s.repository.GetById(ID)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
