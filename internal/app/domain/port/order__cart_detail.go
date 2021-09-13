package port

import (
	"context"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/pkg/sqkit"
)

type OrderCartDetailRepository interface {
	Find(ctx context.Context, opts ...sqkit.SelectOption) (orderCartDetails []*domain.OrderCartDetail, err error)
}

