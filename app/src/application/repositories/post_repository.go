package repositories

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
	"gorm.io/gorm"
)

type PostRepository interface {
	FindByID(db *gorm.DB, id string) (post *entities.Posts, err error)
	FindByUserID(db *gorm.DB, userID string) (posts []*entities.Posts, err error)
	Create(db *gorm.DB, post *entities.Posts) (newPost *entities.Posts, err error)
}
