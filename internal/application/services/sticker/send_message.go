package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	sticker2 "github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
	"github.com/google/uuid"
	"time"
)

// SendMessageToSticker New feature idea:
// - Lets get IP address of sender and store it in the message
// - Lets allow sticker owner to block sender IP address
// - If the sender IP address is blocked, do not allow sending message.
func (s *Service) SendMessageToSticker(ctx context.Context, req *web.SendMessageToStickerRequest) (*web.SendMessageToStickerResponse, error) {
	stx, err := s.stickerRepository.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if stx == nil {
		return nil, sticker2.ErrStickerNotFound
	}

	if !stx.Active {
		return nil, sticker2.ErrStickerNotFound
	}

	msg := &sticker2.Message{
		ID:           uuid.NewString(),
		StickerID:    req.ID,
		Message:      req.Message,
		UrgencyLevel: req.UrgencyLevel,
		Read:         false,
		CreatedAt:    time.Now(),
	}
	if err := s.messageRepository.Create(ctx, msg); err != nil {
		return nil, err
	}

	return &web.SendMessageToStickerResponse{
		ID:      msg.ID,
		Message: "Message sent successfully",
	}, nil
}
