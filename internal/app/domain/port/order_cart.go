package port

import (
	"context"
	"database/sql"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
)

type OrderCartUseCase interface {
	AddCart(ctx context.Context, param domain.AddCartItemPayload) (cartResponse *domain.AddCartItemResponse, err error)
	Checkout(ctx context.Context, param domain.OrderCheckoutPayload) (checkoutResponse *domain.OrderCheckoutResponse, err error)
}

type OrderCartRepository interface {
	Insert(ctx context.Context, dbTx *sql.Tx,ent *domain.OrderCart) (int, error)
}