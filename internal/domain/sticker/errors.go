package sticker

import "errors"

var (
	ErrStickerNotFound       = errors.New("sticker not found")
	ErrStickerNotOwnedByUser = errors.New("sticker not owned by user")
)
