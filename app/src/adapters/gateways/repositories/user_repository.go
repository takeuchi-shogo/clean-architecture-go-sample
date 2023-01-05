package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/takeuchi-shogo/clean-architecture-golang/src/entities"
)

type UserRepository struct{}

func (r *UserRepository) FindByID(db *gorm.DB, userID int) (user entities.Users, err error) {
	user = entities.Users{}
	db.First(&user, userID)
	if user.ID <= 0 {
		return entities.Users{}, err
	}
	return user, nil
}
