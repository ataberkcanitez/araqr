package notification

import "github.com/ataberkcanitez/araqr/pgsql"

type Repository struct {
	DB *pgsql.DB
}

func NewNotificationRepository(db *pgsql.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
