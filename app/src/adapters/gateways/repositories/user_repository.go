package repositories

import (
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UserRepository struct{}

func (r *UserRepository) FindByID(userID int) (user entities.Users, err error) {
	user = entities.Users{}
	// db.First(&user, userID)
	// if user.ID <= 0 {
	// 	return entities.Users{}, err
	// }
	user.ID = 100
	user.DisplayName = "test taro"
	return user, nil
}

func (r *UserRepository) FindByScreenName(screenName string) (user entities.Users, err error) {
	user = entities.Users{}
	// db.Where("sreen_name = ?", screenName).First(&user)
	// if user.ID <= 0 {
	// 	return entities.Users{}, err
	// }
	user.ID = 100
	user.DisplayName = "test taro"
	return user, nil
}

func (r *UserRepository) Create(user entities.Users) (createdUser entities.Users, err error) {
	createdUser = entities.Users{}
	// createdUser.DisplayName = user.DisplayName
	// createdUser.ScreenName = user.ScreenName
	// createdUser.Email = user.Email
	// createdUser.Password = user.Password
	createdUser = user

	return createdUser, nil
}
