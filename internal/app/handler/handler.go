package handler

import (
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/config"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/port"
	"github.com/dedensmkn4/ev-ecommerce-backend/pkg/validation"
)

type (
	Handler struct {
		config    *config.Config
		validate  *validation.Validation
		productUseCase port.ProductUseCase
		orderCartUseCase port.OrderCartUseCase
	}

	HandlerConfig struct {
		Config      *config.Config
		Validator   *validation.Validation
		ProductUseCase port.ProductUseCase
		OrderCartUseCase port.OrderCartUseCase
	}
)

func NewHandler(hc HandlerConfig) *Handler {
	return &Handler{
		config:    hc.Config,
		validate:  hc.Validator,
		productUseCase : hc.ProductUseCase,
		orderCartUseCase: hc.OrderCartUseCase,
	}
}