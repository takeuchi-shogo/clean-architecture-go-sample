package api_errors

type AppError interface {
	BadRequest() AppError
	Unauthorized() AppError
	PaymentRequired() AppError
	NotFound() AppError
	TooManyRequests() AppError
	InternalServerError() AppError
	ServiceUnavailable() AppError
}
