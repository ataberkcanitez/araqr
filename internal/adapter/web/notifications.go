package web

import (
	"context"
	"github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type notificationSvc interface {
	RegisterToken(ctx context.Context, req *RegisterTokenReq) (*RegisterTokenRes, error)
}

type NotificationHandler struct {
	svc notificationSvc
	tp  tokenParser
}

func NewNotificationHandler(svc notificationSvc, parser tokenParser) *NotificationHandler {
	return &NotificationHandler{
		svc: svc,
		tp:  parser,
	}
}

func (h *NotificationHandler) RegisterRoutes(e *echo.Echo) {
	basePrefix := "/api/v1/notifications"
	private := e.Group(basePrefix, MiddlewareTokenVerification(h.tp))

	private.POST("/register-token", h.registerToken)
}

type (
	RegisterTokenReq struct {
		PushToken  string     `json:"push_token"`
		Platform   string     `json:"platform"` // "ios" or "android"
		DeviceInfo DeviceInfo `json:"device_info"`

		UserID string
	}

	DeviceInfo struct {
		DeviceName string `json:"deviceName"`
		OSName     string `json:"osName"`
		OSVersion  string `json:"osVersion"`
	}

	RegisterTokenRes struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

func (h *NotificationHandler) registerToken(c echo.Context) error {
	var req RegisterTokenReq
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(errors.New("bad request"), err.Error())
	}

	if req.PushToken == "" {
		return c.JSON(400, RegisterTokenRes{Success: false, Message: "push_token is required"})
	}
	if req.Platform != "ios" && req.Platform != "android" {
		return c.JSON(400, RegisterTokenRes{Success: false, Message: "platform must be 'ios' or 'android'"})
	}
	if req.DeviceInfo.DeviceName == "" || req.DeviceInfo.OSName == "" || req.DeviceInfo.OSVersion == "" {
		return c.JSON(400, RegisterTokenRes{Success: false, Message: "device_info fields are required"})
	}

	claims := GetClaims(c)
	req.UserID = claims.ID

	res, err := h.svc.RegisterToken(c.Request().Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
