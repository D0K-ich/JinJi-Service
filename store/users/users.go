package users

import (
	"context"
	"gorm.io/gorm"
)

func NewStorage(db *gorm.DB, context context.Context) (s *Storage, err error) {
	s = &Storage{
		db				: db,
		context			: context,
	}
	return
}

type Storage struct {
	context 	context.Context

	db      	*gorm.DB
}
