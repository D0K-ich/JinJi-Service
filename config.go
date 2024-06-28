package KanopyService

import (
	"errors"
	"fmt"
	"github.com/D0K-ich/KanopyService/gpt"
	"github.com/D0K-ich/KanopyService/network"
	"github.com/D0K-ich/KanopyService/network/rest"
	"os"
	"strings"

	"github.com/kr/pretty"
	"gopkg.in/yaml.v3"

	"github.com/D0K-ich/KanopyService/logs"
	"github.com/D0K-ich/KanopyService/store"
)

type Config struct {
	Logger		*logs.Config		`yaml:"logger"`
	Store		*store.Config		`yaml:"store"`
	Server		*network.Config		`yaml:"network"`
	Gpt 		*gpt.Config			`yaml:"gpt"`
	Rest 		*rest.Config		`yaml:"rest"`
}

func NewConfig(path *string) (config *Config, err error) {
	if *path = strings.TrimSpace(*path); *path == "" {err = errors.New("empty path for create main conf");return}

	var bytes []byte
	if bytes, err = os.ReadFile(*path); err != nil {return}

	if err = yaml.Unmarshal(bytes, &config); err != nil {return}

	if err = config.Validate(); err != nil {return}

	return
}
func(c *Config) Print() {fmt.Printf("%# v\n", pretty.Formatter(c))}

func(c *Config) Validate() (err error) {
	if err = c.Server.Validate(); err != nil {return}
	if err = c.Store.Validate(); err != nil {return}
	if err = c.Logger.Validate(); err != nil {return}
	if err = c.Gpt.Validate(); err != nil {return}

	return
}