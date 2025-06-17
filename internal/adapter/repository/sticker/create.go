package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
)

const insertStickersQuery = `
INSERT INTO stickers ( id, active, name, description, image_url, show_phone_number, phone_number,
show_email, email, show_instagram, instagram_url, show_facebook, facebook_url, user_id, created_at,
updated_at
) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16 );
`

func (r *Repository) Create(ctx context.Context, s *sticker.Sticker) (*sticker.Sticker, error) {
	_, err := r.DB.Exec(ctx, insertStickersQuery,
		s.ID,
		s.Active,
		s.Name,
		s.Description,
		s.ImageURL,
		s.ShowPhoneNumber,
		s.PhoneNumber,
		s.ShowEmail,
		s.Email,
		s.ShowInstagram,
		s.InstagramURL,
		s.ShowFacebook,
		s.FacebookURL,
		s.UserID,
		s.CreatedAt,
		s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return s, nil
}
