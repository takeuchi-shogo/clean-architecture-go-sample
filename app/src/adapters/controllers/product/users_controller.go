package product

import (
	"strconv"

	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases/product"
)

type UsersController struct {
	Interactor product.UserInteractor
}

func NewUsersController() *UsersController {
	return &UsersController{
		Interactor: product.UserInteractor{},
	}
}

func (c *UsersController) Get(ctx controllers.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user, res := c.Interactor.Get(id)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.ErrorType, nil))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", user))
}
