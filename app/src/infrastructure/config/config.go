package config

import "github.com/takeuchi-shogo/clean-architecture-golang/lib"

type Config struct {
	Environment string
	Database    struct {
		Production struct {
			Host         string
			DatabaseName string
			UserName     string
			Password     string
		}
	}
	Server struct {
		Port string
	}
}

func NewConfig(env lib.Env) *Config {

	c := new(Config)

	c.Environment = env.Environment

	c.Database.Production.Host = env.DBHost
	c.Database.Production.DatabaseName = env.DBName
	c.Database.Production.UserName = env.DBUsername
	c.Database.Production.Password = env.DBPassword

	c.Server.Port = ":" + env.ServerPort

	return c
}
