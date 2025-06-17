package auth

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
)

func (s *Service) GetProfile(ctx context.Context, p *web.ProfileReq) (*web.ProfileRes, error) {
	u, err := s.userRepository.GetByID(ctx, p.UserID)
	if err != nil {
		return nil, err
	}

	return &web.ProfileRes{
		User: u,
	}, nil

}
