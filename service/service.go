package service

import (
	"github.com/D0K-ich/JinJi-Service/service/neo"
	"github.com/D0K-ich/JinJi-Service/service/shikimory"
	"github.com/rs/zerolog/log"
)

var Default *Services

type Services struct {
	//Local Services
	NeoJinJi 	*neo.NeoJinJi

	//Remotly services
	Shikimory 	*shikimory.Shikimory
}

func NewServices(config *Config) (services *Services, err error) {
	if err = config.Validate(); err != nil {return}

	services = &Services{}

	log.Info().Msgf("(service) >> Creating shiki service...")
	if services.Shikimory, err = shikimory.NewShiki(config.Shikimory); err != nil {return}

	log.Info().Msgf("(service) >> Creating neojinji service...")
	if services.NeoJinJi, err = neo.NewNeo(config.NeoJinJi); err != nil {return}

	return
}