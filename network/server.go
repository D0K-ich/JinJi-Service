package network

import (
	"fmt"
	"github.com/D0K-ich/JinJi-Service/network/websocket"
	"github.com/fasthttp/session/v2"

	"go.uber.org/zap"

	"github.com/valyala/fasthttp"
)

func NewServer(config *Config, git_ver string, user_session *session.Session) (server *fasthttp.Server, err error) {
	log.Info("(network) >> Creating new server...")
	if err = config.Validate(); err != nil {
		return
	}

	var main_router = CreateRouter(config.FilesPath, config.AccessToken, user_session)

	server = &fasthttp.Server{
		ReadBufferSize:     4096,             // if we will have big cookies need to increase
		MaxRequestBodySize: 10 * 1024 * 1024, // 10 MB limit
		Name:               fmt.Sprintf("kanopy/%s", git_ver),
		Handler:            fasthttp.CompressHandler(cors(main_router.Handler)),
		StreamRequestBody:  true,
	}

	go func() {
		if err := server.ListenAndServe(config.Address); err != nil {
			log.Fatal("Failed to start rest server", zap.Any("error", err))
		}
	}()

	go func() {
		if err := websocket.NewWsConnection(); err != nil {
			log.Fatal("Failed to start rest server", zap.Any("error", err))
		}
	}()

	return
}
