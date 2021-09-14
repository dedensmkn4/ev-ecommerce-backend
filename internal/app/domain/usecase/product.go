package usecase

import (
	"context"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/port"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/infra/postgresdb"
	"github.com/labstack/echo/v4"
	"github.com/typical-go/typical-rest-server/pkg/sqkit"
	"strconv"
	"strings"
)

type productUseCaseImpl struct {
	productRepository port.ProductRepository
	pgDb *postgresdb.PgDb
}



func NewProductUseCase(repository port.ProductRepository, pgDb *postgresdb.PgDb) port.ProductUseCase {
	return &productUseCaseImpl{
		productRepository: repository,
		pgDb: pgDb,
	}
}

func (p productUseCaseImpl) GetAll(ctx context.Context, param domain.FindProductFilter) ([]*domain.Product, error) {
	dbTx, err := p.pgDb.BeginTrx(ctx , nil)

	var opts []sqkit.SelectOption
	opts = append(opts, &sqkit.OffsetPagination{Offset: param.Offset, Limit: param.Limit})
	if param.Sort != "" {
		opts = append(opts, sqkit.Sorts(strings.Split(param.Sort, ",")))
	}

	products, err := p.productRepository.Find(ctx, dbTx, false, opts...)
	if err != nil {
		return nil, err
	} else if len(products) < 1 {
		return nil, echo.ErrNotFound
	}
	return products, nil
}

func (p productUseCaseImpl) FindById(ctx context.Context, paramId string) (*domain.Product, error) {
	id, _ := strconv.ParseInt(paramId, 10, 64)
	return p.findById(ctx, id)
}

func (p productUseCaseImpl)  findById(ctx context.Context, id int64)(*domain.Product, error){
	dbTx, err := p.pgDb.BeginTrx(ctx , nil)
	products, err := p.productRepository.Find(ctx, dbTx, false, sqkit.Eq{domain.ProductTable.ID : id})
	if err != nil {
		return nil, err
	} else if len(products) < 1 {
		return nil, echo.ErrNotFound
	}
	return products[0], nil
}

