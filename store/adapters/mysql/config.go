package mysql

import (
	"errors"
	"strings"
)

type Config struct {
	Dsn		string		`yaml:"dsn"`
	Debug	bool		`yaml:"debug"`
}

func(c *Config) Validate() (err error) {
	if c == nil {return errors.New("nil mysql config")}
	if c.Dsn = strings.TrimSpace(c.Dsn); c.Dsn == "" {return errors.New("empty dsn string")}
	return
}
