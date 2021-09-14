package postgresrepository

import (
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/port"
	log "github.com/sirupsen/logrus"
	"time"
)

type orderCartRepositoryImpl struct {
	pg *sql.DB
}

func NewOrderCartRepository(pg *sql.DB)  port.OrderCartRepository{
	return &orderCartRepositoryImpl{
		pg: pg,
	}
}

func (o orderCartRepositoryImpl) Insert(ctx context.Context, dbTx *sql.Tx, ent *domain.OrderCart) (int, error) {
	builder := sq.
		Insert(domain.OrderCartTableName).
		Columns(
			domain.OrderCartTable.UserId,
			domain.OrderCartTable.TotalPrice,
			domain.OrderCartTable.OrderStatus,
			domain.OrderCartTable.Date,
			domain.OrderCartTable.TimeLimitPayment,
		).
		Suffix(
			fmt.Sprintf("RETURNING \"%s\"", domain.OrderCartTable.ID),
		).
		PlaceholderFormat(sq.Dollar).
		Values(
			ent.UserId,
			ent.TotalPrice,
			ent.OrderStatus,
			time.Now(),
			time.Now().Add(time.Minute*60),
		)

	scanner := builder.RunWith(dbTx).QueryRowContext(ctx)
	var id int
	if err := scanner.Scan(&id); err != nil {
		log.Error(err)
		return -1, err
	}
	return id, nil
}
