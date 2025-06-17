package user

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/auth"
	"time"
)

const createQuery = `
INSERT INTO users (id, email, first_name, last_name, password, phone_number, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

func (r *Repository) Create(ctx context.Context, u *auth.User) (*auth.User, error) {
	now := time.Now().UTC()
	_, err := r.DB.Exec(ctx, createQuery,
		u.ID,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Password,
		u.PhoneNumber,
		now,
		now,
	)
	if err != nil {
		return nil, err
	}
	return u, nil
}
