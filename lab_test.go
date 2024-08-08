package JinJi_Service

import (
	"fmt"
	"github.com/rs/zerolog/log"

	"testing"

)

func TestLab(t *testing.T) {
	fmt.Println("All  is okey")
	log.Fatal().Msgf("asda")
	log.Info().Msgf("asdasd")
}
