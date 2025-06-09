package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/handler"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
	"github.com/google/uuid"
	"time"
)

func (svc *StickerService) SendMessageToSticker(ctx context.Context, req *handler.SendMessageToStickerRequest) (*handler.SendMessageToStickerResponse, error) {
	stx, err := svc.stickerRepository.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if stx == nil {
		return nil, sticker.ErrStickerNotFound
	}

	if !stx.Active {
		return nil, sticker.ErrStickerNotFound
	}

	msg := &sticker.Message{
		ID:           uuid.NewString(),
		StickerID:    req.ID,
		Message:      req.Message,
		UrgencyLevel: req.UrgencyLevel,
		CreatedAt:    time.Now(),
	}
	if err := svc.messageRepository.Create(ctx, msg); err != nil {
		return nil, err
	}

	return &handler.SendMessageToStickerResponse{
		ID:      msg.ID,
		Message: "Message sent successfully",
	}, nil
}
