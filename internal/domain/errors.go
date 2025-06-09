package domain

import "errors"

var (
	// ErrInvalidToken is returned when the token is invalid
	ErrInvalidToken = errors.New("invalid token")
	// ErrInvalidTokenSigningMethod is returned when the token signing method is invalid
	ErrInvalidTokenSigningMethod = errors.New("invalid token signing method")
	// ErrInvalidPassword is returned when the password is invalid
	ErrInvalidPassword = errors.New("invalid password")
	// ErrPasswordMismatch is returned when the password does not match
	ErrPasswordMismatch = errors.New("password mismatch")
	// ErrUserAlreadyExists is returned when the user already exists
	ErrUserAlreadyExists = errors.New("user already exists")
	// ErrUserNotFound is returned when the user is not found
	ErrUserNotFound = errors.New("user not found")
	// ErrAuthBearerHeaderNotFound is returned when the Authorization header is not found
	ErrAuthBearerHeaderNotFound = errors.New("authorization bearer header not found")
	// ErrResetTokenNotFound is returned when the user is not found
	ErrResetTokenNotFound = errors.New("reset token not found")
	// ErrResetTokenExpired is returned when the reset token is expired
	ErrResetTokenExpired = errors.New("reset token is expired")
	// ErrSamePassword is returned when the user attempts to update their password but provides the same password as the current one.
	ErrSamePassword = errors.New("same password")
)
