package input

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UserInteractorInput interface {
	Get(userID int) (user entities.Users, resultStatus *usecases.ResultStatus)
	Create(user *entities.Users) (newUser entities.Users, resultStatus *usecases.ResultStatus)
	Save(user *entities.Users) (updatedUser *entities.Users, resultStatus *usecases.ResultStatus)
}
