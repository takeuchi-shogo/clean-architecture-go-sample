package product

import (
	"fmt"

	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/gateways/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases/interactor/product"
)

type AccountsController struct {
	Interactor product.AccountInteractor
}

func NewAccountsController() *AccountsController {
	return &AccountsController{
		Interactor: product.AccountInteractor{
			User: &repositories.UserRepository{},
		},
	}
}

func (c *AccountsController) Post(ctx controllers.Context) {

	screenName := ctx.PostForm("screenName")
	email := ctx.PostForm("email")

	fmt.Println(screenName, email)
}
