package entities

var (
	// user error list
	getOneUserError = "ユーザーの取得に失敗しました"
	// Account error list
	getAccountError    = "アカウントの取得に失敗しました"
	createAccountError = "アカウントの作成に失敗しました"
	updateAccountError = "アカウントの更新に失敗しました"
	deleteAccountError = "アカウントの削除に失敗しました"
	// Validate error list
	validateError            = "正しく入力されていない項目があります"
	validateDisplayNameError = "不適切なアカウント名です"
)

/*
error response example:When I get a user that doesn't exist.
{
	"error": {
		"requestId": "",
		"code": 400,
		"type": "bad request",
		"errors": [
			{
				"domain": "global",
				"resource": "users.get",
				"reason": "bad Request",
				"title": "user not found",
				"message: "user does not exist",
				"errorUserTitle": "ユーザーを取得できませんでした",
				"errorUserMessage": "存在しないユーザーです",
			},
		],
	}
}
*/
