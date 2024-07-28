package db

import (
	"fmt"
	"time"

	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"
	"gorm.io/driver/mysql"

	// _ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type DB struct {
	Connection *gorm.DB
}

func NewDB(c *config.Config) *DB {

	db := &DB{}
	// デバック中の為。コメントアウトしておく
	db.Connection = db.connect(
		c.Database.Production.Host,
		c.Database.Production.UserName,
		c.Database.Production.Password,
		c.Database.Production.DatabaseName,
	)

	return db
}

func (db *DB) connect(host string, username string, password string, db_name string) *gorm.DB {
	connection, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, db_name)), &gorm.Config{})
	if err != nil {
		count := 0
		for {
			fmt.Print(".\n")
			time.Sleep(time.Second)
			count++
			// connection wait 10 seconds for database starting...
			if count > 10 {
				fmt.Print("database connection failed\n")
				panic(err.Error())
			}
			connection, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, db_name)), &gorm.Config{})
			if err == nil {
				return connection
			}
		}
	}

	connection.Use(dbresolver.Register(dbresolver.Config{
		Policy: dbresolver.RandomPolicy{},
	}))
	return connection
}

func (db *DB) Conn() *gorm.DB {
	return db.Connection
}
