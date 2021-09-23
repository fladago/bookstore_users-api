package errors

import "net/http"

type RestErr struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"code,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not found",
	}
}
