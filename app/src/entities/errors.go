package entities

type ErrorResponse struct {
	RequestID  string         `json:"requestId"`
	StatusCode int            `json:"code"`
	ErrorType  string         `json:"type"`
	Errors     []ErrorDetails `json:"errors"`
}

type ErrorDetails struct {
	Resource      string `json:"resource"`
	ErrorTitle    string `json:"title"`
	ErrorMessages string `json:"messages"`
}

type ErrorMessage struct {
	UserErr struct {
		NotFound string
	}
}

func NewErrorMessage() *ErrorMessage {

	em := &ErrorMessage{}

	em.UserErr.NotFound = "user account is not found"

	return em
}

func NewErrorResponse(code int) *ErrorResponse {

	er := &ErrorResponse{}

	er = er.setError(code)

	return er
}

func (er *ErrorResponse) setError(code int) *ErrorResponse {

	switch code {
	case 200:
		return nil
	case 400:

	case 401:
	case 403:
	case 404:
	case 405:
	case 408:
	case 409:
	case 422:
	case 429:

	case 500:
	case 502:
	case 503:
	default:
		return nil
	}

	return nil
}

/*
error response example:When I get a user that doesn't exist.
{
	"errorResponse": {
		"requestId": "",
		"cose": 400,
		"type": "not found",
		"errors": [
			"resource": "users.get",
			"title": "user not found",
			"message: "user not exists",
		]
	}
}
status code list
###############################
200 OK
201 Created
202 Accepted

400 Bad Request
401 Unauthorized
403 Forbidden
404 Not Found
405 Method Not Allowed
408 Request Timeout
409 Conflict
422 Unprocessable Entity
429 Too Many Requests

500 Internal Server Error
502 Bad Gateways
503 Service Unavailable
###############################
*/
var ()
