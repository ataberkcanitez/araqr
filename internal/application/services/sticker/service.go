package sticker

import (
	"github.com/ataberkcanitez/araqr/internal/application/ports/outbound/auth"
	"github.com/ataberkcanitez/araqr/internal/application/ports/outbound/message"
	"github.com/ataberkcanitez/araqr/internal/application/ports/outbound/sticker"
)

type Service struct {
	userRepository    auth.UserRepository
	stickerRepository sticker.Repository
	messageRepository message.Repository
}

func NewService(
	userRepository auth.UserRepository,
	stickerRepository sticker.Repository,
	messageRepository message.Repository,
) *Service {
	return &Service{
		userRepository:    userRepository,
		stickerRepository: stickerRepository,
		messageRepository: messageRepository,
	}
}
