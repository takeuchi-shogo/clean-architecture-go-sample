package config

type Config struct {
	Database struct {
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

func NewConfig() *Config {

	c := new(Config)

	c.Database.Production.Host = ""
	c.Database.Production.DatabaseName = ""
	c.Database.Production.UserName = ""
	c.Database.Production.Password = ""

	c.Server.Port = ""

	return c
}
