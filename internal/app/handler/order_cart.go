package handler

import (
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/pkg/echokit"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) AddToCart(c echo.Context) error {
	ctx := c.Request().Context()

	var req domain.AddToCartPayload
	c.Bind(&req)
	data, err := h.orderUseCase.AddCart(ctx, req)

	if err != nil {
		return echokit.HTTPError(err)
	}

	return c.JSON(http.StatusOK, data)
}