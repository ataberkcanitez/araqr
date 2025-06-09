package message

import "github.com/ataberkcanitez/araqr/pgsql"

type Repository struct {
	DB *pgsql.DB
}

func NewMessageRepository(db *pgsql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
