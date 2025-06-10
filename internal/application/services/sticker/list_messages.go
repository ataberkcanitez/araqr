package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
)

func (svc *Service) ListMessages(ctx context.Context, req *web.ListMessagesRequest) ([]*sticker.Message, error) {
	stx, err := svc.Get(ctx, &web.GetStickerRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}
	if stx == nil {
		return nil, sticker.ErrStickerNotFound
	}
	if stx.UserID != req.UserID {
		return nil, sticker.ErrStickerNotOwnedByUser
	}
	messages, err := svc.messageRepository.GetByStickerID(ctx, stx.ID, req.Limit, req.Page)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
