package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/cockroachdb/errors"
)

func (svc *StickerService) ListMyStickers(ctx context.Context, req *web.ListMyStickersRequest) (*web.ListMyStickersResponse, error) {
	stickers, err := svc.stickerRepository.ListByUserID(ctx, req.UserID, req.Limit, req.Page)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list stickers for user")
	}
	return &web.ListMyStickersResponse{
		Stickers: stickers,
	}, nil
}
