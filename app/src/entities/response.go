package entities

type Response struct {
	Code  int
	Error *ErrorResponse
}

func NewResponse(code int, t, resource string) *Response {
	r := &Response{}

	r.Code = code

	if t != "" && resource != "" {
		r.Error = &ErrorResponse{}

		r.setErrorDetails()
	}

	return r
}

/*
codeによってErrorTypeは決まる
resourceによってErrorMessageは決まる
*/

func (e *Response) setErrorDetails() {

}
