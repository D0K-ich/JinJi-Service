package rest

import (
	"github.com/fasthttp/session/v2"
	"github.com/fasthttp/session/v2/providers/mysql"
	"github.com/valyala/fasthttp"
	"time"
)

func NewSession(config *Config, cookie_name string, table_name string) (serverSession *session.Session, err error) {
	var cfg = session.NewDefaultConfig()

	cfg.Secure = true
	cfg.Expiration = time.Hour * 24 * 365 * 10
	cfg.CookieSameSite = fasthttp.CookieSameSiteNoneMode
	cfg.CookieName = cookie_name

	serverSession = session.New(cfg)

	var cfg_mysql = mysql.NewConfigWith(config.Session.Host, config.Session.Port, config.Session.User, config.Session.Password, config.Session.DbName, table_name)
	var provider *mysql.Provider
	if provider, err = mysql.New(cfg_mysql); err != nil {return}

	if err = serverSession.SetProvider(*provider); err != nil {return}
	return
}