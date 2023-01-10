package entities

type ErrorResponse struct {
	RequestID string         `json:"requestId"` // request id
	Code      int            `json:"code"`      // status code
	ErrorType string         `json:"type"`      // maybe unnecessary?
	Message   string         `json:"message"`   // error message
	Errors    []ErrorDetails `json:"errors"`
}

type ErrorDetails struct {
	Doamin   string `json:"domain"`   // Set the default to global.
	Resource string `json:"resource"` // where did it happen.
	Reason   string `json:"reason"`   // error reason.
	Message  string `json:"message"`  // error message.
	// ↓ Error content displayed on the frontend
	// ErrorUserTitle   string `json:"errorUserTitle"`
	ErrorUserMessage string `json:"errorUserMessage"`
}

type ErrorType struct {
	BadRequest       string
	Unauthorized     string
	Forbidden        string
	NotFound         string
	MethodNotAllowed string
	Conflict         string
	TooManyRequests  string

	InternalServerError string
	BadGateways         string
	ServiceUnavailable  string
}

func NewErrorType() *ErrorType {
	return &ErrorType{
		BadRequest:       "Bad Request",
		Unauthorized:     "Unauthorized",
		Forbidden:        "Forbidden",
		NotFound:         "Not Found",
		MethodNotAllowed: "Method Not Allowed",
		Conflict:         "Conflict",
		TooManyRequests:  "Too Many Requests",

		InternalServerError: "Internal Server Error",
		BadGateways:         "Bad Gateways",
		ServiceUnavailable:  "Service Unavailable",
	}
}

type ErrorResource struct {
	User struct {
		Get     string
		GetList string
		Create  string
		Update  string
		Delete  string
	}
	Account struct {
		Get    string
		Create string
		Update string
		Delete string
	}
}

func NewErrorResource() *ErrorResource {

	er := &ErrorResource{}

	er.setErrorResource()

	return er
}

func (er *ErrorResource) setErrorResource() {
	er.User.Get = "users.get"
	er.User.GetList = "users.getList"
	er.User.Create = "users.create"
	er.User.Update = "users.save"
	er.User.Delete = "users.delete"

	er.Account.Get = "account.get"
	er.Account.Create = "account.create"
	er.Account.Update = "account.save"
	er.Account.Delete = "account.delete"
}

type ErrorReason struct {
	// 400
	BadRequest struct {
		BadRequest       string
		Invalid          string
		InvalidParameter string
		InvalidQuery     string
		NotDownload      string
		NotUpload        string
		PaeseError       string
		Required         string
		UnknownApi       string
	}
	// 401
	Unauthorized struct {
		Unauthorized        string
		authError           string
		Expired             string
		LockedDomainExpired string
		Required            string
	}
	// 402
	PaymentRequired struct {
		DailyLimitExceeded402 string
		QuotaExceeded402      string
		user402               string
	}
	// 403
	Forbidden struct {
		Forbidden           string
		AccessNotConfigured string
		AccountDeleted      string
		AccountDisabled     string
		AccountUnverfied    string
	}
	// 404
	NotFound struct {
		NotFound string
	}
	// 405
	MethodNotAllowed struct {
		HttpMethodNotAllowed string
	}
	// 406
	Conflict struct {
		Conflict  string
		Duplicate string
	}
	// 409
	Gone struct {
		deleted string
	}
	// 429
	TooManyRequests struct {
		RateLimitExceeded string
	}
	// 500
	InternalServerError struct {
		internalError string
	}
	// 503
	ServiceUnavailable struct {
		BackendError     string
		BackendConnected string
		NotReady         string
		Maintenance      string
	}
}

func NewErrorReason(lang string) *ErrorReason {
	et := &ErrorReason{}

	et.setBadRequest(lang)
	et.setInternalServerError()
	et.setServiceUnavailable()

	return et
}

// 400 error
func (et *ErrorReason) setBadRequest(lang string) {
	switch lang {
	case "en":
		et.BadRequest.BadRequest = "badRequest"
		et.BadRequest.Invalid = "invalid"
		et.BadRequest.InvalidParameter = "invaldParameter"
		et.BadRequest.InvalidQuery = "invaldQuery"
		et.BadRequest.NotDownload = "notDownload"
		et.BadRequest.NotUpload = "notUpload"
		et.BadRequest.PaeseError = "parseError"
		et.BadRequest.Required = "required"
		et.BadRequest.UnknownApi = "unknownApi"
	case "ja":
		et.BadRequest.BadRequest = "リクエストを正常に処理できませんでした"
		et.BadRequest.Invalid = "このリクエストは無効です"
		et.BadRequest.InvalidParameter = "無効なパラメータです"
		et.BadRequest.InvalidQuery = "無効なクエリです"
		et.BadRequest.NotDownload = "ダウンロードに失敗しました"
		et.BadRequest.NotUpload = "アップロードに失敗しました"
		et.BadRequest.PaeseError = "サーバーがリクエスト本文を解析できません"
		et.BadRequest.Required = "リクエストに必要な情報がありません。 必要な情報は、パラメーターまたはリソース プロパティです"
		et.BadRequest.UnknownApi = "リクエストが呼び出している API が認識されていません"
	}
}

// 500 error
func (et *ErrorReason) setInternalServerError() {
	et.InternalServerError.internalError = "internalServerError"
}

//503 error
func (et *ErrorReason) setServiceUnavailable() {
	et.ServiceUnavailable.BackendError = ""
	et.ServiceUnavailable.BackendConnected = ""
	et.ServiceUnavailable.NotReady = ""
	et.ServiceUnavailable.Maintenance = "maintenanceMode"
}

func NewErrorResponse(code int, resources []string, err error) *ErrorResponse {

	er := &ErrorResponse{}

	er = er.setError(code, resources, err)

	return er
}

func (er *ErrorResponse) setError(code int, resources []string, err error) *ErrorResponse {

	etp := NewErrorType()

	switch code {
	case 200:
		return nil
	case 400:
		er.Code = code
		er.ErrorType = etp.BadRequest
		er.Message = err.Error()
		er.Errors = setErrors(resources)
	case 401:
		er.Code = code
		er.ErrorType = etp.Unauthorized
		er.Message = err.Error()
	case 403:
		er.Code = code
		er.ErrorType = etp.Forbidden
		er.Message = err.Error()
	case 404:
		er.Code = code
		er.ErrorType = etp.NotFound
		er.Message = err.Error()
	case 405:
		er.Code = code
		er.ErrorType = etp.MethodNotAllowed
		er.Message = err.Error()
	case 406:
		er.Code = code
		er.ErrorType = etp.Conflict
		er.Message = err.Error()
	case 429:
		er.Code = code
		er.ErrorType = etp.TooManyRequests
		er.Message = err.Error()
	case 500:
		er.Code = code
		er.ErrorType = etp.InternalServerError
		er.Message = err.Error()
	case 502:
		er.Code = code
		er.ErrorType = etp.BadGateways
		er.Message = err.Error()
	case 503:
		er.Code = code
		er.ErrorType = etp.ServiceUnavailable
		er.Message = err.Error()
	default:
		return nil
	}

	return er
}

func setErrors(resources []string) []ErrorDetails {
	eds := []ErrorDetails{}

	et := NewErrorReason("ja")
	er := NewErrorResource()
	for _, resource := range resources {
		ed := ErrorDetails{}

		ed.Doamin = "global"

		switch resource {
		case er.User.Get:
			ed.Resource = resource
			ed.Reason = et.BadRequest.BadRequest
			ed.Message = "テスト"
			// ed.ErrorUserTitle = et.BadRequest.BadRequest
			ed.ErrorUserMessage = getOneUserErr.Error()
		}

		eds = append(eds, ed)
	}

	return eds
}
