package message

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
)

const insertMessageQuery = `
INSERT INTO messages (id, sticker_id, message, urgency_level, read, updated_at, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7);
`

func (r *Repository) Create(ctx context.Context, m *sticker.Message) error {
	_, err := r.DB.Exec(ctx, insertMessageQuery,
		m.ID,
		m.StickerID,
		m.Message,
		m.UrgencyLevel,
		m.Read,
		m.UpdatedAt,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
