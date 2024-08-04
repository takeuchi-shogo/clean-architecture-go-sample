package repositories

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(db *gorm.DB, userID string) (user *entities.Users, err error)
	FindByScreenName(db *gorm.DB, screenName string) (user *entities.Users, err error)
	Create(db *gorm.DB, user *entities.Users) (createdUser *entities.Users, err error)
	Save(db *gorm.DB, user *entities.Users) (updatedUser *entities.Users, err error)
}
