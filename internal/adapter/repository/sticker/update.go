package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
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
    show_email = $7,
    email = $8,
    show_instagram = $9,
    instagram_url = $10,
    show_facebook = $11,
    facebook_url = $12,
    user_id = $13,
    updated_at = $14
WHERE id = $15;
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
		stx.ShowEmail,
		stx.Email,
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
