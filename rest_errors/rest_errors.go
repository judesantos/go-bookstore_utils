package rest_errors

import (
	"net/http"
)

type RestErrorInterfaceInterface interface {
	Error() string
	Message() string
	Status() int
	Causes() []interface{}
}

type restError struct {
	message string
	status  int
	error   string
	causes  []interface{}
}

func (e *restError) Error() string {
	return e.error
}

func (e *restError) Message() string {
	return e.message
}

func (e *restError) Status() int {
	return e.status
}

func (e *restError) Causes() []interface{} {
	return e.causes
}

//
// CustomError
//
func RestError(
	mesage string,
	status int,
	err string,
	causes []interface{},
) RestErrorInterface {
	return restError{
		message: message,
		status:  status,
		error:   err,
		causes:  causes,
	}
}

//
// UnauthorizedError
//
func UnauthorizedError(message string) RestErrorInterface {
	return restError{
		message: message,
		status:  http.StatusUnauthorized,
		error:   "Unauthorized",
	}
}

//
// NotFoundError
//
func NotFoundError(message string) RestErrorInterface {
	return restError{
		message: message,
		status:  http.StatusNotFound,
		error:   "Not Found",
	}
}

//
// BadRequestError
//
func BadRequestError(message string) RestErrorInterface {
	return restError{
		message: message,
		status:  http.StatusBadRequest,
		error:   "Bad Request",
	}
}

//
// InvalidParameterError
//
func InvalidParameterError(message string) RestErrorInterface {
	return restError{
		message: message,
		status:  http.StatusPreconditionFailed,
		error:   "Invalid Parameter",
	}
}

//
// InternalServerError
//
func InternalServerError(message string, err error) RestErrorInterface {
	res = restError{
		message: message,
		status:  http.StatusInternalServerError,
		error:   "Internal Server Error",
	}

	if err != nil {
		res.causes = append(res.causes, err.Error())
	}

	return res
}
