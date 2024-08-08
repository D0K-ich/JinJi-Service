package neo

import (
	"errors"
	"strings"
	"time"
)

type Config struct {
	Address 		string			`yaml:"address"`
	Port 			int				`yaml:"port"`
	Version			string			`yaml:"version"`

	Temperature 	float64			`yaml:"temperature"` 	//creation
	MaxLenght 		int				`yaml:"max_lenght"`		// count max tokens
	Beams			int				`yaml:"beams"`			//

	AccessToken     string			`yaml:"access_token"`

	PythonVersion 	string			`yaml:"python_version"`

	Timeout			time.Duration	`yaml:"timeout"`
}

func(c *Config) Validate() (err error) {
	if c == nil {return errors.New("empty neo config")}

	if c.Address = strings.TrimSpace(c.Address); c.Address == "" {return errors.New("empty neo address")}
	if c.Version = strings.TrimSpace(c.Version); c.Version == "" {return errors.New("empty neo Version")}
	if c.AccessToken = strings.TrimSpace(c.AccessToken); c.AccessToken == "" {return errors.New("empty neo AccessToken")}

	if c.Temperature 	== 0 {return errors.New("neo Temperature must be greaten than 0")}
	if c.MaxLenght 		== 0 {return errors.New("neo MaxLenght must be greaten than 0")}
	if c.Beams 			== 0 {return errors.New("neo Beams must be greaten than 0")}
	if c.Port 			== 0 {return errors.New("neo Port must be greaten than 0")}

	return
}