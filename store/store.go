package store

import (
	"context"

	"github.com/D0K-ich/KanopyService/logs"
	"github.com/D0K-ich/KanopyService/store/adapters/elastic"
	"github.com/D0K-ich/KanopyService/store/adapters/mysql"
	"github.com/D0K-ich/KanopyService/store/users"
)

var Default *Store
var log = logs.NewLog()

func NewStore(config *Config) (store *Store, err error) {
	if err = config.Validate(); err != nil {return}

	store = &Store{config: config}
	store.context, store.cancel = context.WithCancel(context.Background())

	log.Info("(store) >> CreateEcom mysql adapter...")
	if store.mysql, err = mysql.NewAdapter(config.Mysql); err != nil {return}

	log.Info("(store) >> CreateEcom elastic adapter...")
	if store.elastic, err = elastic.NewAdapter(store.context, config.Elastic); err != nil {return}

	//log.Info("(store) >> CreateEcom actions register domain...")
	//if store.Actions, err = actions.NewStorage(store.context, ); err != nil {return}


	log.Info("(store) >> Create users storage...")
	if store.Users, err = users.NewStorage(store.mysql.GetDB(), store.elastic, store.context); err != nil {return}



	return
}

type Store struct {
	config  *Config
	context context.Context
	cancel  context.CancelFunc

	// adapters
	//cache		*cache.Adapter
	mysql   *mysql.Adapter
	elastic *elastic.Adapter

	// domains
	//Actions		*actions.Storage
	//Fetch       *fetch.Storage
	//Dictionary  *dictionary.Storage
	//Parser      *parser.Storage
	Users       *users.Storage
}

func (s *Store) Close() {
	log.Info("(store) >> closing...")
	s.cancel()
}