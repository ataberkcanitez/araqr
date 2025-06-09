package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/handler"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
	"github.com/cockroachdb/errors"
)

func (svc *StickerService) Get(ctx context.Context, req *handler.GetStickerRequest) (*sticker.Sticker, error) {
	stx, err := svc.stickerRepository.GetByID(ctx, req.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sticker")
	}
	return stx, nil
}
