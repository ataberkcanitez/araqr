package sticker

import (
	"context"

	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
)

const UpdateStickerQuery = `
UPDATE stickers
SET
    active = $1,
    name = $2,
    description = $3,
    image_url = $4,
    show_phone_number = $5,
    phone_number = $6,
    show_instagram = $7,
    instagram_url = $8,
    show_facebook = $9,
    facebook_url = $10,
    user_id = $11,
    updated_at = $12
WHERE id = $13;
`

func (r *Repository) Update(ctx context.Context, stx *sticker.Sticker) error {
	_, err := r.DB.Exec(
		ctx,
		UpdateStickerQuery,
		stx.Active,
		stx.Name,
		stx.Description,
		stx.ImageURL,
		stx.ShowPhoneNumber,
		stx.PhoneNumber,
		stx.ShowInstagram,
		stx.InstagramURL,
		stx.ShowFacebook,
		stx.FacebookURL,
		stx.UserID,
		stx.UpdatedAt,
		stx.ID,
	)
	return err
}
