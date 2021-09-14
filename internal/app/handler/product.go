package handler

import (
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/pkg/echokit"
	"github.com/dedensmkn4/ev-ecommerce-backend/pkg/httpkit"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) GetAllProduct(c echo.Context) error {
	ctx := c.Request().Context()
	req := domain.FindProductFilter{}

	bodyByte := httpkit.DumpHTTPRequest(c.Request())

	// biding payload to struct
	if err := c.Bind(&req); err != nil {
		logrus.WithContext(c.Request().Context()).WithFields(logrus.Fields{
			"request": string(bodyByte),
		}).Error(errors.Wrap(err, "Error binding request"))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	data, err := h.productUseCase.GetAll(ctx, req)

	if err != nil {
		return echokit.HTTPError(err)
	}

	return c.JSON(http.StatusOK, data)
}

func (h *Handler) FindProductById(c echo.Context) error {
	ctx := c.Request().Context()

	data, err := h.productUseCase.FindById(ctx, c.Param("id"))

	if err != nil {
		return echokit.HTTPError(err)
	}

	return c.JSON(http.StatusOK, data)
}