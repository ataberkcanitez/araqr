package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
	"time"
)

func (s *Service) UpdateSticker(ctx context.Context, req *web.UpdateMyStickerRequest) (*sticker.Sticker, error) {
	stx, err := s.Get(ctx, &web.GetStickerRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}
	if stx == nil {
		return nil, sticker.ErrStickerNotFound
	}
	if stx.UserID != req.UserID {
		return nil, sticker.ErrStickerNotOwnedByUser
	}

	stx.Active = req.Active
	stx.Name = req.Name
	stx.Description = req.Description
	stx.ShowPhoneNumber = req.ShowPhoneNumber
	stx.PhoneNumber = req.PhoneNumber
	stx.ShowEmail = req.ShowEmail
	stx.Email = req.Email
	stx.ShowInstagram = req.ShowInstagram
	stx.InstagramURL = req.InstagramURL
	stx.ShowFacebook = req.ShowFacebook
	stx.FacebookURL = req.FacebookURL
	stx.UpdatedAt = time.Now()

	if err := s.stickerRepository.Update(ctx, stx); err != nil {
		return nil, err
	}
	return stx, nil
}
