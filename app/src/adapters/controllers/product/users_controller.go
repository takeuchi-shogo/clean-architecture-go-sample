package product

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/gateways/middlewares"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/gateways/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases/interactor/product"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UsersController struct {
	Autholization product.UserAuthInteractor
	Interactor    product.UserInteractor
}

type UsersControllerProvider struct {
	DB  repositories.DB
	Jwt middlewares.Jwt
}

func NewUsersController(p UsersControllerProvider) *UsersController {
	return &UsersController{
		Autholization: product.UserAuthInteractor{
			DB:   &repositories.DBRepository{DB: p.DB},
			Jwt:  &middlewares.JwtMiddleware{Jwt: p.Jwt},
			User: &repositories.UserRepository{},
		},
		Interactor: product.UserInteractor{
			DB:   &repositories.DBRepository{DB: p.DB},
			User: &repositories.UserRepository{},
		},
	}
}

func (c *UsersController) Get(ctx controllers.Context) {
	id := ctx.Param("id")

	user, res := c.Interactor.Get(id)
	if res.Error != nil {
		ctx.JSON(res.Code, entities.NewErrorResponse(res.Code, res.Resources, res.Error))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", user))
}
