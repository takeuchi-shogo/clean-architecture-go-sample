package repositories

import "gorm.io/gorm"

type DBRepository interface {
	Conn() *gorm.DB
}
