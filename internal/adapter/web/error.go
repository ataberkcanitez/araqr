package web

import (
	"github.com/ataberkcanitez/araqr/internal/application/domain/auth"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
	"github.com/ataberkcanitez/araqr/log"
	"github.com/cockroachdb/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type httpErr struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func newHttpErr(code, msg string) *httpErr {
	return &httpErr{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

var ErrBadRequest = errors.New("bad request")

const ErrCodeBadRequest = "9000"

func ErrorHandler(err error, c echo.Context) {
	if err == nil {
		return
	}

	ctx := c.Request().Context()

	log.Error(ctx, "request failed, error handler captured", err)

	if errors.Is(err, ErrBadRequest) {
		_ = c.JSON(400, newHttpErr(ErrCodeBadRequest, err.Error()))
		return
	}

	if handleAuthErr(err, c) {
		return
	}

	if handleStickerErr(err, c) {
		return
	}

	_ = c.JSON(http.StatusInternalServerError, newHttpErr("-1", err.Error()))

}

var (
	ErrCodeUserNotFound             = "1000"
	ErrCodeUserAlreadyExists        = "1001"
	ErrCodeAuthBearerHeaderNotFound = "1002"
	ErrCodeAuthInvalidToken         = "1003"
	ErrCodeResetTokenNotFound       = "1004"
	ErrCodeResetTokenExpired        = "1006"
	ErrCodeSamePassword             = "1007"
	ErrCodePasswordMismatch         = "1008"
	ErrInvalidPassword              = "1009"
)

// handleAuthErr sends the auth error to the client
// returns true if the error is handled by this function otherwise false
func handleAuthErr(err error, c echo.Context) bool {
	if errors.Is(err, auth.ErrAuthBearerHeaderNotFound) {
		_ = c.JSON(http.StatusUnauthorized, newHttpErr(ErrCodeAuthBearerHeaderNotFound, err.Error()))
	} else if errors.Is(err, auth.ErrInvalidToken) {
		_ = c.JSON(http.StatusUnauthorized, newHttpErr(ErrCodeAuthInvalidToken, err.Error()))
	} else if errors.Is(err, auth.ErrUserAlreadyExists) {
		_ = c.JSON(http.StatusConflict, newHttpErr(ErrCodeUserAlreadyExists, err.Error()))
	} else if errors.Is(err, auth.ErrUserNotFound) {
		_ = c.JSON(http.StatusNotFound, newHttpErr(ErrCodeUserNotFound, err.Error()))
	} else if errors.Is(err, auth.ErrResetTokenNotFound) {
		_ = c.JSON(http.StatusNotFound, newHttpErr(ErrCodeResetTokenNotFound, err.Error()))
	} else if errors.Is(err, auth.ErrResetTokenExpired) {
		_ = c.JSON(http.StatusBadRequest, newHttpErr(ErrCodeResetTokenExpired, err.Error()))
	} else if errors.Is(err, auth.ErrSamePassword) {
		_ = c.JSON(http.StatusBadRequest, newHttpErr(ErrCodeSamePassword, err.Error()))
	} else if errors.Is(err, auth.ErrPasswordMismatch) {
		_ = c.JSON(http.StatusBadRequest, newHttpErr(ErrCodePasswordMismatch, err.Error()))
	} else if errors.Is(err, auth.ErrInvalidPassword) {
		_ = c.JSON(http.StatusBadRequest, newHttpErr(ErrInvalidPassword, err.Error()))
	} else {
		return false
	}

	return true
}

var (
	ErrCodeStickerNotFound        = "2000"
	ErrCodeStickerNotOwnedByUser  = "2001"
	ErrCodeStickerAlreadyAssigned = "2002"
)

func handleStickerErr(err error, c echo.Context) bool {
	if errors.Is(err, sticker.ErrStickerNotFound) {
		_ = c.JSON(http.StatusBadRequest, newHttpErr(ErrCodeStickerNotFound, err.Error()))
	} else if errors.Is(err, sticker.ErrStickerNotOwnedByUser) {
		_ = c.JSON(http.StatusUnauthorized, newHttpErr(ErrCodeStickerNotOwnedByUser, err.Error()))
	} else if errors.Is(err, sticker.ErrStickerAlreadyAssigned) {
		_ = c.JSON(http.StatusConflict, newHttpErr(ErrCodeStickerAlreadyAssigned, err.Error()))
	} else {
		return false
	}

	return true
}
