package postgresrepository

import (
	"context"
	"database/sql"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/port"
)

type orderCartRepositoryImpl struct {
	pg *sql.DB
}


func NewOrderCartRepository(pg *sql.DB)  port.OrderCartRepository{
	return &orderCartRepositoryImpl{
		pg: pg,
	}
}

func (o orderCartRepositoryImpl) CreateCart(ctx context.Context, dbTx *sql.Tx, userID int) (cartID int, err error) {
	panic("implement me")
}

func (o orderCartRepositoryImpl) InsertCartItems(ctx context.Context, dbTx *sql.Tx, cartID int, productCode string, Qty int) error {
	panic("implement me")
}

func (o orderCartRepositoryImpl) CheckUserByID(ctx context.Context, dbTx *sql.Tx, userID int) (int, error) {
	panic("implement me")
}

func (o orderCartRepositoryImpl) GetCartData(ctx context.Context, dbTx *sql.Tx, cartID int, isRowLocking bool) (domain.CartData, error) {
	panic("implement me")
}

func (o orderCartRepositoryImpl) CreateOrders(ctx context.Context, dbTx *sql.Tx, userID int) (int, error) {
	panic("implement me")
}

func (o orderCartRepositoryImpl) InsertOrderItems(ctx context.Context, dbTx *sql.Tx, orderID int, productCode string, qty int) error {
	panic("implement me")
}

func (o orderCartRepositoryImpl) DeleteCart(ctx context.Context, dbTx *sql.Tx, cartID int) error {
	panic("implement me")
}

func (o orderCartRepositoryImpl) CleanCartAndOrders(ctx context.Context) error {
	panic("implement me")
}

