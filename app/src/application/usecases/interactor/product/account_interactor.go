package product

import (
	"errors"

	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type AccountInteractor struct {
	User repositories.UserRepository
}

func (i *AccountInteractor) Create(user entities.Users) (createdUser entities.Users, resultStatus *usecases.ResultStatus) {
	createdUser, err := i.User.Create(user)
	if err != nil {
		return entities.Users{}, usecases.NewResultStatus(400, []string{"users.create"}, errors.New("test create error"))
	}
	return createdUser, usecases.NewResultStatus(200, []string{}, nil)
}
