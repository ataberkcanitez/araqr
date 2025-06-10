package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
	"github.com/cockroachdb/errors"
)

func (svc *StickerService) Get(ctx context.Context, req *web.GetStickerRequest) (*sticker.Sticker, error) {
	stx, err := svc.stickerRepository.GetByID(ctx, req.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sticker")
	}
	return stx, nil
}
