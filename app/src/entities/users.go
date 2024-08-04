package entities

import "errors"

type Users struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	ScreenName  string `json:"screenName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
	DeletedAt   *int64 `json:"deletedAt"`

	Posts []*Posts `json:"posts" gorm:"foreignKey:UserID;references:ID"`
}

type UsersResponse struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	ScreenName  string `json:"screenName"`
	Email       string `json:"email"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`

	Posts []*Posts `json:"posts"`
}

func (e *Users) Validate() error {
	if err := e.validateDisplayName(); err != nil {
		return err
	}
	if err := e.validateScreenName(); err != nil {
		return err
	}
	if err := e.validatePassword(); err != nil {
		return err
	}
	return nil
}

func (e *Users) validateDisplayName() error {
	if e.DisplayName == "" {
		return errors.New("displayName required")
	}
	return nil
}

func (e *Users) validateScreenName() error {
	if e.ScreenName == "" {
		return errors.New("screenName required")
	}
	return nil
}

func (e *Users) validatePassword() error {
	if e.Password == "" {
		return errors.New("password required")
	}
	return nil
}

// func setErrorList(errList []error, message string) []error {
// 	errList = append(errList, errors.New(message))
// 	return errList
// }

func (e *Users) BuildResponse() *UsersResponse {
	return &UsersResponse{
		ID:          e.ID,
		DisplayName: e.DisplayName,
		ScreenName:  e.ScreenName,
		Email:       e.Email,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
