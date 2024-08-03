package shikimory

import (
	"errors"
	"strings"
)

type Config struct {
	ClientId 		string 	`yaml:"client_id"`
	ClientSecret 	string 	`yaml:"client_secret"`
	RedirectUrl 	string 	`yaml:"redirect_url"`
}

func(c *Config) Validate() (err error) {
	if c == nil {return errors.New("empty shikimory config")}

	if c.ClientId 		= strings.TrimSpace(c.ClientId);		c.ClientId 		== "" {return errors.New("empty client_id in shikimory")}
	if c.RedirectUrl 	= strings.TrimSpace(c.RedirectUrl); 	c.RedirectUrl 	== "" {return errors.New("empty RedirectUrl in shikimory")}
	if c.ClientSecret 	= strings.TrimSpace(c.ClientSecret); 	c.ClientSecret 	== "" {return errors.New("empty ClientSecret in shikimory")}

	return
}