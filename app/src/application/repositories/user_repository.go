package repositories

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UserRepository interface {
	FindByID(userID int) (user entities.Users, err error)
	FindByScreenName(screenName string) (user entities.Users, err error)
	Create(user entities.Users) (createdUser entities.Users, err error)
}
