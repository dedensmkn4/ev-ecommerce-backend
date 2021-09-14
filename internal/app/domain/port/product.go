package port

import (
	"context"
	"database/sql"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/typical-go/typical-rest-server/pkg/sqkit"
)

type ProductUseCase interface {
	GetAll(ctx context.Context, param domain.FindProductFilter) ([]*domain.Product, error)
	FindById(ctx context.Context, paramId string) (*domain.Product, error)
}

type ProductRepository interface {
	Find(ctx context.Context, dbTx *sql.Tx, isRowLocking bool, opts ...sqkit.SelectOption) (products []*domain.Product, err error)
	Update(ctx context.Context, dbTx *sql.Tx, ent *domain.Product, opts ...sqkit.UpdateOption) (int, error)
}