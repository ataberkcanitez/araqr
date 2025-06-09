package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/handler"
	"github.com/ataberkcanitez/araqr/internal/domain/sticker"
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	"time"
)

func (svc *StickerService) Create(ctx context.Context, req *handler.CreateStickerRequest) ([]string, error) {
	now := time.Now()
	var stickerIds []string
	for i := 0; i < req.NumberOfStickers; i++ {
		sticker := &sticker.Sticker{
			ID:              uuid.NewString(),
			Active:          false,
			ShowPhoneNumber: false,
			ShowEmail:       false,
			ShowInstagram:   false,
			ShowFacebook:    false,
			CreatedAt:       now,
			UpdatedAt:       now,
		}
		_, err := svc.stickerRepository.Create(ctx, sticker)
		if err != nil {
			return []string{}, errors.Wrap(err, "failed to create sticker")
		}
		stickerIds = append(stickerIds, sticker.ID)
	}
	return stickerIds, nil
}
