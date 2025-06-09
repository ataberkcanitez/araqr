package sticker

import (
	"context"
	"fmt"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
)

const ListStickersByUserIDQuery = `
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
FROM stickers
WHERE user_id = $1
LIMIT $2 OFFSET $3
`

func (r *Repository) ListByUserID(ctx context.Context, userID string, limit, page int) ([]*sticker.Sticker, error) {
	fmt.Println("page:", page, "limit:", limit, "userID:", userID)
	rows, err := r.DB.Query(ctx, ListStickersByUserIDQuery, userID, limit, page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stickers []*sticker.Sticker
	for rows.Next() {
		var s sticker.Sticker
		err := rows.Scan(
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
			return nil, err
		}
		stickers = append(stickers, &s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stickers, nil
}
