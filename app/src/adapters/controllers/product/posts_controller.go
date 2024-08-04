package product

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/pkg/apierrors"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/controllers"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/adapters/gateways/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases/interactor/product"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type PostsController struct {
	Interactor product.PostInteractor
}

type PostsControllerProvider struct {
	DB repositories.DB
}

func NewPostsController(p PostsControllerProvider) *PostsController {
	return &PostsController{
		Interactor: product.PostInteractor{
			DB:   &repositories.DBRepository{DB: p.DB},
			Post: &repositories.PostRepository{},
		},
	}
}

func (c *PostsController) Get(ctx controllers.Context) {

}

func (c *PostsController) Post(ctx controllers.Context) {
	post := &entities.Posts{}
	if err := ctx.BindJSON(post); err != nil {
		ctx.JSON(400, apierrors.BadRequest.New(err.Error()))
		return
	}
	createdPost, res := c.Interactor.Create(post)
	if res.Error != nil {
		ctx.JSON(res.Code, apierrors.BadRequest.New(res.Error.Error()))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", createdPost))

}
