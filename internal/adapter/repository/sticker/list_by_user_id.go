package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
)

const ListStickersByUserIDQuery = `
SELECT 
s.id, 
    s.active, 
    s.name, 
    s.description, 
    s.image_url, 
    s.show_phone_number, 
    s.phone_number, 
    s.show_email, 
    s.email, 
    s.show_instagram, 
    s.instagram_url, 
    s.show_facebook, 
    s.facebook_url, 
    s.user_id, 
    s.created_at, 
    s.updated_at,
    COUNT(m.id) FILTER (WHERE m.read = false) AS unread_messages_count
FROM stickers s
LEFT JOIN messages m ON m.sticker_Id = s.id
WHERE s.user_id = $1
GROUP BY
    s.id, s.active, s.name, s.description, s.image_url,
    s.show_phone_number, s.phone_number,
    s.show_email, s.email,
    s.show_instagram, s.instagram_url,
    s.show_facebook, s.facebook_url,
    s.user_id, s.created_at, s.updated_at
LIMIT $2 OFFSET $3
`

func (r *Repository) ListByUserID(ctx context.Context, userID string, limit, page int) ([]*sticker.Sticker, error) {
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
			&s.UnreadMessageCount,
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
