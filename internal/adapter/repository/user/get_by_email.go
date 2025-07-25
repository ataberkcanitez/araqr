package user

import (
	"context"
	auth2 "github.com/ataberkcanitez/araqr/internal/application/domain/auth"
	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v4"
)

const getByEmailQuery = `
SELECT * FROM users where email = $1
`

func (r *Repository) GetByEmail(ctx context.Context, email string) (*auth2.User, error) {
	var user auth2.User
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, auth2.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil

}
