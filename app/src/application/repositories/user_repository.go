package repositories

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UserRepository interface {
	FindByID(userID int) (user entities.Users, err error)
}
