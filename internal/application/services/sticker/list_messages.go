package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	sticker2 "github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
)

func (svc *Service) ListMessages(ctx context.Context, req *web.ListMessagesRequest) ([]*sticker2.Message, error) {
	stx, err := svc.Get(ctx, &web.GetStickerRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}
	if stx == nil {
		return nil, sticker2.ErrStickerNotFound
	}
	if stx.UserID != req.UserID {
		return nil, sticker2.ErrStickerNotOwnedByUser
	}
	messages, err := svc.messageRepository.GetByStickerID(ctx, stx.ID, req.Limit, req.Page)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
