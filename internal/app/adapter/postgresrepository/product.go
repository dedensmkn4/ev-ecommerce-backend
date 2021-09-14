package postgresrepository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/port"
	"github.com/dedensmkn4/ev-ecommerce-backend/pkg/pgkit"
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-rest-server/pkg/sqkit"
)

type productRepositoryImpl struct {
	pg *sql.DB
}


func NewProductRepository(pg *sql.DB) port.ProductRepository {
	return &productRepositoryImpl{pg}
}

func (p productRepositoryImpl) Find(ctx context.Context, dbTx *sql.Tx, isRowLocking bool, opts ...sqkit.SelectOption) (products []*domain.Product, err error) {

	builder := sq.
		Select(
			domain.ProductTable.ID,
			domain.ProductTable.Code,
			domain.ProductTable.Name,
			domain.ProductTable.Desc,
			domain.ProductTable.Stock,
			domain.ProductTable.Price,
			).
		From(domain.ProductTableName).
		Suffix(pgkit.SuffixRawLocking(isRowLocking)).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbTx)

	for _, opt := range opts {
		builder = opt.CompileSelect(builder)
	}

	rows, err := builder.QueryContext(ctx)
	if err != nil {
		log.Error(err)
		return
	}
	products = make([]*domain.Product, 0)
	for  rows.Next(){
		product := new(domain.Product)
		if err = rows.Scan(
			&product.ID,
			&product.Code,
			&product.Name,
			&product.Desc,
			&product.Stock,
			&product.Price,
		); err != nil {
			log.Error(err)
			return
		}
		products = append(products, product)
	}

	return
}

func (p productRepositoryImpl) Update(ctx context.Context, dbTx *sql.Tx, ent *domain.Product, opts ...sqkit.UpdateOption) (int, error) {
	builder := sq.
		Update(domain.ProductTableName).
		Set(domain.ProductTable.Price, ent.Price).
		Set(domain.ProductTable.Stock, ent.Stock).
		Set(domain.ProductTable.Code, ent.Code).
		Set(domain.ProductTable.Name, ent.Name).
		Set(domain.ProductTable.Desc, ent.Desc).
		PlaceholderFormat(sq.Dollar).
		RunWith(dbTx)

	for _, opt := range opts {
		builder = opt.CompileUpdate(builder)
	}

	res, err := builder.ExecContext(ctx)
	if err != nil {
		log.Error(err)
		return -1, err
	}
	affectedRow, err := res.RowsAffected()
	return int(affectedRow), err
}
