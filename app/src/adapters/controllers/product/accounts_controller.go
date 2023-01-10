package product

import (
	"strconv"

	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/gateways/middlewares"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/gateways/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases/interactor/product"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type AccountsController struct {
	Interactor product.AccountInteractor
}

type AccountsControllerProvider struct {
	DB  repositories.DB
	Jwt middlewares.JwtMiddleware
}

func NewAccountsController(p AccountsControllerProvider) *AccountsController {
	return &AccountsController{
		Interactor: product.AccountInteractor{
			DB:   &repositories.DBRepository{DB: p.DB},
			User: &repositories.UserRepository{},
		},
	}
}

func (c *AccountsController) Get(ctx controllers.Context) {
	userID, err := strconv.Atoi("100")

	if err != nil {
		code := 400
		ctx.JSON(code, entities.NewErrorResponse(code, []string{}, err))
		return
	}

	account, res := c.Interactor.Get(userID)

	if res.Error != nil {
		ctx.JSON(res.Code, entities.NewErrorResponse(res.Code, res.Resources, res.Error))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", account))
}

func (c *AccountsController) Post(ctx controllers.Context) {

	displayName := ctx.PostForm("displayName")
	screenName := ctx.PostForm("screenName")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	createdAccount, res := c.Interactor.Create(entities.Users{
		DisplayName: displayName,
		ScreenName:  screenName,
		Email:       email,
		Password:    password,
	})

	if res.Error != nil {
		ctx.JSON(res.Code, entities.NewErrorResponse(res.Code, res.Resources, res.Error))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", createdAccount))
}
