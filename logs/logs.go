package logs

import (
	"os"
	"errors"
	"strconv"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Level string `yaml:"level"`
	//Formatter 	*logrus.Formatter 	`yaml:"formatter"`
	Output string `yaml:"encoding"`
}

func (c *Config) Validate() (err error) {
	if c == nil {return errors.New("nil log config")}

	if c.Level == "" {return errors.New("nil level config")}
	if c.Output == "" {return errors.New("nil Output config")}
	return
}

func SetConf(config *Config) (err error) {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})

	return
}