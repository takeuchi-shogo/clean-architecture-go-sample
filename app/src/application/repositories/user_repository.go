package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UserRepository interface {
	FindByID(db *gorm.DB, userID int) (user entities.Users, err error)
}
