package postgresrepository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/port"
	"github.com/typical-go/typical-rest-server/pkg/sqkit"
	"time"

	sq "github.com/Masterminds/squirrel"
	log "github.com/sirupsen/logrus"
)

type cartItemRepositoryImpl struct {
	pg *sql.DB
}



func NewCartItemRepository(pg *sql.DB) port.CartItemRepository {
	return &cartItemRepositoryImpl{pg}
}

func (c cartItemRepositoryImpl) Find(ctx context.Context, opts ...sqkit.SelectOption) (cartItems []*domain.CartItem, err error) {
	builder := sq.
		Select(
			domain.CartItemTable.ID,
			domain.CartItemTable.UserId,
			domain.CartItemTable.ProductId,
			domain.CartItemTable.Quantity,
			domain.CartItemTable.Date,
		).
		From(domain.CartItemTableName).
		PlaceholderFormat(sq.Dollar).
		RunWith(c.pg)


	for _, opt := range opts {
		builder = opt.CompileSelect(builder)
	}

	rows, err := builder.QueryContext(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	cartItems = make([]*domain.CartItem, 0)
	for  rows.Next(){
		cartItem := new(domain.CartItem)
		if err = rows.Scan(
			&cartItem.ID,
			&cartItem.UserID,
			&cartItem.ProductId,
			&cartItem.Quantity,
			&cartItem.Date,
		); err != nil {
			log.Error(err)
			return
		}
		cartItems = append(cartItems, cartItem)
	}
	return
}

func (c cartItemRepositoryImpl) Insert(ctx context.Context,  dbTx *sql.Tx, ent *domain.CartItem) (int, error) {
	builder := sq.
		Insert(domain.CartItemTableName).
		Columns(
			domain.CartItemTable.UserId,
			domain.CartItemTable.ProductId,
			domain.CartItemTable.Quantity,
			domain.CartItemTable.Date,
			).
		Suffix(
			fmt.Sprintf("RETURNING \"%s\"", domain.CartItemTable.ID),
		).
		PlaceholderFormat(sq.Dollar).
		Values(
			ent.UserID,
			ent.ProductId,
			ent.Quantity,
			time.Now(),
		)

	scanner := builder.RunWith(dbTx).QueryRowContext(ctx)
	var id int
	if err := scanner.Scan(&id); err != nil {
		log.Error(err)
		return -1, err
	}
	return id, nil

}

func (c cartItemRepositoryImpl) Delete(ctx context.Context, dbTx *sql.Tx, opts ...sqkit.DeleteOption) (int, error) {
	builder := sq.
		Delete(domain.CartItemTableName).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbTx)

	for _, opt := range opts {
		builder = opt.CompileDelete(builder)
	}

	res, err := builder.ExecContext(ctx)
	if err != nil {
		return -1, err
	}

	affectedRow, err := res.RowsAffected()
	return int(affectedRow), err
}
