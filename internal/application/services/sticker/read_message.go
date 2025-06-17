package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
)

func (s *Service) SetMessageAsRead(ctx context.Context, req *web.SetMessageAsReadRequest) (*web.SetMessageAsReadResponse, error) {
	stx, err := s.Get(ctx, &web.GetStickerRequest{ID: req.StickerID})
	if err != nil {
		return nil, err
	}
	if stx.UserID != req.UserID {
		return nil, sticker.ErrStickerNotOwnedByUser
	}
	msg, err := s.messageRepository.GetByID(ctx, req.MessageID)
	if err != nil {
		return nil, err
	}
	msg.SetAsRead()
	if err := s.messageRepository.Update(ctx, msg); err != nil {
		return nil, err
	}
	return &web.SetMessageAsReadResponse{}, nil
}
