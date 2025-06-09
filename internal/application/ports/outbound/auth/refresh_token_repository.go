package auth

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/domain/auth"
)

type RefreshTokenRepository interface {
	Upsert(ctx context.Context, rt *auth.RefreshToken) error
	Invalidate(ctx context.Context, userID string) error
	Get(ctx context.Context, token string) (*auth.RefreshToken, error)
}
