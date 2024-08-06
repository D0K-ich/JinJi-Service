package store

import (
	"context"
	"github.com/D0K-ich/JinJi-Service/store/dialogs"
	"github.com/rs/zerolog/log"

	"github.com/D0K-ich/JinJi-Service/store/users"
	"github.com/D0K-ich/JinJi-Service/store/adapters/elastic"
	"github.com/D0K-ich/JinJi-Service/store/adapters/mysql"
)

var Default *Store

func NewStore(config *Config) (store *Store, err error) {
	if err = config.Validate(); err != nil {return}

	store = &Store{config: config}
	store.context, store.cancel = context.WithCancel(context.Background())

	log.Info().Msg("(store) >> Create mysql adapter...")
	if store.mysql, err = mysql.NewAdapter(config.Mysql); err != nil {return}

	log.Info().Msg("(store) >> Create elastic adapter...")
	if store.elastic, err = elastic.NewAdapter(store.context, config.Elastic); err != nil {return}

	log.Info().Msg("(store) >> Create users storage...")
	if store.Users, err = users.NewStorage(store.mysql.GetDB(), store.context); err != nil {return}

	log.Info().Msg("(store) >> Create dialogs storage...")
	if store.Dialogs, err = dialogs.NewStorage(store.elastic, store.context); err != nil {return}

	return
}

type Store struct {
	config  	*Config
	context 	context.Context
	cancel  	context.CancelFunc

	mysql   	*mysql.Adapter
	elastic 	*elastic.Adapter

	Users 		*users.Storage
	Dialogs		*dialogs.Storage
}

func (s *Store) Close() {
	log.Info().Msg("(store) >> closing...")
	s.cancel()
}
