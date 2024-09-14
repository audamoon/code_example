package http_errors

import (
	"errors"
	"net/http"
)

//Error

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

var (
	ErrNotFound       = errors.New("not found")
	ErrNotEnoughMoney = errors.New("not enough money")
	ErrBadRequest     = errors.New("bad request")
)

func GetCodeAndResponse(err error) (int, *ErrorResponse) {
	var (
		code      int
		errorText string
	)

	switch {
	case errors.Is(err, ErrNotFound):
		code = http.StatusNotFound
		errorText = "not found"
	case errors.Is(err, ErrNotEnoughMoney):
		code = http.StatusInternalServerError
		errorText = "not enough money"
	default:
		code = http.StatusInternalServerError
		errorText = "internal error"
	}

	return code, &ErrorResponse{
		Status: code,
		Error:  errorText,
	}
}
