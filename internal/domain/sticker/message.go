package sticker

import "time"

type Message struct {
	ID           string    `json:"id"`
	StickerID    string    `json:"sticker_id"`
	Message      string    `json:"message"`
	UrgencyLevel string    `json:"urgency_level"`
	CreatedAt    time.Time `json:"created_at"`
}
