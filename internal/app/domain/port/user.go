package port

import (
	"context"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/typical-go/typical-rest-server/pkg/sqkit"
)

type UserRepository interface {
	Find(ctx context.Context, opts ...sqkit.SelectOption) (users []*domain.User, err error)
}
