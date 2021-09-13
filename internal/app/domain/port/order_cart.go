package port

import (
	"context"
	"database/sql"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
)

type OrderCartUseCase interface {
	AddCart(ctx context.Context, param domain.AddToCartPayload) (cartResponse domain.AddToCartResponse, err error)
	Checkout(ctx context.Context, param domain.CheckoutPayload) (checkoutResponse domain.CheckoutResponse, err error)
}

type OrderCartRepository interface {
	CreateCart(ctx context.Context, dbTx *sql.Tx, userID int) (cartID int, err error)
	InsertCartItems(ctx context.Context, dbTx *sql.Tx, cartID int, productCode string, Qty int) error
	CheckUserByID(ctx context.Context, dbTx *sql.Tx, userID int) (int, error)
	GetCartData(ctx context.Context, dbTx *sql.Tx, cartID int, isRowLocking bool) (domain.CartData, error)
	CreateOrders(ctx context.Context, dbTx *sql.Tx, userID int) (int, error)
	InsertOrderItems(ctx context.Context, dbTx *sql.Tx, orderID int, productCode string, qty int) error
	DeleteCart(ctx context.Context, dbTx *sql.Tx, cartID int) error
	CleanCartAndOrders(ctx context.Context) error
}