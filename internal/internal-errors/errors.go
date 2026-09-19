package internalerrors

import "errors"

var InternalServerError error = errors.New("Internal servererror")
var NotFoundError error = errors.New("Not found")
var ValidationError error = errors.New("Validation error")
var UnauthorizedError error = errors.New("Unauthorized")
var ForbiddenError error = errors.New("Forbidden")
var ConflictError error = errors.New("Conflict")
var BadRequestError error = errors.New("Bad request")
var ServiceUnavailableError error = errors.New("Service unavailable")
var GatewayTimeoutError error = errors.New("Gateway timeout")
