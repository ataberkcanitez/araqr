package sticker

import (
	"context"
	"fmt"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v4"
)

const GetStickerQuery = `
SELECT 
	id, 
	active, 
	name, 
	description, 
	image_url, 
	show_phone_number, 
	phone_number, 
	show_email, 
	email, 
	show_instagram, 
	instagram_url, 
	show_facebook, 
	facebook_url, 
	user_id, 
	created_at, 
	updated_at
FROM stickers WHERE id = $1
`

func (r *Repository) GetByID(ctx context.Context, id string) (*sticker.Sticker, error) {
	fmt.Println("Executing GetByID with ID:", id)
	var s sticker.Sticker
	err := r.DB.QueryRow(ctx, GetStickerQuery, id).Scan(
		&s.ID,
		&s.Active,
		&s.Name,
		&s.Description,
		&s.ImageURL,
		&s.ShowPhoneNumber,
		&s.PhoneNumber,
		&s.ShowEmail,
		&s.Email,
		&s.ShowInstagram,
		&s.InstagramURL,
		&s.ShowFacebook,
		&s.FacebookURL,
		&s.UserID,
		&s.CreatedAt,
		&s.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, sticker.ErrStickerNotFound
		}
		return nil, err
	}
	return &s, nil
}
