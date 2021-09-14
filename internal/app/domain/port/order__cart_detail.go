package port

import (
	"context"
	"database/sql"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/typical-go/typical-rest-server/pkg/sqkit"
)

type OrderCartDetailRepository interface {
	Find(ctx context.Context, opts ...sqkit.SelectOption) (orderCartDetails []*domain.OrderCartDetail, err error)
	Insert(ctx context.Context,  dbTx *sql.Tx, ent *domain.OrderCartDetail) (int, error)
}

