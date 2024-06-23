package network

import (
	"errors"
	"strings"
)

type Config struct {
	Address     	string 		`yaml:"address"`
	FilesPath   	string 		`yaml:"files_path"`
	AccessToken 	string 		`yaml:"access_token"`
	Session     	*Session 	`yaml:"session"`
}

type Session struct {
	Host     		string 		`yaml:"host"`
	Port     		int    		`yaml:"port"`
	User     		string 		`yaml:"user"`
	Password 		string 		`yaml:"password"`
	DbName   		string 		`yaml:"dbname"`
}

func (c *Config) Validate() (err error) {
	if c == nil {return errors.New("nil rest config")}
	if c.Address = strings.TrimSpace(c.Address); c.Address == "" {return errors.New("empty listen address")}
	if c.AccessToken = strings.TrimSpace(c.AccessToken); len(c.AccessToken) < 10 {return errors.New("to short admin token")}
	return
}
