package auth

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/auth"
)

type UserRepository interface {
	Create(ctx context.Context, u *auth.User) (*auth.User, error)
	GetByEmail(ctx context.Context, email string) (*auth.User, error)
	GetByID(ctx context.Context, ID string) (*auth.User, error)
	Update(ctx context.Context, u *auth.User) (*auth.User, error)
}
