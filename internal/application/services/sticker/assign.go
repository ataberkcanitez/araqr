package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/handler"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
	"github.com/cockroachdb/errors"
)

func (svc *StickerService) Assign(ctx context.Context, req *handler.AssignStickerRequest) (*sticker.Sticker, error) {
	stx, err := svc.stickerRepository.GetByID(ctx, req.StickerID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sticker")
	}

	if stx == nil {
		return nil, sticker.ErrStickerNotFound
	}

	if stx.UserID != "" {
		return nil, sticker.ErrStickerAlreadyAssigned
	}

	stx.Assign(req.UserID)
	assign, err := svc.stickerRepository.Assign(ctx, stx)
	if err != nil {
		stx.UserID = ""
		return nil, errors.Wrap(err, "failed to assign sticker")
	}
	return assign, nil
}
