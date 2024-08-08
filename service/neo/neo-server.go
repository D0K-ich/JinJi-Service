package neo

import (
	"context"
	"github.com/D0K-ich/types/uuid"
	"github.com/elastic/go-elasticsearch/v8/typedapi/tasks/cancel"
)

type NeoJinJi struct {
	config 		*Config

	context 	context.Context
	cancel 		cancel.Cancel

	Version 	string

	UserUuid    uuid.UserUuid
	//Store todo quadrant
}

func NewNeo(config *Config) (neo_server *NeoJinJi, err error) {
	//todo request to load model, and health check


	return
}

func(n *NeoJinJi) HealhCheck() (err error) {

	return
}