package message

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
	"time"
)

const updateMessageQuery = `
UPDATE messages
SET
  read = $1,
  updated_at = $2
WHERE
  id = $3;
`

func (r *Repository) Update(ctx context.Context, msg *sticker.Message) error {
	now := time.Now().UTC()
	msg.UpdatedAt = now
	_, err := r.DB.Exec(ctx, updateMessageQuery,
		msg.Read,
		msg.UpdatedAt,
		msg.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
