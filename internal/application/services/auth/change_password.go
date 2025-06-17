package auth

import (
	"context"
	"fmt"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	auth2 "github.com/ataberkcanitez/araqr/internal/application/domain/auth"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) ChangePassword(ctx context.Context, in *web.ChangePasswordReq) (*web.ChangePasswordRes, error) {
	u, err := s.userRepository.GetByEmail(ctx, in.Email)
	if err != nil {
		return nil, err
	}

	fmt.Println("current password hash:", u.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.OldPassword)); err != nil {
		return nil, auth2.ErrInvalidPassword
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	fmt.Println("new password hash:", string(hashedPassword))

	err = u.UpdatePassword(string(hashedPassword))
	if err != nil {
		return nil, err
	}
	_, err = s.userRepository.Update(ctx, u)
	if err != nil {
		return nil, err
	}
	return &web.ChangePasswordRes{}, nil
}
