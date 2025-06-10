package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
	"github.com/cockroachdb/errors"
)

func (svc *Service) ListMyStickers(ctx context.Context, req *web.ListMyStickersRequest) (*web.ListMyStickersResponse, error) {
	stickers, err := svc.stickerRepository.ListByUserID(ctx, req.UserID, req.Limit, req.Page)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list stickers for user")
	}
	if stickers == nil {
		return &web.ListMyStickersResponse{
			Stickers: []*sticker.Sticker{},
		}, nil

	}
	return &web.ListMyStickersResponse{
		Stickers: stickers,
	}, nil
}
