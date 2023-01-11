package apierrors

import (
	"errors"
)

type ApiError struct {
	// Return a typical error
	Next        error  `json:"next"`        // err
	Message     string `json:"message"`     // invalid string value: 'asdf'.
	InfoMessage string `json:"infoMessage"` // 入力項目が無効です
	StatusCode  int    `json:"statusCode"`  // 400
	Code        string `json:"code"`        // invalid_parameter

	// Errors []ErrorItems
	level string
}

func (e ApiError) New(msg string) ApiError {
	err := errors.New(msg)
	e.Next = err
	e.InfoMessage = msg

	return e
}
