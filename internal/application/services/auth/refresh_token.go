package auth

import (
	"context"
	"time"

	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/ataberkcanitez/araqr/internal/application/domain/auth"
)

func (s *Service) RefreshToken(ctx context.Context, req *web.RefreshTokenReq) (*web.RefreshTokenRes, error) {
	refreshToken, err := s.refreshTokenRepository.Get(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	if !refreshToken.Valid || refreshToken.ValidUntil.Before(now) {
		return nil, auth.ErrInvalidToken
	}
	user, err := s.userRepository.GetByID(ctx, refreshToken.UserID)
	if err != nil {
		return nil, err
	}

	t, err := s.generateToken(ctx, user)
	if err != nil {
		return nil, err
	}
	refreshToken.Token = t.RefreshToken.Token
	if err := s.refreshTokenRepository.Upsert(ctx, refreshToken); err != nil {
		return nil, err
	}
	return &web.RefreshTokenRes{
		AccessToken:  t.Token,
		RefreshToken: t.RefreshToken.Token,
		ExpiresAt:    t.ExpiresAt,
	}, nil

}
