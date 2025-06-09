package refresh_token

import (
	"context"
	"time"
)

const invalidateRefreshTokenQuery = `
UPDATE refresh_tokens
SET valid = false, updated_at = $1
WHERE user_id = $2
`

func (r *Repository) Invalidate(ctx context.Context, userID string) error {
	now := time.Now().UTC()
	_, err := r.DB.Exec(ctx, invalidateRefreshTokenQuery, now, userID)
	if err != nil {
		return err
	}
	return nil
}
