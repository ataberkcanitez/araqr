package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	sticker2 "github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
	"github.com/cockroachdb/errors"
)

func (svc *Service) Assign(ctx context.Context, req *web.AssignStickerRequest) (*sticker2.Sticker, error) {
	stx, err := svc.stickerRepository.GetByID(ctx, req.StickerID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sticker")
	}

	if stx == nil {
		return nil, sticker2.ErrStickerNotFound
	}

	if stx.UserID != "" {
		return nil, sticker2.ErrStickerAlreadyAssigned
	}

	stx.Assign(req.UserID)
	assign, err := svc.stickerRepository.Assign(ctx, stx)
	if err != nil {
		stx.UserID = ""
		return nil, errors.Wrap(err, "failed to assign sticker")
	}
	return assign, nil
}
