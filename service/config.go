package service

import (
	"errors"

	"github.com/D0K-ich/JinJi-Service/service/neo"

	"github.com/D0K-ich/JinJi-Service/service/shikimory"
)

type Config struct {
	NeoJinJi 	*neo.Config			`yaml:"neo_jin_ji"`

	Shikimory 	*shikimory.Config	`yaml:"shikimory"`
}

func(c *Config) Validate() (err error) {
	if c == nil {return errors.New("empty service config")}

	if err = c.NeoJinJi.Validate(); err != nil {return}
	if err = c.Shikimory.Validate(); err != nil {return}

	return
}