package store

import (
	"errors"

	"github.com/D0K-ich/JinJi-Service/store/adapters/elastic"
	"github.com/D0K-ich/JinJi-Service/store/adapters/mysql"
)

type Config struct {
	Mysql   *mysql.Config   `yaml:"mysql"`
	Elastic *elastic.Config `yaml:"elastic"`
}

func (c *Config) Validate() (err error) {
	if c == nil {
		return errors.New("nil config")
	}

	if err = c.Mysql.Validate(); err != nil {
		return err
	}
	if err = c.Elastic.Validate(); err != nil {
		return err
	}

	return
}
