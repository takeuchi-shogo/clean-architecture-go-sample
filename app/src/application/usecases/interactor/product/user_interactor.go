package product

import (
	"errors"

	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases/presenter"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UserInteractor struct {
	User          repositories.UserRepository
	UserPresenter presenter.UserPresenterOutput
}

func (i *UserInteractor) Get(id int) (user entities.Users, resultStatus *usecases.ResultStatus) {
	user, err := i.User.FindByID(id)
	if err != nil {
		return entities.Users{}, usecases.NewResultStatus(400, []string{"users.get"}, errors.New("e"))
	}
	return user, usecases.NewResultStatus(200, []string{}, nil)
}
