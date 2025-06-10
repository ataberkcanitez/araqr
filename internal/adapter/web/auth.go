package web

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type authSvc interface {
	Register(ctx context.Context, req *RegisterReq) (*RegisterRes, error)
	Login(ctx context.Context, req *LoginReq) (*LoginRes, error)
}

type AuthHandler struct {
	svc authSvc
}

func NewAuthHandler(svc authSvc) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

func (h *AuthHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/v1/register", h.register)
	e.POST("/v1/login", h.login)
}

type (
	RegisterReq struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		PhoneNumber string `json:"phone_number"`
	}

	RegisterRes struct {
		ID          string    `json:"id"`
		Email       string    `json:"email"`
		FirstName   string    `json:"first_name"`
		LastName    string    `json:"last_name"`
		PhoneNumber string    `json:"phone_number"`
		CreatedAt   time.Time `json:"created_at"`
	}
)

func (h *AuthHandler) register(c echo.Context) error {
	var req RegisterReq
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(errors.New("bad request"), err.Error())
	}

	if req.Email == "" || req.Password == "" {
		return errors.Wrap(errors.New("bad request"), "email and password are required")
	}

	res, err := h.svc.Register(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

type (
	LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginRes struct {
		AccessToken  string    `json:"access_token"`
		RefreshToken string    `json:"refresh_token"`
		ExpiresAt    time.Time `json:"expires_at"`
	}
)

func (h *AuthHandler) login(c echo.Context) error {
	var req LoginReq
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(err, "failed to bind login request")
	}

	if req.Email == "" || req.Password == "" {
		return errors.Wrap(errors.New("Bad request"), "email and password are required")
	}

	res, err := h.svc.Login(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
