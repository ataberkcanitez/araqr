package user

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/auth"
	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v4"
)

const getByIDQuery = `
SELECT * FROM users where id = $1
`

func (r *Repository) GetByID(ctx context.Context, ID string) (*auth.User, error) {
	var user auth.User
	err := r.DB.QueryRow(ctx, getByIDQuery, ID).Scan(
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
			return nil, auth.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil

}
