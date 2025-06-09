package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
)

type Repository interface {
	Create(ctx context.Context, s *sticker.Sticker) (*sticker.Sticker, error)
	Assign(ctx context.Context, s *sticker.Sticker) (*sticker.Sticker, error)
	GetByID(ctx context.Context, id string) (*sticker.Sticker, error)
	ListByUserID(ctx context.Context, userID string, limit, page int) ([]*sticker.Sticker, error)
	Update(ctx context.Context, stx *sticker.Sticker) error
}
