package product

import (
	"fmt"

	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/gateways/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases/interactor/product"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type userAuthController struct {
	Interactor product.UserAuthInteractor
}

type UserAuthControllerProvider struct {
	DB repositories.DB
}

func NewUserAuthController(p UserAuthControllerProvider) *userAuthController {
	return &userAuthController{
		Interactor: product.UserAuthInteractor{
			DB:   &repositories.DBRepository{DB: p.DB},
			User: &repositories.UserRepository{},
		},
	}
}

func (c *userAuthController) Post(ctx controllers.Context) {
	screenName := ctx.PostForm("screenName")
	password := ctx.PostForm("password")

	// errList := checkParam(screenName, password)
	// if len(errList) > 0 {

	// }

	auth, res := c.Interactor.Create(entities.Users{
		ScreenName: screenName,
		Password:   password,
	})
	fmt.Println("www")

	if res.Error != nil {
		ctx.JSON(res.Code, entities.NewErrorResponse(res.Code, res.Resources, res.Error))
		return
	}

	ctx.Header("Authorization", auth.JwtToken)
	ctx.JSON(res.Code, controllers.NewH("success", auth.User))
}

func checkParam(sceenName, password string) []string {
	errList := []string{}
	if sceenName != "" {
		errList = append(errList, "auth.screenName")
	}
	if len(sceenName) >= 30 {
		errList = append(errList, "auth.maxLength")
	}

	if password != "" {
		errList = append(errList, "auth.password")
	}
	if len(password) >= 30 {
		errList = append(errList, "auth.maxLength")
	}
	return errList
}
