package repositories

import "gorm.io/gorm"

type DBRepository struct {
	DB DB
}

func (r *DBRepository) Conn() *gorm.DB {
	return r.DB.Conn()
}
