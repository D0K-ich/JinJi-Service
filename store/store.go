package store

import (
	"context"

	"github.com/D0K-ich/JinJi-Service/logs"

	"github.com/D0K-ich/JinJi-Service/store/adapters/elastic"
	"github.com/D0K-ich/JinJi-Service/store/adapters/mysql"
	"github.com/D0K-ich/JinJi-Service/store/users"
)

var Default *Store
var log = logs.NewLog()

func NewStore(config *Config) (store *Store, err error) {
	if err = config.Validate(); err != nil {
		return
	}

	store = &Store{config: config}
	store.context, store.cancel = context.WithCancel(context.Background())

	log.Info("(store) >> CreateEcom mysql adapter...")
	if store.mysql, err = mysql.NewAdapter(config.Mysql); err != nil {
		return
	}

	log.Info("(store) >> Create users storage...")
	if store.Users, err = users.NewStorage(store.mysql.GetDB(), store.elastic, store.context); err != nil {
		return
	}

	return
}

type Store struct {
	config  *Config
	context context.Context
	cancel  context.CancelFunc

	mysql   *mysql.Adapter
	elastic *elastic.Adapter

	Users *users.Storage
}

func (s *Store) Close() {
	log.Info("(store) >> closing...")
	s.cancel()
}
