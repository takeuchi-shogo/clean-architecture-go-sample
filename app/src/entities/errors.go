package entities

type ErrorResponse struct {
	RequestID  string         `json:"requestId"`
	StatusCode int            `json:"code"`
	ErrorType  string         `json:"type"` // Maybe unnecessary?
	Errors     []ErrorDetails `json:"errors"`
}

type ErrorDetails struct {
	Doamin   string `json:"domain"`   // Set the default to global.
	Resource string `json:"resource"` // where did it happen.
	Title    string `json:"title"`    // error reason.
	Message  string `json:"message"`  // error message.
	// ↓ Error content displayed on the frontend
	ErrorUserTitle   string `json:"errorUserTitle"`
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

type ErrorTitle struct {
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

func NewErrorTitle(lang string) *ErrorTitle {
	et := &ErrorTitle{}

	et.setBadRequest(lang)
	et.setInternalServerError()
	et.setServiceUnavailable()

	return et
}

// 400 error
func (et *ErrorTitle) setBadRequest(lang string) {
	switch lang {
	case "en":
		et.BadRequest.BadRequest = "bad request"
		et.BadRequest.Invalid = "invalid"
		et.BadRequest.InvalidParameter = "invald parameter"
		et.BadRequest.InvalidQuery = "invald query"
		et.BadRequest.NotDownload = "not download"
		et.BadRequest.NotUpload = "not upload"
		et.BadRequest.PaeseError = "parse error"
		et.BadRequest.Required = "required"
		et.BadRequest.UnknownApi = "unknown api"
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
func (et *ErrorTitle) setInternalServerError() {
	et.InternalServerError.internalError = "internal server error"
}

//503 error
func (et *ErrorTitle) setServiceUnavailable() {
	et.ServiceUnavailable.BackendError = ""
	et.ServiceUnavailable.BackendConnected = ""
	et.ServiceUnavailable.NotReady = ""
	et.ServiceUnavailable.Maintenance = "Maintenance mode"
}

func NewErrorResponse(code int, resources []string) *ErrorResponse {

	er := &ErrorResponse{}

	etp := NewErrorType()

	er = er.setError(code, resources, etp)

	return er
}

func (er *ErrorResponse) setError(code int, resources []string, etp *ErrorType) *ErrorResponse {

	switch code {
	case 200:
		return nil
	case 400:
		er.StatusCode = code
		er.ErrorType = etp.BadRequest
		er.Errors = setErrors(resources)
	case 401:
		er.StatusCode = code
		er.ErrorType = etp.Unauthorized
	case 403:
		er.StatusCode = code
		er.ErrorType = etp.Forbidden
	case 404:
		er.StatusCode = code
		er.ErrorType = etp.NotFound
	case 405:
		er.StatusCode = code
		er.ErrorType = etp.MethodNotAllowed
	case 406:
		er.StatusCode = code
		er.ErrorType = etp.Conflict
	case 429:
		er.StatusCode = code
		er.ErrorType = etp.TooManyRequests

	case 500:
		er.StatusCode = code
		er.ErrorType = etp.InternalServerError
	case 502:
		er.StatusCode = code
		er.ErrorType = etp.BadGateways
	case 503:
		er.StatusCode = code
		er.ErrorType = etp.ServiceUnavailable
	default:
		return nil
	}

	return er
}

func setErrors(resources []string) []ErrorDetails {
	eds := []ErrorDetails{}

	et := NewErrorTitle("ja")
	er := NewErrorResource()
	for _, r := range resources {
		ed := ErrorDetails{}

		ed.Doamin = "global"

		switch r {
		case er.User.Get:
			ed.Resource = r
			ed.Title = et.BadRequest.BadRequest
			ed.Message = "テスト"
			ed.ErrorUserTitle = et.BadRequest.BadRequest
			ed.ErrorUserMessage = getOneUserError
		}

		eds = append(eds, ed)
	}

	return eds
}

/*
error response example:When I get a user that doesn't exist.
{
	"error": {
		"requestId": "",
		"code": 400,
		"type": "bad request",
		"errors": [
			"domain": "global",
			"resource": "users.get",
			"title": "user not found",
			"message: "user does not exist",
			"errorUserTitle": "ユーザーを取得できませんでした",
			"errorUserMessage": "存在しないユーザーです",
		]
	}
}
*/

var (
	// user error list
	getOneUserError = "ユーザーの取得に失敗しました"
	// Account error list
	getAccountError    = "アカウントの取得に失敗しました"
	createAccountError = "アカウントの作成に失敗しました"
	updateAccountError = "アカウントの更新に失敗しました"
	deleteAccountError = "アカウントの削除に失敗しました"
)
