package entities

type Users struct {
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"`
	ScreenName  string `json:"screenName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}
