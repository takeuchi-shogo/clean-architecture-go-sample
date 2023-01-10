package product

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases/interactor/product"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type userAuthController struct {
	Interactor product.UserAuthInteractor
}

func NewUserAuthController() *userAuthController {
	return &userAuthController{
		Interactor: product.UserAuthInteractor{},
	}
}

func (c *userAuthController) Post(ctx controllers.Context) {
	screenName := ctx.PostForm("screenName")
	password := ctx.PostForm("password")

	auth, res := c.Interactor.Create(entities.Users{
		ScreenName: screenName,
		Password:   password,
	})

	if res.Error != nil {
		ctx.JSON(res.Code, entities.NewErrorResponse(res.Code, res.Resources, res.Error))
		return
	}

	ctx.Header("Authorization", auth.JwtToken)
	ctx.JSON(res.Code, controllers.NewH("success", auth.User))
}
