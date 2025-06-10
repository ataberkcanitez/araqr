package auth

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/cockroachdb/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, in *web.LoginReq) (*web.LoginRes, error) {
	user, err := s.userRepository.GetByEmail(ctx, in.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return nil, errors.Wrapf(errors.New("password mismatch"), "login: password does not match %s", err.Error())
	}

	t, err := s.generateToken(ctx, user)
	if err != nil {
		return nil, err
	}

	if err := s.refreshTokenRepository.Upsert(ctx, t.RefreshToken); err != nil {
		return nil, err
	}

	return &web.LoginRes{
		AccessToken:  t.Token,
		RefreshToken: t.RefreshToken.Token,
		ExpiresAt:    t.ExpiresAt,
	}, nil
}
