package input

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UserInteractorInput interface {
	Get(userID int) (user entities.Users, resultStatus *usecases.ResultStatus)
}
