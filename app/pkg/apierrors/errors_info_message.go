package apierrors

/* Error message for user */
var (
	// User error
	GetUserErr     = "ユーザーが見つかりません"
	GetUserListErr = "ユーザーの取得に失敗しました"

	// Account error
	GetAccountErr    = "アカウントの取得に失敗しました"
	CreateAccountErr = "アカウントの作成に失敗しました"
	UpdateAccountErr = "アカウントの更新に失敗しました"
	DeleteAccountErr = "アカウントの削除に失敗しました"

	// Login error
	LoginErr = "ログインに失敗しました"

	// Signup error
	SignupErr = "サインアップに作成しました"

	// invalid parameter error
	InvalidParameterErr = "パラメーターが無効です"

	// Empty error
	EmptyParameterErr = "入力されていません"
)
