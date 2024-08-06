package elastic

import (
	"fmt"
	"sync"
	"errors"
	"strings"
)

type IndexKey string
func(i IndexKey) String()	string  {return string(i)}
func(i IndexKey) IndexKey() IndexKey {return i}
func(i IndexKey) Validate() error {
	if i.String() == "" {return errors.New("empty index key")}
	return nil
}

// the keys are binded to config, do not change
const (
	IndexKeyFetchDialogs IndexKey = "dialogs" // "manager.fetch.profiles-1"
)

var validIndexKeys = []IndexKey{IndexKeyFetchDialogs}

type Config struct {
	lock		sync.RWMutex
	Host		string                 	`yaml:"host"`
	Indices		map[IndexKey]string 	`yaml:"indices"` // index_key => current_index_name
	//Prefixes	map[string]string	`yaml:"prefixes"`
}

func(c *Config) Validate() (err error) {
	if c == nil {return errors.New("nil elastic config")}
	if len(c.Indices) == 0 {return errors.New("empty index names config")}

	c.lock.Lock()
	defer c.lock.Unlock()

	for _, idx_key := range validIndexKeys {
		if c.Indices[idx_key] = strings.TrimSpace(c.Indices[idx_key]); c.Indices[idx_key] == "" {
			return fmt.Errorf("empty index for key %s", idx_key)
		}
	}

	return
}