package message

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v4"
)

const getMessageByIDQuery = `
SELECT id, sticker_id, message, urgency_level, read, updated_at, created_at FROM messages
WHERE id = $1
`

func (r *Repository) GetByID(ctx context.Context, ID string) (*sticker.Message, error) {
	var message sticker.Message
	err := r.DB.QueryRow(ctx, getMessageByIDQuery, ID).Scan(
		&message.ID,
		&message.StickerID,
		&message.Message,
		&message.UrgencyLevel,
		&message.Read,
		&message.UpdatedAt,
		&message.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, sticker.ErrMessageNotFound
		}
		return nil, err
	}
	return &message, nil
}
