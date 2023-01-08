package product

import (
	"errors"
	"time"

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
	foundUser, err := i.User.FindByScreenName(user.ScreenName)
	if err != nil {
		return entities.Users{}, usecases.NewResultStatus(404, []string{}, err)
	}

	if utilities.CheckPasswordHash(user.Password, foundUser.Password) {
		return entities.Users{}, usecases.NewResultStatus(401, []string{"account.authorization"}, errors.New("failed autholization"))
	}

	claims, err := i.Jwt.ParseToken(jwtToken)
	if err != nil {
		return entities.Users{}, usecases.NewResultStatus(401, []string{"account.authorization"}, err)
	}
	if foundUser.ID != claims["userId"] {
		return entities.Users{}, usecases.NewResultStatus(401, []string{"account.authorization"}, errors.New("failed user"))
	}
	// check token
	if isTokenExpire(int64(claims["exp"].(float64))) {
		return entities.Users{}, usecases.NewResultStatus(401, []string{"auth.expireAt"}, errors.New("token is expire"))
	}

	return foundUser, usecases.NewResultStatus(200, []string{}, nil)
}

func isTokenExpire(expireAt int64) bool {
	currentTime := time.Now().Unix()
	if expireAt < currentTime {
		return true
	}
	return false
}
