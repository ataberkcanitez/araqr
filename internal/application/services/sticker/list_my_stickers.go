package sticker

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
	"github.com/cockroachdb/errors"
	"sort"
)

func (svc *Service) ListMyStickers(ctx context.Context, req *web.ListMyStickersRequest) (*web.ListMyStickersResponse, error) {
	stickers, err := svc.stickerRepository.ListByUserID(ctx, req.UserID, req.Limit, req.Page)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list stickers for user")
	}
	if stickers == nil {
		return &web.ListMyStickersResponse{
			Stickers: []*sticker.Sticker{},
		}, nil

	}

	sort.SliceStable(stickers, func(i, j int) bool {
		return stickers[i].Active && !stickers[j].Active
	})

	return &web.ListMyStickersResponse{
		Stickers: stickers,
	}, nil
}
