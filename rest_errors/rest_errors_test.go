package rest_errors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInternalServerError(t *testing.T) {

	err := InternalServerError("Internal server error test")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Internal server error test", err.Message)
}

func TestBadRequestError(t *testing.T) {

	err := BadRequestError("Bad request error test")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "Bad request error test", err.Message)
}

func TestNotFoundError(t *testing.T) {

	err := NotFoundError("Not found error test")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Not found error test", err.Message)
}

func TestInvalidParameterError(t *testing.T) {

	err := InvalidParameterError("Invalid parameter error test")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusPreconditionFailed, err.Status)
	assert.EqualValues(t, "Invalid parameter error test", err.Message)
}
