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
	default:
		return http.StatusInternalServerError
	}
}

func MessageFor(err error) string {
	if isUnknownError(err) {
		return "Internal server error"
	}
	return err.Error()
}

func isUnknownError(err error) bool {
	return !errors.Is(err, ValidationError) &&
		!errors.Is(err, BadRequestError) &&
		!errors.Is(err, UnauthorizedError) &&
		!errors.Is(err, ForbiddenError) &&
		!errors.Is(err, NotFoundError) &&
		!errors.Is(err, ConflictError) &&
		!errors.Is(err, ServiceUnavailableError) &&
		!errors.Is(err, GatewayTimeoutError) &&
		!errors.Is(err, InternalServerError)
}
