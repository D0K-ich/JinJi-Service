package network

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/valyala/fasthttp"
)

func NewServer(config *Config, git_ver string) (server *fasthttp.Server, err error) {

	if err = config.Validate(); err != nil {return}

	var main_router = CreateRouter(config.FilesPath, config.AccessToken)

	server = &fasthttp.Server{
		ReadBufferSize		: 4096,             // if we will have big cookies need to increase
		MaxRequestBodySize	: 10 * 1024 * 1024, // 10 MB limit
		Name				: fmt.Sprintf("kanopy/%s", git_ver),
		Handler				: fasthttp.CompressHandler(cors(main_router.Handler)),
		StreamRequestBody	: true,
	}

	// ========================
	// Start main route serving

	// SHADOW ERR
	go func() {
		if err := DefaultServer.ListenAndServe(config.Address); err != nil {
			log.Fatal("Failed to start rest server", zap.Any("error", err))
		}
	}()
	return
}
