package notification

import (
	"context"
	"github.com/ataberkcanitez/araqr/internal/application/domain/notification"
)

type Repository interface {
	RegisterToken(ctx context.Context, token notification.Token) error
}
