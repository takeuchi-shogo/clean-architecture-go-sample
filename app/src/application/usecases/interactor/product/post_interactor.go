package product

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type PostInteractor struct {
	DB   repositories.DBRepository
	Post repositories.PostRepository
}

func (i *PostInteractor) Get(id string) (post *entities.Posts, resultStatus *usecases.ResultStatus) {

	db := i.DB.Conn()

	post, err := i.Post.FindByID(db, id)
	if err != nil {
		return &entities.Posts{}, usecases.NewResultStatus(404, []string{}, err)
	}
	return post, usecases.NewResultStatus(200, []string{}, nil)
}

func (i *PostInteractor) Create(post *entities.Posts) (newPost *entities.Posts, resultStatus *usecases.ResultStatus) {
	db := i.DB.Conn()

	newPost, err := i.Post.Create(db, post)
	if err != nil {
		return &entities.Posts{}, usecases.NewResultStatus(400, []string{}, err)
	}
	return newPost, usecases.NewResultStatus(200, []string{}, nil)
}
