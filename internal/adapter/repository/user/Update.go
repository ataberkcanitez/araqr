package user

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/auth"
	"time"
)

const UpdateUserQuery = `
Update users
Set first_name = $1,
    last_name = $2,
    password = $3,
	phone_number = $4,
	updated_at = $5
Where id = $6
`

func (r *Repository) Update(ctx context.Context, user *auth.User) (*auth.User, error) {
	now := time.Now().UTC()
	user.UpdatedAt = now
	_, err := r.DB.Exec(ctx, UpdateUserQuery,
		user.FirstName,
		user.LastName,
		user.Password,
		user.PhoneNumber,
		user.UpdatedAt,
		user.ID,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
