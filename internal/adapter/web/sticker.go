package web

import (
	"context"
	"fmt"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
	"github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type stickerSvc interface {
	Create(ctx context.Context, req *CreateStickerRequest) ([]string, error)
	Assign(ctx context.Context, req *AssignStickerRequest) (*sticker.Sticker, error)
	GetPublicProfile(ctx context.Context, req *GetStickerProfileRequest) (*GetStickerProfileResponse, error)
	SendMessageToSticker(ctx context.Context, s *SendMessageToStickerRequest) (*SendMessageToStickerResponse, error)
	ListMyStickers(ctx context.Context, l *ListMyStickersRequest) (*ListMyStickersResponse, error)
	Get(ctx context.Context, req *GetStickerRequest) (*sticker.Sticker, error)
	UpdateSticker(ctx context.Context, u *UpdateMyStickerRequest) (*sticker.Sticker, error)
	ListMessages(ctx context.Context, l *ListMessagesRequest) ([]*sticker.Message, error)
	CreateQrCode(ctx context.Context, d *DownloadQRCodeRequest) ([]byte, error)
}

type StickerHandler struct {
	svc stickerSvc
	tp  tokenParser
}

func NewStickerHandler(svc stickerSvc, parser tokenParser) *StickerHandler {
	return &StickerHandler{
		svc: svc,
		tp:  parser,
	}
}

func (h *StickerHandler) RegisterRoutes(e *echo.Echo) {
	// Public
	// Get Car Owner sticker profile
	e.GET("/public/v1/stickers/:id", h.GetStickerProfile)
	// Send message to sticker
	e.POST("/public/v1/stickers/:id/message", h.SendMessageToSticker)

	// Create a new sticker
	e.POST("/v1/sticker", h.Create, MiddlewareTokenVerification(h.tp))
	// Assign a sticker to a user
	e.POST("/v1/sticker/:stickerID/assign", h.AssignSticker, MiddlewareTokenVerification(h.tp))
	// List my stickers
	e.GET("/v1/stickers", h.ListMyStickers, MiddlewareTokenVerification(h.tp))
	// Get Specific sticker
	e.GET("/v1/stickers/:id", h.GetSticker, MiddlewareTokenVerification(h.tp))
	// Update sticker settings
	e.PUT("/v1/stickers/:id", h.UpdateMySticker, MiddlewareTokenVerification(h.tp))
	// Get sticker messages
	e.GET("/v1/sticker/:id/messages", h.ListMessages, MiddlewareTokenVerification(h.tp))

	// Download QR code for sticker
	e.GET("/v1/stickers/:id/download", h.DownloadQRCode, MiddlewareTokenVerification(h.tp))
}

type (
	CreateStickerRequest struct {
		NumberOfStickers int `json:"number_of_stickers"`
	}
)

func (h *StickerHandler) Create(c echo.Context) error {
	var req CreateStickerRequest
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(ErrBadRequest, err.Error())
	}

	fmt.Println(req.NumberOfStickers)

	if req.NumberOfStickers <= 0 {
		return errors.Wrap(ErrBadRequest, "number_of_stickers must be greater than 0")
	}

	stickerIds, err := h.svc.Create(c.Request().Context(), &req)
	if err != nil {
		return errors.Wrap(err, "failed to create sticker")
	}
	return c.JSON(http.StatusOK, stickerIds)
}

type (
	AssignStickerRequest struct {
		StickerID string `param:"stickerID"`
		UserID    string
	}
)

func (h *StickerHandler) AssignSticker(c echo.Context) error {
	var req AssignStickerRequest
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(ErrBadRequest, err.Error())
	}

	claims := GetClaims(c)
	req.UserID = claims.ID
	_, err := h.svc.Assign(c.Request().Context(), &req)
	return err
}

type (
	GetStickerProfileRequest struct {
		ID string `param:"id"`
	}

	GetStickerProfileResponse struct {
		ID           string  `json:"id"`
		Name         *string `json:"name"`
		Description  *string `json:"description"`
		ImageURL     *string `json:"image_url"`
		PhoneNumber  *string `json:"phone_number"`
		Email        *string `json:"email"`
		InstagramURL *string `json:"instagram_url"`
		FacebookURL  *string `json:"facebook_url"`
	}
)

// GetStickerProfile is a public API that retrieves the sticker profile by ID.
func (h *StickerHandler) GetStickerProfile(c echo.Context) error {
	var req GetStickerProfileRequest
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(ErrBadRequest, err.Error())
	}

	publicProfile, err := h.svc.GetPublicProfile(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, publicProfile)
}

type (
	SendMessageToStickerRequest struct {
		ID           string `param:"id"`
		UrgencyLevel string `json:"urgency_level"`
		Message      string `json:"message"`
	}
	SendMessageToStickerResponse struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	}
)

func (h *StickerHandler) SendMessageToSticker(c echo.Context) error {
	var req SendMessageToStickerRequest
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(ErrBadRequest, err.Error())
	}

	if req.UrgencyLevel == "" || req.Message == "" {
		return errors.Wrap(ErrBadRequest, "urgency_level and message are required")
	}

	res, err := h.svc.SendMessageToSticker(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)

}

type (
	ListMyStickersRequest struct {
		UserID string
		Page   int `query:"page" default:"1"`
		Limit  int `query:"limit" default:"100"`
	}
	ListMyStickersResponse struct {
		Stickers []*sticker.Sticker `json:"stickers"`
	}
)

func (h *StickerHandler) ListMyStickers(c echo.Context) error {
	var req ListMyStickersRequest
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(ErrBadRequest, err.Error())
	}

	if req.Limit == 0 {
		req.Limit = 100
	}

	claims := GetClaims(c)
	req.UserID = claims.ID
	res, err := h.svc.ListMyStickers(c.Request().Context(), &req)
	if err != nil {
		return errors.Wrap(err, "failed to list my stickers")
	}
	return c.JSON(http.StatusOK, res)

}

type (
	GetStickerRequest struct {
		ID string `param:"id"`
	}
)

func (h *StickerHandler) GetSticker(c echo.Context) error {
	var req GetStickerRequest
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(ErrBadRequest, err.Error())
	}

	_ = GetClaims(c)

	stx, err := h.svc.Get(c.Request().Context(), &req)
	if err != nil {
		return errors.Wrap(err, "failed to get sticker")
	}

	return c.JSON(http.StatusOK, stx)
}

type (
	UpdateMyStickerRequest struct {
		ID string `param:"id"`

		Active          bool    `json:"active"`
		Name            *string `json:"name"`
		Description     *string `json:"description"`
		ShowPhoneNumber bool    `json:"show_phone_number"`
		PhoneNumber     *string `json:"phone_number"`
		ShowEmail       bool    `json:"show_email"`
		Email           *string `json:"email"`
		ShowInstagram   bool    `json:"show_instagram"`
		InstagramURL    *string `json:"instagram_url"`
		ShowFacebook    bool    `json:"show_facebook"`
		FacebookURL     *string `json:"facebook_url"`

		UserID string `json:"-"`
	}
)

func (h *StickerHandler) UpdateMySticker(c echo.Context) error {
	var req UpdateMyStickerRequest
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(ErrBadRequest, err.Error())
	}
	claims := GetClaims(c)
	req.UserID = claims.ID

	stx, err := h.svc.UpdateSticker(c.Request().Context(), &req)
	if err != nil {
		return errors.Wrap(err, "failed to update sticker")
	}

	return c.JSON(http.StatusOK, stx)
}

type (
	ListMessagesRequest struct {
		ID string `param:"id"`

		Page  int `query:"page" default:"1"`
		Limit int `query:"limit" default:"100"`

		UserID string
	}
)

func (h *StickerHandler) ListMessages(c echo.Context) error {
	var req ListMessagesRequest
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(ErrBadRequest, err.Error())
	}

	if req.Limit == 0 {
		req.Limit = 100
	}

	claims := GetClaims(c)
	req.UserID = claims.ID

	messages, err := h.svc.ListMessages(c.Request().Context(), &req)
	if err != nil {
		return errors.Wrap(err, "failed to list sticker messages")
	}
	return c.JSON(http.StatusOK, messages)
}

type (
	DownloadQRCodeRequest struct {
		ID string `param:"id"`
	}
)

func (h *StickerHandler) DownloadQRCode(c echo.Context) error {
	var req DownloadQRCodeRequest
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(ErrBadRequest, err.Error())
	}

	if req.ID == "" {
		return errors.Wrap(ErrBadRequest, "missing sticker id")
	}

	qrCode, err := h.svc.CreateQrCode(c.Request().Context(), &req)
	if err != nil {
		return errors.Wrap(err, "failed to create QR code")
	}
	c.Response().Header().Set(echo.HeaderContentType, "image/png")
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=sticker-"+req.ID+"-qrcode.png")
	return c.Blob(http.StatusOK, "image/png", qrCode)
}
