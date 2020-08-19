package rest_errors

import "net/http"

type RestError struct {
	Message string
	Status  int
}

func UnauthorizedError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusUnauthorized,
	}
}

func NotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
	}
}

func BadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
	}
}

func InternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
	}
}

func InvalidParameterError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusPreconditionFailed,
	}
}
