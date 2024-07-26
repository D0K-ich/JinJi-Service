package users

import (
	"context"

	"gorm.io/gorm"

	"github.com/D0K-ich/JinJi-Service/store/adapters/elastic"
)

func NewStorage(db *gorm.DB, adapter *elastic.Adapter, context context.Context) (s *Storage, err error) {
	s = &Storage{
		db:         db,
		elastic:    adapter,
		context:    context,
		indexUsers: "users", //todo create type for name index
	}
	//if err = s.elastic.EnsureIndexExist(s.indexUsers, tpls.DefaultIndexSettings, billing.FieldTransactionsMap); err != nil {return}//todo create elastick for ... what?
	return
}

type Storage struct {
	context context.Context
	db      *gorm.DB
	elastic *elastic.Adapter

	indexUsers string
}
