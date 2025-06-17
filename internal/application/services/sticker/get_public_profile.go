package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
)

func (svc *Service) GetPublicProfile(ctx context.Context, request *web.GetStickerProfileRequest) (*web.GetStickerProfileResponse, error) {
	stx, err := svc.stickerRepository.GetByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if stx == nil {
		return nil, sticker.ErrStickerNotFound
	}

	publicProfile := &web.GetStickerProfileResponse{}
	publicProfile.ID = stx.ID
	publicProfile.Name = stx.Name
	publicProfile.Description = stx.Description
	publicProfile.ImageURL = stx.ImageURL
	if stx.ShowPhoneNumber {
		publicProfile.PhoneNumber = stx.PhoneNumber
	}
	if stx.ShowEmail {
		publicProfile.Email = stx.Email
	}
	if stx.ShowInstagram {
		publicProfile.InstagramURL = stx.InstagramURL
	}
	if stx.ShowFacebook {
		publicProfile.FacebookURL = stx.FacebookURL
	}

	return publicProfile, nil
}
