package product

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases/presenter"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UserInteractor struct {
	User          repositories.UserRepository
	UserPresenter presenter.UserPresenterOutput
}

func (i *UserInteractor) Get(id int) (user entities.Users, response *entities.Response) {

	user, err := i.User.FindByID(id)
	if err == nil {
		return entities.Users{}, entities.NewResponse(400, "notFound", "user")
	}
	return user, entities.NewResponse(200, "", "")
}
