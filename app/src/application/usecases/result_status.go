package usecases

// Controllerへ情報をまとめて返すためのオブジェクト
type ResultStatus struct {
	Code      int
	Resources []string
	Error     error
}

func NewResultStatus(code int, resources []string, err error) *ResultStatus {
	return &ResultStatus{
		Code:      code,
		Resources: resources,
		Error:     err,
	}
}
