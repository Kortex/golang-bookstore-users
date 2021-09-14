package errors

import "net/http"

type RestError struct {
	Message      string `json:"message"`
	Status       int    `json:"status"`
	ErrorMessage string `json:"errorMessage"`
}

func BadRequest(message string) *RestError {
	return &RestError{
		Message: message,
		Status: http.StatusBadRequest,
		ErrorMessage: "bad_request",
	}
}

