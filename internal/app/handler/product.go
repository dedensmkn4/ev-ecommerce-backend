package handler

import (
	"github.com/dedensmkn4/ev-ecommerce-backend/pkg/echokit"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) FindProductById(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := h.productUseCase.FindById(ctx, c.Param("id"))

	if err != nil {
		return echokit.HTTPError(err)
	}

	return c.JSON(http.StatusOK, data)
}