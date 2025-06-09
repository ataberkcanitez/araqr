package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/handler"
	"github.com/cockroachdb/errors"
)

func (svc *StickerService) ListMyStickers(ctx context.Context, req *handler.ListMyStickersRequest) (*handler.ListMyStickersResponse, error) {
	stickers, err := svc.stickerRepository.ListByUserID(ctx, req.UserID, req.Limit, req.Page)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list stickers for user")
	}
	return &handler.ListMyStickersResponse{
		Stickers: stickers,
	}, nil
}
