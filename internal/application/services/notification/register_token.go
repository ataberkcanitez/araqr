package notification

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/ataberkcanitez/araqr/internal/application/domain/notification"
	"github.com/google/uuid"
	"time"
)

func (s *Service) RegisterToken(ctx context.Context, req *web.RegisterTokenReq) (*web.RegisterTokenRes, error) {
	now := time.Now().UTC()
	in := notification.Token{
		ID:         uuid.NewString(),
		UserID:     req.UserID,
		PushToken:  req.PushToken,
		Platform:   req.Platform,
		DeviceName: req.DeviceInfo.DeviceName,
		OSName:     req.DeviceInfo.OSName,
		OSVersion:  req.DeviceInfo.OSVersion,
		IsActive:   true,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	if err := s.notificationRepository.RegisterToken(ctx, in); err != nil {
		return nil, err
	}

	return &web.RegisterTokenRes{
		Success: true,
		Message: "Token registered successfully",
	}, nil

}
