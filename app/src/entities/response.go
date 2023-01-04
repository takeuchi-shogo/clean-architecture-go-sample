package entities

type Response struct {
	Code  int
	Error *ErrorResponse
}

func NewResponse(code int, t, resource string) *Response {
	r := &Response{
		Code: code,
		Error: &ErrorResponse{
			RequestID: "",
			ErrorType: t,
			Errors: []ErrorDetails{
				{
					Resource:      resource,
					ErrorTitle:    "",
					ErrorMessages: "",
				},
			},
		},
	}

	r.setErrorDetails()

	return r
}

/*
codeによってErrorTypeは決まる
resourceによってErrorMessageは決まる
*/

func (e *Response) setErrorDetails() {

}
