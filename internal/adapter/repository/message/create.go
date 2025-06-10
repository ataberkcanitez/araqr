package message

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
)

const insertMessageQuery = `
INSERT INTO messages (id, sticker_id, message, urgency_level, created_at) VALUES ($1, $2, $3, $4, $5);
`

func (r *Repository) Create(ctx context.Context, m *sticker.Message) error {
	_, err := r.DB.Exec(ctx, insertMessageQuery,
		m.ID,
		m.StickerID,
		m.Message,
		m.UrgencyLevel,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
