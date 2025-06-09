package message

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
)

const GetMessageByStickerIDQuery = `
SELECT id, sticker_id, message, urgency_level, created_at
FROM messages 
WHERE sticker_id = $1
LIMIT $2 OFFSET $3
`

func (r *Repository) GetByStickerID(ctx context.Context, id string, limit int, page int) ([]*sticker.Message, error) {
	rows, err := r.DB.Query(ctx, GetMessageByStickerIDQuery, id, limit, page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var messages []*sticker.Message
	for rows.Next() {
		var m sticker.Message
		err := rows.Scan(
			&m.ID,
			&m.StickerID,
			&m.Message,
			&m.UrgencyLevel,
			&m.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &m)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil

}
