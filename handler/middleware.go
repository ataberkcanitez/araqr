package handler

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/domain"
	"github.com/cockroachdb/errors"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	ParseTokenReq struct {
		Token string `json:"token"`
	}

	ParseTokenRes struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}
)

type tokenParser interface {
	Parse(ctx context.Context, token *ParseTokenReq) (*ParseTokenRes, error)
}

// MiddlewareTokenVerification is the middleware function to verify the Cognito token
func MiddlewareTokenVerification(parser tokenParser) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorizationHeaderToken := c.Request().Header.Get("Authorization")
			if authorizationHeaderToken == "" || !strings.HasPrefix(authorizationHeaderToken, "Bearer ") {
				return domain.ErrAuthBearerHeaderNotFound
			}

			authorizationToken := strings.TrimPrefix(authorizationHeaderToken, "Bearer ")
			claims, err := parser.Parse(c.Request().Context(), &ParseTokenReq{Token: authorizationToken})
			if err != nil {
				return errors.Wrap(err, "failed to verify token")
			}

			c.Set("claims", claims)
			return next(c)
		}
	}
}

// GetClaims returns the claims from the context
func GetClaims(c echo.Context) *ParseTokenRes {
	return c.Get("claims").(*ParseTokenRes)
}
