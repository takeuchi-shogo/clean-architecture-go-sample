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
	DB   repositories.DBRepository
	Jwt  gateways.JwtMiddleware
	User repositories.UserRepository
}

type AuthToken struct {
	JwtToken string
	User     *entities.Users
}

// Check screenName and password after check jwt
func (i *UserAuthInteractor) Autholization(jwtToken string) (foundUser *entities.Users, resultStatus *usecases.ResultStatus) {

	claims, err := i.Jwt.ParseToken(jwtToken)

	if err != nil {
		return &entities.Users{}, usecases.NewResultStatus(401, []string{"account.authorization"}, err)
	}
	if isTokenExpire(int64(claims["exp"].(float64))) {
		return &entities.Users{}, usecases.NewResultStatus(401, []string{"auth.expireAt"}, errors.New("token is expire"))
	}

	db := i.DB.Conn()

	foundUser, err = i.User.FindByID(db, claims["aud"].(string))
	if err != nil {
		return &entities.Users{}, usecases.NewResultStatus(404, []string{}, err)
	}
	if foundUser.ID != claims["aud"] {
		return &entities.Users{}, usecases.NewResultStatus(401, []string{"account.authorization"}, errors.New("failed user"))
	}

	return foundUser, usecases.NewResultStatus(200, []string{}, nil)
}

func (i *UserAuthInteractor) Create(user entities.Users) (auth AuthToken, resultStatus *usecases.ResultStatus) {
	db := i.DB.Conn()

	foundUser, err := i.User.FindByScreenName(db, user.ScreenName)
	if err != nil {
		return AuthToken{}, usecases.NewResultStatus(400, []string{}, errors.New("user is not found"))
	}

	if utilities.CheckPasswordHash(user.Password, foundUser.Password) {
		return AuthToken{}, usecases.NewResultStatus(400, []string{}, errors.New("test auth error"))
	}

	token := i.Jwt.CreateToken(foundUser.ID)

	return AuthToken{JwtToken: token, User: foundUser}, usecases.NewResultStatus(200, []string{}, nil)
}

func isTokenExpire(expireAt int64) bool {
	currentTime := time.Now().Unix()
	return expireAt < currentTime
}
