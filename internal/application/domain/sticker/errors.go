package sticker

import "errors"

var (
	ErrStickerNotFound        = errors.New("sticker not found")
	ErrStickerNotOwnedByUser  = errors.New("sticker not owned by user")
	ErrStickerAlreadyAssigned = errors.New("sticker already assigned to a user")

	ErrMessageNotFound = errors.New("message not found")
)
