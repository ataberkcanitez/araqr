package refresh_token

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/domain/auth"
	"time"
)

const upsertRefreshTokenQuery = `
INSERT INTO refresh_tokens (user_id, token, valid,  valid_until, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (user_id) DO UPDATE SET
    token = EXCLUDED.token,
    valid = EXCLUDED.valid,
    valid_until = EXCLUDED.valid_until,
    updated_at = NOW()
`

func (r *Repository) Upsert(ctx context.Context, rt *auth.RefreshToken) error {
	now := time.Now().UTC()
	_, err := r.DB.Exec(ctx, upsertRefreshTokenQuery,
		rt.UserID,
		rt.Token,
		rt.Valid,
		rt.ValidUntil,
		now,
		now,
	)
	if err != nil {
		return err
	}

	return nil
}
