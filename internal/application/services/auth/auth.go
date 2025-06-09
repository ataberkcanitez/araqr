package auth

import (
	"context"
	"fmt"
	"github.com/ataberkcanitez/araqr/handler"
	auth2 "github.com/ataberkcanitez/araqr/internal/application/ports/outbound/auth"
	"github.com/ataberkcanitez/araqr/internal/domain"
	"github.com/ataberkcanitez/araqr/internal/domain/auth"
	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthConfig struct {
	SecretKey          string        `mapstructure:"secret-key"`
	AccessTokenExpiry  time.Duration `mapstructure:"access-token-expiry"`
	RefreshTokenExpiry time.Duration `mapstructure:"refresh-token-expiry"`
}

type AuthService struct {
	cfg                    *AuthConfig
	userRepository         auth2.UserRepository
	refreshTokenRepository auth2.RefreshTokenRepository
}

func NewAuthService(cfg *AuthConfig, userRepository auth2.UserRepository, refreshTokenRepository auth2.RefreshTokenRepository) *AuthService {
	return &AuthService{
		cfg:                    cfg,
		userRepository:         userRepository,
		refreshTokenRepository: refreshTokenRepository,
	}
}

func (s *AuthService) Register(ctx context.Context, in *handler.RegisterReq) (*handler.RegisterRes, error) {
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

	return &handler.RegisterRes{
		ID:          insertedUser.ID,
		Email:       insertedUser.Email,
		FirstName:   insertedUser.FirstName,
		LastName:    insertedUser.LastName,
		PhoneNumber: insertedUser.PhoneNumber,
		CreatedAt:   insertedUser.CreatedAt,
	}, nil

}

func (s *AuthService) Login(ctx context.Context, in *handler.LoginReq) (*handler.LoginRes, error) {
	user, err := s.userRepository.GetByEmail(ctx, in.Email)
	if err != nil {
		fmt.Println("Error fetching user by email:", err)
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

	return &handler.LoginRes{
		AccessToken:  t.Token,
		RefreshToken: t.RefreshToken.Token,
		ExpiresAt:    t.ExpiresAt,
	}, nil
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

func (s *AuthService) generateToken(_ context.Context, user *auth.User) (*token, error) {
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

// Parse parses a token
func (s *AuthService) Parse(ctx context.Context, in *handler.ParseTokenReq) (*handler.ParseTokenRes, error) {
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

	return &handler.ParseTokenRes{
		ID:    claims.Id,
		Email: claims.Email,
	}, nil
}
