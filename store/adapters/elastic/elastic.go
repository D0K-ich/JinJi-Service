package elastic

import (
	"context"
	"errors"
	"github.com/D0K-ich/KanopyService/logs"
	"go.uber.org/zap"
	"time"

	E8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/info"
)

var log = logs.NewLog()

func NewAdapter(parent_ctx context.Context, config *Config) (a *Adapter, err error) {

	if err = config.Validate(); err != nil {return}

	a = &Adapter{
		config	:  config,
		context	: parent_ctx,
	}

	if a.Client, err = E8.NewTypedClient(E8.Config{
		Addresses		: []string{config.Host},
		RetryBackoff	: func(i int) time.Duration { return time.Duration(i) * 500 * time.Millisecond },
		MaxRetries		: 5,
		//Transport		: &FastTransport{},	 //while not working, freezein on some request
	}); err != nil {return}

	var ping_ok bool
	if ping_ok, err = a.Client.Core.Ping().IsSuccess(a.context); err != nil {return}
	if !ping_ok {err = errors.New("elastic host ping failure");return}

	var info_resp *info.Response
	if info_resp, err = a.Client.Info().Do(a.context); err != nil {return}
	log.Info("(elastic) >> Adapter created",
		zap.Any("version", info_resp.Version),
		zap.Any("name", info_resp.Name),
		zap.Any("cluster", info_resp.ClusterName))

	return
}

type Adapter struct {
	config  *Config
	context context.Context

	Client *E8.TypedClient
}