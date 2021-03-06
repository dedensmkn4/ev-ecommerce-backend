package response_test

import (
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/response"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccess(t *testing.T) {
	act := response.Success("message", nil)
	ref := act.(response.SuccessBody)

	assert.Equal(t, false, ref.Error)
	assert.Equal(t, "message", ref.Message)
	assert.Equal(t, nil, ref.Data)
}

func TestSuccessForTest(t *testing.T) {
	body := `{"error":false,"message":"message","data":null}`
	act, err := response.SuccessForTest(body)

	assert.NoError(t, err)
	assert.Equal(t, false, act.Error)
	assert.Equal(t, "message", act.Message)
	assert.Equal(t, nil, act.Data)

	body = "error"
	act, err = response.SuccessForTest(body)
	assert.Error(t, err)
	assert.Empty(t, act)
}