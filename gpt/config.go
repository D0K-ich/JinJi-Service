package gpt

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Config struct {
	Port 				int				`json:"port"`
	Address 			string			`json:"address"`
	Model 				string			`json:"model"`
	Temperature 		float64			`json:"temperature"`
	MaxTokens			int				`json:"max_tokens"`

	TimeoutResponseMin 	time.Duration 	`json:"timeout_response_min"`
	PersonRole			MessagesGPT		`json:"person_role"`
}


func(c *Config) Validate() (err error) {
	if c == nil {return errors.New("nil gpt config")}

	if c.Port == 0 {return errors.New("invalid port for gpt")}
	if c.Address = strings.TrimSpace(c.Address); c.Address == "" {return errors.New("invalid address for gpt")}
	if c.Model = strings.TrimSpace(c.Model); c.Model == "" {return errors.New("invalid model for gpt")}
	if c.Temperature == 0.0 {return errors.New("invalid temperature for gpt")}

	return
}

func defaultConfig() *Config {return &Config{
	Port				: defaultPort,
	Address				: defaultAddress,
	Model				: defaultModel,
	Temperature			: defaultTemp,
	TimeoutResponseMin	: defaultTimeout,
	MaxTokens			: defaultMaxTokens,

	PersonRole			: MessagesGPT{{
		Role	: "system",
		Content	: "Always answer in rhymes.",
	}},
}}

func(c *Config) serverAddress() string {return fmt.Sprintf(c.Address, c.Port)}