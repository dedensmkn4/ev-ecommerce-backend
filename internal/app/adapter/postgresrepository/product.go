package postgresrepository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/port"
	"github.com/dedensmkn4/ev-ecommerce-backend/pkg/sqkit"
	log "github.com/sirupsen/logrus"
)

type productRepositoryImpl struct {
	pg *sql.DB
}



func NewProductRepository(pg *sql.DB) port.ProductRepository {
	return &productRepositoryImpl{pg}
}

func (p productRepositoryImpl) Find(ctx context.Context, opts ...sqkit.SelectOption) (products []*domain.Product, err error) {

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
		PlaceholderFormat(sq.Dollar).
		RunWith(p.pg)


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
	}

	return
}
