package gpt

import (
	"time"
	"errors"
	"strings"
)

type Config struct {
	RootPath			string			`yaml:"root_path"`
	Model 				string			`yaml:"model"`
	Temperature 		float64			`yaml:"temperature"`
	MaxTokens			int				`yaml:"max_tokens"`
	Token 				string			`yaml:"token"`
	Beams 				string			`yaml:"beams"`

	TimeoutResponseMin 	time.Duration 	`json:"timeout_response_min"`
	//PersonRole			MessagesGPT		`json:"person_role"`
}


func(c *Config) Validate() (err error) {
	if c == nil {return errors.New("nil gpt config")}

	if c.RootPath = strings.TrimSpace(c.RootPath); c.RootPath == "" {return errors.New("invalid address for gpt")}
	if c.Token = strings.TrimSpace(c.Token); c.Token == "" {return errors.New("invalid Token for gpt")}
	if c.Model = strings.TrimSpace(c.Model); c.Model == "" {return errors.New("invalid model for gpt")}
	if c.Temperature == 0.0 {return errors.New("invalid temperature for gpt")}

	return
}

func defaultConfig() *Config {return &Config{
	RootPath			: defaultRootPath,
	Token				: defaultToken,
	Model				: defaultModel,
	Temperature			: defaultTemp,
	TimeoutResponseMin	: defaultTimeout,
	MaxTokens			: defaultMaxTokens,

}}