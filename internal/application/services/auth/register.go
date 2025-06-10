package auth

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/ataberkcanitez/araqr/internal/domain/auth"
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Register(ctx context.Context, in *web.RegisterReq) (*web.RegisterRes, error) {
	if _, err := s.userRepository.GetByEmail(ctx, in.Email); err == nil {
		return nil, errors.Wrap(errors.New("bad request"), "email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "could not hash password")
	}

	u := &auth.User{
		ID:          uuid.NewString(),
		Email:       in.Email,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		Password:    string(hashedPassword),
		PhoneNumber: in.PhoneNumber,
	}

	insertedUser, err := s.userRepository.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return &web.RegisterRes{
		ID:          insertedUser.ID,
		Email:       insertedUser.Email,
		FirstName:   insertedUser.FirstName,
		LastName:    insertedUser.LastName,
		PhoneNumber: insertedUser.PhoneNumber,
		CreatedAt:   insertedUser.CreatedAt,
	}, nil

}
