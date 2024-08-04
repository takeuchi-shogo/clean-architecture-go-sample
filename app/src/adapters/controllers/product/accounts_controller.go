package product

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/pkg/apierrors"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/gateways/middlewares"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/gateways/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases/interactor/product"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type AccountsController struct {
	Auth       product.UserAuthInteractor
	Interactor product.AccountInteractor
}

type AccountsControllerProvider struct {
	DB  repositories.DB
	Jwt middlewares.Jwt
}

func NewAccountsController(p AccountsControllerProvider) *AccountsController {
	return &AccountsController{
		Auth: product.UserAuthInteractor{
			DB:   &repositories.DBRepository{DB: p.DB},
			Jwt:  &middlewares.JwtMiddleware{Jwt: p.Jwt},
			User: &repositories.UserRepository{},
		},
		Interactor: product.AccountInteractor{
			DB:   &repositories.DBRepository{DB: p.DB},
			User: &repositories.UserRepository{},
		},
	}
}

func (c *AccountsController) Get(ctx controllers.Context) {
	token, res := c.Auth.Autholization(ctx.GetHeader("Authorization"))

	if res.Error != nil {
		ctx.JSON(res.Code, apierrors.BadRequest.New(res.Error.Error()))
		return
	}

	account, res := c.Interactor.Get(token.ID)

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

	createdAccount, res := c.Interactor.Create(&entities.Users{
		DisplayName: displayName,
		ScreenName:  screenName,
		Email:       email,
		Password:    password,
	})
	if res.Error != nil {
		ctx.JSON(res.Code, apierrors.BadRequest.New(res.Error.Error()))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", createdAccount))
}

func (c *AccountsController) Patch(ctx controllers.Context) {
	user := &entities.Users{}
	if err := ctx.BindJSON(user); err != nil {
		ctx.JSON(400, apierrors.BadRequest.New(err.Error()))
		return
	}

	id := ctx.Param("id")

	user.ID = id

	updatedAccount, res := c.Interactor.Save(user)
	if res.Error != nil {
		ctx.JSON(res.Code, apierrors.BadRequest.New(res.Error.Error()))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", updatedAccount))
}
