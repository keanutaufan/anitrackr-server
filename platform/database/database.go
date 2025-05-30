package database

import (
	"fmt"
)

type Config struct {
	Protocol string
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SslMode  string
}

func (conf *Config) Dsn() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		conf.Protocol,
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.SslMode,
	)
}
