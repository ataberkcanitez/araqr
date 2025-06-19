package sticker

import (
	"time"
)

type Sticker struct {
	ID                 string    `json:"id"`
	Active             bool      `json:"active"`
	Name               *string   `json:"name"`
	Description        *string   `json:"description"`
	ImageURL           *string   `json:"image_url"`
	ShowPhoneNumber    bool      `json:"show_phone_number"`
	PhoneNumber        *string   `json:"phone_number"`
	ShowEmail          bool      `json:"show_email"`
	Email              *string   `json:"email"`
	ShowInstagram      bool      `json:"show_instagram"`
	InstagramURL       *string   `json:"instagram_url"`
	ShowFacebook       bool      `json:"show_facebook"`
	FacebookURL        *string   `json:"facebook_url"`
	UserID             string    `json:"user_id"`
	UnreadMessageCount int       `json:"unread_message_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (s *Sticker) Assign(userID string) {
	s.UserID = userID
	s.UpdatedAt = time.Now()
}

func (s *Sticker) Update(name, description, imageURL *string, showPhoneNumber bool, phoneNumber *string, showEmail bool, email *string, showInstagram bool, instagramURL *string, showFacebook bool, facebookURL *string) {
	s.Name = name
	s.Description = description
	s.ImageURL = imageURL
	s.ShowPhoneNumber = showPhoneNumber
	s.PhoneNumber = phoneNumber
	s.ShowEmail = showEmail
	s.Email = email
	s.ShowInstagram = showInstagram
	s.InstagramURL = instagramURL
	s.ShowFacebook = showFacebook
	s.FacebookURL = facebookURL
	s.UpdatedAt = time.Now()
}
