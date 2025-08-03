package rest_err

import (
	"net/http"

)

// RestErr represents a standard error response returned by the API
type RestErr struct {
	// Human-readable error message
	// example: invalid user ID
	Message string `json:"message"`

	// Machine-readable error identifier
	// example: bad_request
	Err string `json:"error"`

	// HTTP status code
	// example: 400
	Code int `json:"code"`

}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int) *RestErr{
	return &RestErr{
		Message: message,
		Err: err,
		Code: code,
	}
}

func NewBadRequestError(message string) *RestErr{
	return NewRestErr(message, "bad_request", http.StatusBadRequest)
}

func NewBadRequestValidationError(message string) *RestErr{
	return NewRestErr(message, "bad_request", http.StatusBadRequest)
}

func NewInternalServerError(message string) *RestErr{
	return NewRestErr(message, "internal_server_error", http.StatusInternalServerError)
}

func NewUnauthorizedError(message string) *RestErr{
	return NewRestErr(message, "unauthorized", http.StatusUnauthorized )
}

func NewForbiddenError(message string) *RestErr{
	return NewRestErr(message, "forbidden", http.StatusForbidden)
}

func NewConflictError(message string) *RestErr{
	return NewRestErr(message, "conflict", http.StatusConflict)
}

func NewNotFoundError(message string) *RestErr{
	return NewRestErr(message, "not_found", http.StatusNotFound)
}