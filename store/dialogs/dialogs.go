package dialogs

import (
	"context"
	"github.com/D0K-ich/JinJi-Service/store/adapters/elastic"
)

func NewStorage(adapter *elastic.Adapter, context context.Context) (s *Storage, err error) {
	s = &Storage{
		elastic			: adapter,
		context			: context,
		indexDialogs	: "dialogs", //todo create type for name index
	}
	if err = s.elastic.EnsureIndexExist(s.indexDialogs, elastic.DefaultIndexSettings, FieldDialogsMap); err != nil {return}
	return
}

type Storage struct {
	context 		context.Context

	elastic 		*elastic.Adapter
	indexDialogs 	string
}

