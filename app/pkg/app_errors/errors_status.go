package api_errors

import "net/http"

func (e *ApiError) BadRequest() AppError {
	e.StatusCode = http.StatusBadRequest
	return e
}

func (e *ApiError) Unauthorized() AppError {
	e.StatusCode = http.StatusUnauthorized
	return e
}

func (e *ApiError) PaymentRequired() AppError {
	e.StatusCode = http.StatusPaymentRequired
	return e
}

func (e *ApiError) NotFound() AppError {
	e.StatusCode = http.StatusNotFound
	return e
}

func (e *ApiError) TooManyRequests() AppError {
	e.StatusCode = http.StatusTooManyRequests
	return e
}

func (e *ApiError) InternalServerError() AppError {
	e.StatusCode = http.StatusInternalServerError
	return e
}

func (e *ApiError) ServiceUnavailable() AppError {
	e.StatusCode = http.StatusServiceUnavailable
	return e
}
