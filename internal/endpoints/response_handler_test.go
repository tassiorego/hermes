package endpoints

import (
	"errors"
	internalerrors "hermes/internal/internal-errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseHandler_WhenReturnUnknownError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return nil, errors.New("unknown error")
	}
	handlerFunc := ResponseHandler(200, endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc(res, req)
	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), "Internal server error")
}

func TestResponseHandler_WhenReturnInternalServerError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return nil, internalerrors.InternalServerError
	}
	handlerFunc := ResponseHandler(200, endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc(res, req)
	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), "Internal server error")
}

func TestResponseHandler_WhenReturnBadRequestError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return nil, internalerrors.BadRequestError
	}
	handlerFunc := ResponseHandler(200, endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc(res, req)
	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "Bad request")
}

func TestResponseHandler_WhenReturnUnauthorizedError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return nil, internalerrors.UnauthorizedError
	}
	handlerFunc := ResponseHandler(200, endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc(res, req)
	assert.Equal(http.StatusUnauthorized, res.Code)
	assert.Contains(res.Body.String(), "Unauthorized")
}

func TestResponseHandler_WhenReturnForbiddenError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return nil, internalerrors.ForbiddenError
	}
	handlerFunc := ResponseHandler(200, endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc(res, req)
	assert.Equal(http.StatusForbidden, res.Code)
	assert.Contains(res.Body.String(), "Forbidden")
}

func TestResponseHandler_WhenReturnNotFoundError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return nil, internalerrors.NotFoundError
	}
	handlerFunc := ResponseHandler(200, endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc(res, req)
	assert.Equal(http.StatusNotFound, res.Code)
	assert.Contains(res.Body.String(), "Not found")
}

func TestResponseHandler_WhenReturnConflictError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return nil, internalerrors.ConflictError
	}
	handlerFunc := ResponseHandler(200, endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc(res, req)
	assert.Equal(http.StatusConflict, res.Code)
	assert.Contains(res.Body.String(), "Conflict")
}

func TestResponseHandler_WhenReturnServiceUnavailableError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return nil, internalerrors.ServiceUnavailableError
	}
	handlerFunc := ResponseHandler(200, endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc(res, req)
	assert.Equal(http.StatusServiceUnavailable, res.Code)
	assert.Contains(res.Body.String(), "Service unavailable")
}

func TestResponseHandler_WhenReturnGatewayTimeoutError(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return nil, internalerrors.GatewayTimeoutError
	}
	handlerFunc := ResponseHandler(200, endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc(res, req)
	assert.Equal(http.StatusGatewayTimeout, res.Code)
	assert.Contains(res.Body.String(), "Gateway timeout")
}

func TestResponseHandler_WhenSuccessful(t *testing.T) {
	assert := assert.New(t)
	expectedResponse := map[string]string{"message": "success"}
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		return expectedResponse, nil
	}
	handlerFunc := ResponseHandler(200, endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc(res, req)
	assert.Equal(http.StatusOK, res.Code)
	assert.Contains(res.Body.String(), `"message":"success"`)
}
