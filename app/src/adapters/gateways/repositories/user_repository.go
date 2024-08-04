package repositories

import (
	"errors"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (r *UserRepository) FindByID(db *gorm.DB, userID string) (user *entities.Users, err error) {
	user = &entities.Users{}
	db.Debug().Where("id = ?", userID).Preload("Posts").First(user)
	if user.ID == "" {
		return &entities.Users{}, errors.New("user is not found")
	}
	return user, nil
}

func (r *UserRepository) FindByScreenName(db *gorm.DB, screenName string) (user *entities.Users, err error) {
	user = &entities.Users{}
	db.Where("screen_name = ?", screenName).First(user)
	if user.ID <= "" {
		return &entities.Users{}, errors.New("user is not found")
	}
	// user.ID = 100
	// user.DisplayName = "test taro"
	return user, nil
}

func (r *UserRepository) Create(db *gorm.DB, user *entities.Users) (createdUser *entities.Users, err error) {
	createdUser = &entities.Users{}

	err = user.Validate()
	if err != nil {
		return &entities.Users{}, err
	}

	createdUser = user
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	createdUser.ID = ulid.MustNew(ulid.Timestamp(t), entropy).String()

	err = db.Omit("Posts.*").Create(createdUser).Error

	return createdUser, err
}

func (r *UserRepository) Save(db *gorm.DB, user *entities.Users) (updatedUser *entities.Users, err error) {
	err = user.Validate()
	if err != nil {
		return &entities.Users{}, err
	}

	err = db.Omit("Posts.*").Save(user).Error
	return user, err
}
