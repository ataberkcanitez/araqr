package auth

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/domain/auth"
)

type UserRepository interface {
	Create(ctx context.Context, u *auth.User) (*auth.User, error)
	GetByEmail(ctx context.Context, email string) (*auth.User, error)
}
