package product

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UsersPresenter struct {
}

func (p *UsersPresenter) ViewUser(c controllers.Context, user entities.Users) {
	c.JSON(200, nil)
}
