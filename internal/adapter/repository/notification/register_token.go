package notification

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/notification"
)

const registerTokenQuery = `
INSERT INTO user_push_tokens (id, user_id, push_token, platform, device_name, os_name, os_version, is_active, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT (push_token)
DO UPDATE SET
    user_id = EXCLUDED.user_id,
    platform = EXCLUDED.platform,
    device_name = EXCLUDED.device_name,
    os_name = EXCLUDED.os_name,
    os_version = EXCLUDED.os_version,
    is_active = EXCLUDED.is_active,
    updated_at = CURRENT_TIMESTAMP;
`

func (r *Repository) RegisterToken(ctx context.Context, token notification.Token) error {
	_, err := r.DB.Exec(ctx, registerTokenQuery,
		token.ID, token.UserID, token.PushToken, token.Platform, token.DeviceName, token.OSName, token.OSVersion, token.IsActive, token.CreatedAt, token.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
