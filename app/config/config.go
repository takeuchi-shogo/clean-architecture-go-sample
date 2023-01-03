package config

type Config struct {
	Database struct {
		Host         string
		DatabaseName string
		UserName     string
		Password     string
	}
	Server struct {
		Port string
	}
}

func NewConfig() *Config {

	c := new(Config)

	c.Database.Host = ""
	c.Database.DatabaseName = ""
	c.Database.UserName = ""
	c.Database.Password = ""

	c.Server.Port = ""

	return c
}
