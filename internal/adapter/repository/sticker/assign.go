package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
	"time"
)

const AssignStickerToUserQuery = `
UPDATE stickers
SET user_id = $1, updated_at = $2
WHERE id = $3
`

func (r *Repository) Assign(ctx context.Context, s *sticker.Sticker) (*sticker.Sticker, error) {
	now := time.Now().UTC()
	_, err := r.DB.Exec(ctx, AssignStickerToUserQuery,
		s.UserID,
		now,
		s.ID,
	)
	if err != nil {
		return nil, err
	}
	s.UpdatedAt = now
	return s, nil
}
