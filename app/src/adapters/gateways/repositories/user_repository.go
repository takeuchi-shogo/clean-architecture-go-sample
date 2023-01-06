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

func (r *UserRepository) Create(user entities.Users) (createdUser entities.Users, err error) {
	createdUser = entities.Users{}
	createdUser.DisplayName = user.DisplayName

	return createdUser, nil
}
