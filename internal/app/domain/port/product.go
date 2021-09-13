package port

import (
	"context"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/pkg/sqkit"
)

type ProductUseCase interface {
	FindById(ctx context.Context, paramId string) (*domain.Product, error)
}

type ProductRepository interface {
	Find(ctx context.Context, opts ...sqkit.SelectOption) (products []*domain.Product, err error)
}