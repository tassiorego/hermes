package internalerrors

import (
	"errors"
	"net/http"
)

func StatusFor(err error) int {
	switch {
	case errors.Is(err, ValidationError),
		errors.Is(err, BadRequestError):
		return http.StatusBadRequest
	case errors.Is(err, UnauthorizedError):
		return http.StatusUnauthorized
	case errors.Is(err, ForbiddenError):
		return http.StatusForbidden
	case errors.Is(err, NotFoundError):
		return http.StatusNotFound
	case errors.Is(err, ConflictError):
		return http.StatusConflict
	case errors.Is(err, ServiceUnavailableError):
		return http.StatusServiceUnavailable
	case errors.Is(err, GatewayTimeoutError):
		return http.StatusGatewayTimeout
	default: // InternalServerError and any unknown error
		return http.StatusInternalServerError
	}
}
