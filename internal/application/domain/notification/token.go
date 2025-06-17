package notification

import "time"

type Token struct {
	ID         string    `db:"id"`
	UserID     string    `db:"user_id"`
	PushToken  string    `db:"push_token"`
	Platform   string    `db:"platform"`
	DeviceName string    `db:"device_name"`
	OSName     string    `db:"os_name"`
	OSVersion  string    `db:"os_version"`
	IsActive   bool      `db:"is_active"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
