package sticker

import "time"

type Message struct {
	ID           string    `json:"id"`
	StickerID    string    `json:"sticker_id"`
	Message      string    `json:"message"`
	UrgencyLevel string    `json:"urgency_level"`
	Read         bool      `json:"read"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}

func (m *Message) SetAsRead() {
	m.Read = true
	m.UpdatedAt = time.Now()
}
