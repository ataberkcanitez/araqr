package refresh_token

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/domain/auth"
)

const getRefreshTokenQuery = `SELECT * FROM refresh_tokens WHERE token = $1`

func (r *Repository) Get(ctx context.Context, token string) (*auth.RefreshToken, error) {
	var rt auth.RefreshToken
	err := r.DB.QueryRow(ctx, getRefreshTokenQuery, token).Scan(
		&rt.UserID,
		&rt.Token,
		&rt.Valid,
		&rt.ValidUntil,
		&rt.CreatedAt,
		&rt.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &rt, nil
}
