package echokit

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// NewValidErr create ValidationError
func NewValidErr(message string) *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnprocessableEntity, message)
}

// HTTPError convert error to *echo.HTTPError
func HTTPError(err error) *echo.HTTPError {
	if httpErr, ok := err.(*echo.HTTPError); ok {
		return httpErr
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
