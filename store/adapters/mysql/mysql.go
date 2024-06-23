package mysql

import (
	"errors"
	dblog "log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Adapter struct {
	config		*Config
	db			*gorm.DB	// Use getter for get connection for private methods too
}

func NewAdapter(config *Config) (a *Adapter, err error) {
	if config == nil {err = errors.New("nil config for create adapter mysql");return}

	a = &Adapter{config : config}
	if a.db, err = gorm.Open(mysql.Open(config.Dsn), &gorm.Config{
		Logger	: logger.New(dblog.New(os.Stdout, "\r\n", dblog.LstdFlags), logger.Config{
			SlowThreshold	: 200 * time.Millisecond,
			LogLevel		: logger.Warn,
			Colorful		: true,
			IgnoreRecordNotFoundError: true,
		}),
	}); err != nil {return}

	return
}
func(a *Adapter) GetDB() (db *gorm.DB) {return a.db}