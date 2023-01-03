package db

import (
	"fmt"

	"github.com/takeuchi-shogo/clean-architecture-golang/config"

	"github.com/jinzhu/gorm"
)

type DB struct {
	Connection *gorm.DB
}

func NewDB(c *config.Config) *DB {

	db := &DB{}
	db.Connection = db.connect(
		c.Database.Production.Host,
		c.Database.Production.UserName,
		c.Database.Production.Password,
		c.Database.Production.DatabaseName,
	)

	return db
}

func (db *DB) connect(host string, username string, password string, db_name string) *gorm.DB {
	connection, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, db_name))
	if err != nil {
		panic(err.Error())
	}
	return connection
}
