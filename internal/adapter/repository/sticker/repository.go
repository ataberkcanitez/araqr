package sticker

import "github.com/ataberkcanitez/araqr/pgsql"

type Repository struct {
	DB *pgsql.DB
}

func NewRepository(db *pgsql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
