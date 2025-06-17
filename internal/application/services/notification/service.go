package notification

import "github.com/ataberkcanitez/araqr/internal/application/ports/outbound/notification"

type Service struct {
	notificationRepository notification.Repository
}

func NewService(notificationRepository notification.Repository) *Service {
	return &Service{
		notificationRepository: notificationRepository,
	}
}
