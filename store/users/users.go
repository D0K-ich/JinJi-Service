package users

import (
	"context"

	"gorm.io/gorm"

	"github.com/D0K-ich/KanopyService/store/adapters/elastic"
)

func NewStorage(db *gorm.DB, adapter *elastic.Adapter, context context.Context) (s *Storage, err error) {
	s = &Storage{
		db		: db,
		elastic	: adapter,
		context	: context,
	}
	//if err = s.elastic.EnsureIndexExist(s.indexTransactions, tpls.DefaultIndexSettings, billing.FieldTransactionsMap); err != nil {return}
	return
}

type Storage struct {
	context		context.Context
	db			*gorm.DB
	elastic		*elastic.Adapter
}
