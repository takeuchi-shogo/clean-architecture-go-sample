package product

import (
	"fmt"
	"strconv"

	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/gateways/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases/interactor/product"
)

type UsersController struct {
	Interactor product.UserInteractor
}

func NewUsersController() *UsersController {
	return &UsersController{
		Interactor: product.UserInteractor{
			User: &repositories.UserRepository{},
		},
	}
}

func (c *UsersController) Get(ctx controllers.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user, res := c.Interactor.Get(id)
	fmt.Println(res)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.ErrorType, nil))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", user))
}
