package echokit

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestHTTPError(t *testing.T) {
	testcases := []struct {
		TestName string
		err      error
		expected *echo.HTTPError
	}{
		{
			err:      errors.New("some-error"),
			expected: echo.NewHTTPError(http.StatusInternalServerError, "some-error"),
		},
		{
			err:      echo.NewHTTPError(99, "some-message"),
			expected: echo.NewHTTPError(99, "some-message"),
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.expected, HTTPError(tt.err))
		})
	}
}

func TestNewValidErr(t *testing.T) {
	require.Equal(t, echo.NewHTTPError(422, "some-message"), NewValidErr("some-message"))
}
