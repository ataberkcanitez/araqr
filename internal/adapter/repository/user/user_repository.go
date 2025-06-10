package user

import (
	"github.com/ataberkcanitez/araqr/pgsql"
)

type Repository struct {
	DB *pgsql.DB
}

func NewRepository(db *pgsql.DB) *Repository {
	repository := &Repository{
		DB: db,
	}
	return repository
}
