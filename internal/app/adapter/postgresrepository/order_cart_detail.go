package postgresrepository

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/port"
	"github.com/typical-go/typical-rest-server/pkg/sqkit"
	"time"

	log "github.com/sirupsen/logrus"
)

type orderCartDetailRepositoryImpl struct {
	pg *sql.DB
}

func NewOrderCartDetailRepository(pg *sql.DB)  port.OrderCartDetailRepository{
	return &orderCartDetailRepositoryImpl{
		pg: pg,
	}
}

func (o orderCartDetailRepositoryImpl) Insert(ctx context.Context, dbTx *sql.Tx, ent *domain.OrderCartDetail) (int, error) {
	builder := sq.
		Insert(domain.OrderCartDetailTableName).
		Columns(
			domain.OrderCartDetailTable.ProductId,
			domain.OrderCartDetailTable.Quantity,
			domain.OrderCartDetailTable.Price,
			domain.OrderCartDetailTable.OrderId,
			domain.OrderCartDetailTable.Date,
		).
		Suffix(
			fmt.Sprintf("RETURNING \"%s\"", domain.OrderCartDetailTable.ID),
		).
		PlaceholderFormat(sq.Dollar).
		Values(
			ent.ProductId,
			ent.Quantity,
			ent.Price,
			ent.OrderId,
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

func (o orderCartDetailRepositoryImpl) Find(ctx context.Context, opts ...sqkit.SelectOption) (orderCartDetails []*domain.OrderCartDetail, err error) {
	panic("implement me")
}