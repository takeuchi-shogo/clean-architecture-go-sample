package product

import "github.com/takeuchi-shogo/clean-architecture-golang/src/entities"

type UserInteractor struct {
}

func (i *UserInteractor) Get(id int) (user entities.Users, response *entities.Response) {
	return user, entities.NewResponse(200, "", "")
}
