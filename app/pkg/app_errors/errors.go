package api_errors

type ApiError struct {
	// Return a typical error
	Next        error  // err
	Message     string // invalid string value: 'asdf'.
	InfoMessage string // 入力項目が無効です
	StatusCode  int    // 400
	Code        string // invalid_parameter

	// Errors []ErrorItems
	level string
}
