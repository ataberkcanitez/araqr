package message

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
)

type Repository interface {
	Create(ctx context.Context, message *sticker.Message) error
	GetByStickerID(ctx context.Context, id string, limit int, page int) ([]*sticker.Message, error)
}
