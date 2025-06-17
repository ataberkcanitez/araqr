package message

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
)

type Repository interface {
	Create(ctx context.Context, message *sticker.Message) error
	GetByID(ctx context.Context, id string) (*sticker.Message, error)
	GetByStickerID(ctx context.Context, id string, limit int, page int) ([]*sticker.Message, error)
	Update(ctx context.Context, msg *sticker.Message) error
}
