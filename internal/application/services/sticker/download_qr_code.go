package sticker

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ataberkcanitez/araqr/internal/adapter/web"
	"github.com/ataberkcanitez/araqr/internal/application/domain/sticker"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
)

func (s *Service) CreateQrCode(ctx context.Context, req *web.DownloadQRCodeRequest) ([]byte, error) {
	stx, err := s.Get(ctx, &web.GetStickerRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}
	if stx == nil {
		return nil, sticker.ErrStickerNotFound
	}

	url := fmt.Sprintf("https://www.araqr.com/qr/%s", stx.ID)
	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		return nil, fmt.Errorf("failed to create QR code: %w", err)
	}
	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		return nil, fmt.Errorf("failed to scale QR code: %w", err)
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, qrCode)
	if err != nil {
		return nil, fmt.Errorf("failed to encode QR code to PNG: %w", err)
	}

	return buf.Bytes(), nil
}
