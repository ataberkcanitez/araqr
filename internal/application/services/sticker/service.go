package sticker

import (
	"github.com/ataberkcanitez/araqr/internal/application/ports/outbound/auth"
	"github.com/ataberkcanitez/araqr/internal/application/ports/outbound/message"
	"github.com/ataberkcanitez/araqr/internal/application/ports/outbound/sticker"
)

type StickerService struct {
	userRepository    auth.UserRepository
	stickerRepository sticker.Repository
	messageRepository message.Repository
}

func NewStickerService(
	userRepository auth.UserRepository,
	stickerRepository sticker.Repository,
	messageRepository message.Repository,
) *StickerService {
	return &StickerService{
		userRepository:    userRepository,
		stickerRepository: stickerRepository,
		messageRepository: messageRepository,
	}
}
