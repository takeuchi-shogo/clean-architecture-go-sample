package product

import (
	"errors"

	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/gateways"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/repositories"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/usecases"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/application/utilities"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UserAuthInteractor struct {
	Jwt  gateways.JwtMiddleware
	User repositories.UserRepository
}

// Check screenName and password after check jwt
func (i *UserAuthInteractor) Autholization(user entities.Users, jwtToken string) (foundUser entities.Users, resultStatus *usecases.ResultStatus) {
	// return i.Jwt.
	foundUser, err := i.User.FindByScreenName(user.ScreenName)
	if err != nil {
		return entities.Users{}, usecases.NewResultStatus(404, []string{}, err)
	}

	if utilities.CheckPasswordHash(user.Password, foundUser.Password) {
		return entities.Users{}, usecases.NewResultStatus(401, []string{"account.autholization"}, errors.New("failed autholization"))
	}

	if i.Jwt.CheckJwtToken(jwtToken) {
		return entities.Users{}, usecases.NewResultStatus(401, []string{"account.autholization"}, errors.New("test auth error"))
	}

	return foundUser, usecases.NewResultStatus(200, []string{}, nil)
}
