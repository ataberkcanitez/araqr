package sticker

import "time"

type Message struct {
	ID           string
	StickerID    string
	Message      string
	UrgencyLevel string
	CreatedAt    time.Time
}
