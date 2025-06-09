package auth

import "time"

type RefreshToken struct {
	UserID     string
	Token      string
	Valid      bool
	ValidUntil time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (rt *RefreshToken) Invalidate() {
	rt.Valid = false
}

func (rt *RefreshToken) IsValid() bool {
	return rt.Valid && time.Now().Before(rt.ValidUntil)
}

type PasswordResetToken struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

func (prt *PasswordResetToken) IsExpired() bool {
	return time.Now().After(prt.ExpiresAt)
}
