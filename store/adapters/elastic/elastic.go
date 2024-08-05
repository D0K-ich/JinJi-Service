package elastic

import (
	"time"
	"errors"
	"context"
	"strings"
	"unicode/utf8"

	"github.com/rs/zerolog/log"

	E8 "github.com/elastic/go-elasticsearch/v8"
	ET "github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/info"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
)

func NewAdapter(parent_ctx context.Context, config *Config) (a *Adapter, err error) {

	if err = config.Validate(); err != nil {return}

	a = &Adapter{
		config:  config,
		context: parent_ctx,
	}

	if a.Client, err = E8.NewTypedClient(E8.Config{
		Addresses:    []string{config.Host},
		RetryBackoff: func(i int) time.Duration { return time.Duration(i) * 500 * time.Millisecond },
		MaxRetries:   5,
		//Transport		: &FastTransport{},	 //while not working, freezein on some request
	}); err != nil {return}

	var ping_ok bool
	if ping_ok, err = a.Client.Core.Ping().IsSuccess(a.context); err != nil {return}
	if !ping_ok {err = errors.New("elastic host ping failure");return}

	var info_resp *info.Response
	if info_resp, err = a.Client.Info().Do(a.context); err != nil {return}
	log.Info().Msgf("(elastic) >> Adapter created",
		"version", info_resp.Version,
		"name", info_resp.Name,
		"cluster", info_resp.ClusterName)

	return
}

type Adapter struct {
	config  *Config
	context context.Context

	Client *E8.TypedClient
}

func (a *Adapter) EnsureIndexExist(index_name string, settings *ET.IndexSettings, mapping *ET.TypeMapping) (err error) {

	// check index name
	if index_name = strings.TrimSpace(index_name); utf8.RuneCountInString(index_name) < 3 {return errors.New("empty or to short index name")}
	if mapping == nil {return errors.New("nil index mapping")}

	// check index already created
	var exist bool
	if exist, err = a.Client.Indices.Exists(index_name).Do(a.context); err != nil {return}
	if exist {return} // todo add mapping comparation

	log.Info().Msgf("(elastic) >> need to create index", "name", index_name)

	if settings == nil {return errors.New("get nil index settings")}

	var response *create.Response
	if response, err = a.Client.Indices.Create(index_name).
		Settings(settings).
		Mappings(mapping).
		Do(a.context); err != nil {
		return
	}

	log.Info().Msgf("(elastic) >> index created", "name", response.Index)
	return
}
