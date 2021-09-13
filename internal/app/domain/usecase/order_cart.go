package usecase

import (
	"context"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/port"
)

type orderCartUseCaseImpl struct {
	orderCartRepository port.OrderCartRepository
	orderCartDetailRepository port.OrderCartDetailRepository
	userRepository port.UserRepository
	productRepository port.ProductRepository
}

func NewOrderCartUseCase(orderRepository port.OrderCartRepository,
	orderCartDetailRepository port.OrderCartDetailRepository,
	userRepository port.UserRepository,
	productRepository port.ProductRepository)  port.OrderCartUseCase{
	return &orderCartUseCaseImpl{
		orderCartRepository: orderRepository,
		orderCartDetailRepository: orderCartDetailRepository,
		userRepository: userRepository,
		productRepository: productRepository,
	}
}

func (o orderCartUseCaseImpl) AddCart(ctx context.Context, param domain.AddToCartPayload) (cartResponse domain.AddToCartResponse, err error) {
	panic("implement me")
}

func (o orderCartUseCaseImpl) Checkout(ctx context.Context, param domain.CheckoutPayload) (checkoutResponse domain.CheckoutResponse, err error) {
	panic("implement me")
}


