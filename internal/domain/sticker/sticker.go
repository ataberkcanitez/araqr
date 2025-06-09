package sticker

import (
	"time"
)

type Sticker struct {
	ID              string
	Active          bool
	Name            *string
	Description     *string
	ImageURL        *string
	ShowPhoneNumber bool
	PhoneNumber     *string
	ShowEmail       bool
	Email           *string
	ShowInstagram   bool
	InstagramURL    *string
	ShowFacebook    bool
	FacebookURL     *string
	UserID          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
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
