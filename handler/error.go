package handler

import (
	"github.com/ataberkcanitez/araqr/internal/domain"
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

	log.Error(ctx, "reqest failed, error handler captured", err)

	if errors.Is(err, ErrBadRequest) {
		_ = c.JSON(400, newHttpErr(ErrCodeBadRequest, err.Error()))
		return
	}

	if handleAuthErr(err, c) {
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
	if errors.Is(err, domain.ErrAuthBearerHeaderNotFound) {
		_ = c.JSON(http.StatusUnauthorized, newHttpErr(ErrCodeAuthBearerHeaderNotFound, err.Error()))
	} else if errors.Is(err, domain.ErrInvalidToken) {
		_ = c.JSON(http.StatusUnauthorized, newHttpErr(ErrCodeAuthInvalidToken, err.Error()))
	} else if errors.Is(err, domain.ErrUserAlreadyExists) {
		_ = c.JSON(http.StatusConflict, newHttpErr(ErrCodeUserAlreadyExists, err.Error()))
	} else if errors.Is(err, domain.ErrUserNotFound) {
		_ = c.JSON(http.StatusNotFound, newHttpErr(ErrCodeUserNotFound, err.Error()))
	} else if errors.Is(err, domain.ErrResetTokenNotFound) {
		_ = c.JSON(http.StatusNotFound, newHttpErr(ErrCodeResetTokenNotFound, err.Error()))
	} else if errors.Is(err, domain.ErrResetTokenExpired) {
		_ = c.JSON(http.StatusBadRequest, newHttpErr(ErrCodeResetTokenExpired, err.Error()))
	} else if errors.Is(err, domain.ErrSamePassword) {
		_ = c.JSON(http.StatusBadRequest, newHttpErr(ErrCodeSamePassword, err.Error()))
	} else if errors.Is(err, domain.ErrPasswordMismatch) {
		_ = c.JSON(http.StatusBadRequest, newHttpErr(ErrCodePasswordMismatch, err.Error()))
	} else if errors.Is(err, domain.ErrInvalidPassword) {
		_ = c.JSON(http.StatusBadRequest, newHttpErr(ErrInvalidPassword, err.Error()))
	} else {
		return false
	}

	return true
}
