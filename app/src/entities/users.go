package entities

import "errors"

type Users struct {
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"`
	ScreenName  string `json:"screenName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`

	Posts []Posts `json:"post" gorm:"forgenkey:UserID"`
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

func setErrorList(errList []error, message string) []error {
	errList = append(errList, errors.New(message))
	return errList
}
