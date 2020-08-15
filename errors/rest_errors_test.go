package errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadRequestError(t *testing.T) {
	message := "some message"

	err := NewBadRequestError(message)

	assert.NotNil(t, err)
	assert.EqualValues(t, message, err.Message)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	message := "some message"

	err := NewNotFoundError(message)

	assert.NotNil(t, err)
	assert.EqualValues(t, message, err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "not_found", err.Error)
}

func TestNewInternatServerError(t *testing.T) {
	message := "some message"

	err := NewInternatServerError(message)

	assert.NotNil(t, err)
	assert.EqualValues(t, message, err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "internal_server_error", err.Error)
}
