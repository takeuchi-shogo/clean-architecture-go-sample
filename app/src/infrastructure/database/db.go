package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/takeuchi-shogo/clean-architecture-golang/src/infrastructure/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type DB struct {
	Connection *gorm.DB
}

type MasterDBConfig struct {
	Host         string
	UserName     string
	Password     string
	DatabaseName string
}

type SlaveDBConfig struct {
	Host         string
	UserName     string
	Password     string
	DatabaseName string
}

func NewDB(c *config.Config) *DB {

	db := &DB{}
	// デバック中の為。コメントアウトしておく
	db.Connection = db.connect(
		c.Database.Master.Host,
		c.Database.Master.UserName,
		c.Database.Master.Password,
		c.Database.Master.DatabaseName,
		SlaveDBConfig{
			Host:         c.Database.Slave.Host,
			UserName:     c.Database.Slave.UserName,
			Password:     c.Database.Slave.Password,
			DatabaseName: c.Database.Slave.DatabaseName,
		},
	)

	return db
}

func (db *DB) connect(host string, username string, password string, db_name string, slave SlaveDBConfig) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Silent,    // Log level
			IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,             // Don't include params in the SQL log
			Colorful:                  false,            // Disable color
		},
	)
	connection, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, db_name)), &gorm.Config{
		Logger: newLogger,
	})
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
			connection, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, db_name)), &gorm.Config{
				Logger: newLogger,
			})
			if err == nil {
				return connection
			}
		}
	}

	connection.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", slave.UserName, slave.Password, slave.Password, slave.DatabaseName))},
		Policy:   dbresolver.RandomPolicy{},
	}))

	return connection
}

func (db *DB) Conn() *gorm.DB {
	return db.Connection
}
