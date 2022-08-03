package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(msg string) *RestErr {
	return &RestErr{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
