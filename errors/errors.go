package errors

import "net/http"

type AppErrors struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e *AppErrors) AsMessage() *AppErrors {
	return &AppErrors{
		Message: e.Message,
	}
}

func (e *AppErrors) Error() string {
	return e.Message
}

func NotFoundError(message string) *AppErrors {
	return &AppErrors{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func InternalServerError(message string) *AppErrors {
	return &AppErrors{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
