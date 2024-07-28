package entities

import "errors"

var (
	// user error list
	getOneUserErr = errors.New("ユーザーの取得に失敗しました")
	// Account error list
	getAccountErr    = errors.New("アカウントの取得に失敗しました")
	createAccountErr = errors.New("アカウントの作成に失敗しました")
	updateAccountErr = errors.New("アカウントの更新に失敗しました")
	deleteAccountErr = errors.New("アカウントの削除に失敗しました")
	// Validate error list
	validateErr            = errors.New("正しく入力されていない項目があります")
	validateEmptyErr       = errors.New("入力されていません")
	validateDisplayNameErr = errors.New("不適切なアカウント名です")
	validateScreenNameErr  = errors.New("")
	validateEmailErr       = errors.New("正しいメールアドレスの形式ではありません")
	validatePasswordErr    = errors.New("正しく入力されていません")
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
