package output

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UserInteractorOutput interface {
	ViewUser(user entities.Users, resultStatus *usecases.ResultStatus)
	ViewUserList()
}
