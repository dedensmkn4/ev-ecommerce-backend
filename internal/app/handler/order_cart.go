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

func (h *Handler) AddToCart(c echo.Context) error {
	ctx := c.Request().Context()
	req := domain.AddCartItemPayload{}

	bodyByte := httpkit.DumpHTTPRequest(c.Request())

	// biding payload to struct
	if err := c.Bind(&req); err != nil {
		logrus.WithContext(c.Request().Context()).WithFields(logrus.Fields{
			"request": string(bodyByte),
		}).Error(errors.Wrap(err, "Error binding request"))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate payload struct
	if err := c.Validate(&req); err != nil {
		logrus.Error(errors.Wrap(err, "error Field"))
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	data, err := h.orderCartUseCase.AddCart(ctx, req)

	if err != nil {
		return echokit.HTTPError(err)
	}

	return c.JSON(http.StatusOK, data)
}

func (h *Handler) OrderCheckout(c echo.Context) error {
	ctx := c.Request().Context()
	req := domain.OrderCheckoutPayload{}

	bodyByte := httpkit.DumpHTTPRequest(c.Request())

	// biding payload to struct
	if err := c.Bind(&req); err != nil {
		logrus.WithContext(c.Request().Context()).WithFields(logrus.Fields{
			"request": string(bodyByte),
		}).Error(errors.Wrap(err, "Error binding request"))
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate payload struct
	if err := c.Validate(&req); err != nil {
		logrus.Error(errors.Wrap(err, "error Field"))
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	data, err := h.orderCartUseCase.Checkout(ctx, req)

	if err != nil {
		return echokit.HTTPError(err)
	}

	return c.JSON(http.StatusOK, data)
}