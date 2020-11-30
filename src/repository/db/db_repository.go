package db

import (
	"github.com/kaitus/bookstore_oauth-api-golang/src/domain/access_token"
	"github.com/kaitus/bookstore_oauth-api-golang/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	return nil, errors.NewInternalServerError("database connection not implement yet")
}
