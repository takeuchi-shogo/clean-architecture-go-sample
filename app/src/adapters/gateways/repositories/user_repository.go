package repositories

import (
	"errors"
	"fmt"

	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (r *UserRepository) FindByID(db *gorm.DB, userID int) (user entities.Users, err error) {
	user = entities.Users{}
	db.First(&user, userID)
	if user.ID <= 0 {
		return entities.Users{}, err
	}
	// user.ID = 100
	// user.DisplayName = "test taro"
	return user, nil
}

func (r *UserRepository) FindByScreenName(db *gorm.DB, screenName string) (user entities.Users, err error) {
	user = entities.Users{}
	db.Where("screen_name = ?", screenName).First(&user)
	if user.ID <= 0 {
		return entities.Users{}, errors.New("user is not found")
	}
	// user.ID = 100
	// user.DisplayName = "test taro"
	return user, nil
}

func (r *UserRepository) Create(db *gorm.DB, user entities.Users) (createdUser entities.Users, err error) {
	createdUser = entities.Users{}
	// createdUser.DisplayName = user.DisplayName
	// createdUser.ScreenName = user.ScreenName
	// createdUser.Email = user.Email
	// createdUser.Password = user.Password
	// createdUser = user

	// err = user.Validate()

	createdUser = user
	fmt.Println(createdUser)

	// result := db.Create(&createdUser)

	return createdUser, nil
}
