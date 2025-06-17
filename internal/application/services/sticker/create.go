package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"time"
)

func (svc *Service) Create(ctx context.Context, req *web.CreateStickerRequest) ([]string, error) {
	now := time.Now()
	var stickerIds []string
	for i := 0; i < req.NumberOfStickers; i++ {
		stx := &sticker.Sticker{
			ID:              uuid.NewString(),
			Active:          false,
			ShowPhoneNumber: false,
			ShowEmail:       false,
			ShowInstagram:   false,
			ShowFacebook:    false,
			CreatedAt:       now,
			UpdatedAt:       now,
		}
		_, err := svc.stickerRepository.Create(ctx, stx)
		if err != nil {
			return []string{}, errors.Wrap(err, "failed to create sticker")
		}
		stickerIds = append(stickerIds, stx.ID)
	}
	return stickerIds, nil
}
