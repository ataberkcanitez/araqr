package auth

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	authout "github.com/ataberkcanitez/araqr/internal/application/ports/outbound/auth"
	"github.com/ataberkcanitez/araqr/internal/domain"
	"github.com/ataberkcanitez/araqr/internal/domain/auth"
	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type Config struct {
	SecretKey          string        `mapstructure:"secret-key"`
	AccessTokenExpiry  time.Duration `mapstructure:"access-token-expiry"`
	RefreshTokenExpiry time.Duration `mapstructure:"refresh-token-expiry"`
}

type Service struct {
	cfg                    *Config
	userRepository         authout.UserRepository
	refreshTokenRepository authout.RefreshTokenRepository
}

func NewService(cfg *Config, userRepository authout.UserRepository, refreshTokenRepository authout.RefreshTokenRepository) *Service {
	return &Service{
		cfg:                    cfg,
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
	}
}

type (
	customJwtClaims struct {
		jwt.StandardClaims
		Email string `json:"email"`
	}

	token struct {
		Token        string
		RefreshToken *auth.RefreshToken
		ExpiresAt    time.Time
	}
)

func (s *Service) generateToken(_ context.Context, user *auth.User) (*token, error) {
	now := time.Now().UTC()
	accessTokenExpiresAt := now.Add(s.cfg.AccessTokenExpiry)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, customJwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpiresAt.Unix(),
			Id:        user.ID,
			IssuedAt:  now.Unix(),
		},
		Email: user.Email,
	})

	signedToken, err := jwtToken.SignedString([]byte(s.cfg.SecretKey))
	if err != nil {
		return nil, errors.Wrap(err, "could not sign token")
	}

	return &token{
		Token: signedToken,
		RefreshToken: &auth.RefreshToken{
			UserID:     user.ID,
			Token:      uuid.NewString(),
			Valid:      true,
			ValidUntil: now.Add(s.cfg.RefreshTokenExpiry),
		},
		ExpiresAt: accessTokenExpiresAt,
	}, nil

}

func (s *Service) Parse(_ context.Context, in *web.ParseTokenReq) (*web.ParseTokenRes, error) {
	var claims customJwtClaims

	_, err := jwt.ParseWithClaims(in.Token, &claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.Wrap(domain.ErrInvalidTokenSigningMethod, "parse: invalid signing method")
		}
		return []byte(s.cfg.SecretKey), nil
	})
	if err != nil {
		return nil, errors.Wrap(domain.ErrInvalidToken, err.Error())
	}

	return &web.ParseTokenRes{
		ID:    claims.Id,
		Email: claims.Email,
	}, nil
}
