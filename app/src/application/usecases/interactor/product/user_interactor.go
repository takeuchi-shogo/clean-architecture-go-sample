package product

import (
	"errors"

	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases/output"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UserInteractor struct {
	DB            repositories.DBRepository
	User          repositories.UserRepository
	UserPresenter output.UserInteractorOutput
}

func (i *UserInteractor) Get(id string) (user *entities.Users, resultStatus *usecases.ResultStatus) {
	db := i.DB.Conn()

	user, err := i.User.FindByID(db, id)
	if err != nil {
		return &entities.Users{}, usecases.NewResultStatus(400, []string{"users.get"}, errors.New("e"))
	}
	return user, usecases.NewResultStatus(200, []string{}, nil)
}
