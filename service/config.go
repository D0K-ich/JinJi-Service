package service

import (
	"errors"

	"github.com/D0K-ich/JinJi-Service/service/shikimory"
)

type Config struct {
	Shikimory *shikimory.Config	`yaml:"shikimory"`
}

func(c *Config) Validate() (err error) {
	if c == nil {return errors.New("empty service config")}

	if err = c.Shikimory.Validate(); err != nil {return}
	return
}