package user

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/domain/auth"
)

const getByEmailQuery = `
SELECT * FROM users where email = $1
`

func (r *Repository) GetByEmail(ctx context.Context, email string) (*auth.User, error) {
	var user auth.User
	err := r.DB.QueryRow(ctx, getByEmailQuery, email).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil

}
