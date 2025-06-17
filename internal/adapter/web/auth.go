package web

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/auth"
	"github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type authSvc interface {
	Register(ctx context.Context, req *RegisterReq) (*RegisterRes, error)
	Login(ctx context.Context, req *LoginReq) (*LoginRes, error)
	Logout(ctx context.Context, req *LogoutReq) (*LogoutRes, error)
	ChangePassword(ctx context.Context, c *ChangePasswordReq) (*ChangePasswordRes, error)
	RefreshToken(ctx context.Context, req *RefreshTokenReq) (*RefreshTokenRes, error)
	GetProfile(ctx context.Context, p *ProfileReq) (*ProfileRes, error)
}

type AuthHandler struct {
	svc authSvc
	tp  tokenParser
}

func NewAuthHandler(svc authSvc, parser tokenParser) *AuthHandler {
	return &AuthHandler{
		svc: svc,
		tp:  parser,
	}
}

func (h *AuthHandler) RegisterRoutes(e *echo.Echo) {
	basePrefix := "/api/v1/auth"
	public := e.Group(basePrefix)
	private := e.Group(basePrefix, MiddlewareTokenVerification(h.tp))

	public.POST("/register", h.register)
	public.POST("/login", h.login)
	public.POST("/refresh-token", h.refreshToken)

	private.POST("/logout", h.logout)
	private.POST("/change-password", h.changePassword)
	private.GET("/profile", h.profile)

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

type (
	LogoutReq struct {
		Email string
	}

	LogoutRes struct{}
)

func (h *AuthHandler) logout(c echo.Context) error {
	claims := GetClaims(c)
	req := &LogoutReq{
		Email: claims.Email,
	}

	res, err := h.svc.Logout(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

type (
	ChangePasswordReq struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`

		Email string
	}

	ChangePasswordRes struct{}
)

func (h *AuthHandler) changePassword(c echo.Context) error {
	var req ChangePasswordReq
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(errors.New("bad request"), err.Error())
	}

	if req.OldPassword == "" || req.NewPassword == "" {
		return errors.Wrap(errors.New("bad request"), "old password and new password are required")
	}

	claims := GetClaims(c)
	req.Email = claims.Email

	res, err := h.svc.ChangePassword(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)

}

type (
	RefreshTokenReq struct {
		RefreshToken string `json:"refresh_token"`
	}

	RefreshTokenRes struct {
		AccessToken  string    `json:"access_token"`
		RefreshToken string    `json:"refresh_token"`
		ExpiresAt    time.Time `json:"expires_at"`
	}
)

func (h *AuthHandler) refreshToken(c echo.Context) error {
	var req RefreshTokenReq
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(errors.New("bad request"), err.Error())
	}

	if req.RefreshToken == "" {
		return errors.Wrap(errors.New("bad request"), "refresh token is required")
	}

	res, err := h.svc.RefreshToken(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

type (
	ProfileReq struct {
		UserID string `json:"-"`
	}

	ProfileRes struct {
		*auth.User
	}
)

func (h *AuthHandler) profile(c echo.Context) error {
	var req ProfileReq
	claims := GetClaims(c)
	req.UserID = claims.ID
	res, err := h.svc.GetProfile(c.Request().Context(), &req)
	if err != nil {
		return nil
	}
	return c.JSON(http.StatusOK, res)
}
