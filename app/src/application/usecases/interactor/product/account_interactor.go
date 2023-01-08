package product

import (
	"errors"

	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/utilities"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type AccountInteractor struct {
	User repositories.UserRepository
}

func (i *AccountInteractor) Get(userID int) (account entities.Users, resultStatus *usecases.ResultStatus) {
	foundUser, err := i.User.FindByID(userID)
	if err != nil {
		return entities.Users{}, usecases.NewResultStatus(400, []string{"account.get"}, errors.New("test get account error"))
	}
	return foundUser, usecases.NewResultStatus(200, []string{}, nil)
}

func (i *AccountInteractor) Create(user entities.Users) (createdUser entities.Users, resultStatus *usecases.ResultStatus) {
	newPassword, err := utilities.HashPassword(user.Password)
	if err != nil {
		return entities.Users{}, usecases.NewResultStatus(400, []string{"accounts.create"}, err)
	}

	user.Password = newPassword
	user.CreatedAt = utilities.SetCurrentUnixTime()
	user.UpdatedAt = utilities.SetCurrentUnixTime()

	createdUser, err = i.User.Create(user)
	if err != nil {
		return entities.Users{}, usecases.NewResultStatus(400, []string{"accounts.create"}, errors.New("test create error"))
	}
	return createdUser, usecases.NewResultStatus(200, []string{}, nil)
}
