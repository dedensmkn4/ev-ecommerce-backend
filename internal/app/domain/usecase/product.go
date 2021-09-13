package usecase

import (
	"context"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/port"
	"github.com/dedensmkn4/ev-ecommerce-backend/pkg/sqkit"
	"github.com/labstack/echo/v4"
	"strconv"
)

type productUseCaseImpl struct {
	productRepository port.ProductRepository
}

func NewProductUseCase(repository port.ProductRepository) port.ProductUseCase {
	return &productUseCaseImpl{repository}
}

func (p productUseCaseImpl) FindById(ctx context.Context, paramId string) (*domain.Product, error) {
	id, _ := strconv.ParseInt(paramId, 10, 64)
	return p.findById(ctx, id)
}

func (p productUseCaseImpl)  findById(ctx context.Context, id int64)(*domain.Product, error){
	products, err := p.productRepository.Find(ctx, sqkit.Eq{domain.ProductTable.ID : id})
	if err != nil {
		return nil, err
	} else if len(products) < 1 {
		return nil, echo.ErrNotFound
	}
	return products[0], nil
}