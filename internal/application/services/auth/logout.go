package auth

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
)

func (s *Service) Logout(ctx context.Context, in *web.LogoutReq) (*web.LogoutRes, error) {
	u, err := s.userRepository.GetByEmail(ctx, in.Email)
	if err != nil {
		return nil, err
	}

	if err := s.refreshTokenRepository.Invalidate(ctx, u.ID); err != nil {
		return nil, err
	}

	return &web.LogoutRes{}, nil
}
