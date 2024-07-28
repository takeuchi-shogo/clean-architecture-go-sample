package repositories

import "gorm.io/gorm"

type DB interface {
	Conn() *gorm.DB
}
