package errors

import (
	"net/http"
)

type RestErr struct {
	Status    int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr{
	return &RestErr{
		Status  :  http.StatusBadRequest,
		Message : message,
		Error   : "bad request",
	}
}

func NewBadNotFoundError(message string) *RestErr{
	return &RestErr{
		Status  :  http.StatusNotFound,
		Message : message,
		Error   : "not found",
	}
}
