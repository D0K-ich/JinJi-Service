package mysql

import (
	"log"
	"os"
	"time"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
)

type Adapter struct {
	config		*Config
	db			*gorm.DB	// Use getter for get connection for private methods too
}

func NewAdapter(config *Config) (a *Adapter, err error) {
	if config == nil {err = errors.New("nil config for create adapter mysql");return}

	a = &Adapter{config : config}
	if a.db, err = gorm.Open(mysql.Open(config.Dsn), &gorm.Config{
		Logger	: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold	: 200 * time.Millisecond,
			LogLevel		: logger.Warn,
			Colorful		: true,
			IgnoreRecordNotFoundError: true,
		}),
	}); err != nil {return}

	return
}
func(a *Adapter) GetDB() (db *gorm.DB) {return a.db}