package network

import (
	"fmt"
	"github.com/rs/zerolog/log"

	"github.com/valyala/fasthttp"
	"github.com/fasthttp/session/v2"

	"github.com/D0K-ich/JinJi-Service/network/websocket"
)

func NewServer(config *Config, git_ver string, user_session *session.Session) (server *fasthttp.Server, err error) {
	log.Info().Msg("(network) >> Creating new server...")
	if err = config.Validate(); err != nil {return}

	var main_router = CreateRouter(config.FilesPath, config.AdminToken, user_session)

	server = &fasthttp.Server{
		ReadBufferSize		: 4096,             // if we will have big cookies need to increase
		MaxRequestBodySize	: 10 * 1024 * 1024, // 10 MB limit
		Name				: fmt.Sprintf("jinji/%s", git_ver),
		Handler				: fasthttp.CompressHandler(cors(main_router.Handler, config.AdminToken)),
		StreamRequestBody	: true,
	}

	go func() {
		if err := server.ListenAndServe(config.Address); err != nil {
			log.Fatal().Msgf("Failed to start rest server", "error", err)
		}
	}()

	go func() {
		if err := websocket.NewWsConnection(); err != nil {
			log.Fatal().Msgf("Failed to start rest server", "error", err)
		}
	}()

	return
}
