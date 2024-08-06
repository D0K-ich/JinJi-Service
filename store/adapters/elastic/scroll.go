package elastic

import (
	"context"
	"errors"
	ET "github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/scroll"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/clearscroll"
	"github.com/rs/zerolog/log"
	"strings"
	"unicode/utf8"
)

const scrollDuration = "1m"
type ScrollCb = func(total, fetched int, hits_batch []ET.Hit)
func(a *Adapter) Scroll(ctx context.Context, index_name string, query *ET.Query, size int, cb ScrollCb) (err error) {

	if size < 100	{size = 100}
	if size > 10000	{size = 10000}

	if index_name = strings.TrimSpace(index_name); utf8.RuneCountInString(index_name) < 3 {
		err = errors.New("empty or too short index name in adapter.Scroll(): " + index_name)
		return
	}	//if err = indexnamey.Validate(); err != nil {return}

	if query == nil {
		err = errors.New("nil query")
		return
	}

	if cb == nil {
		err = errors.New("nil scroll callback")
		return
	}

	var client = a.Client

	var response *search.Response
	if response, err = client.Search().
		Index(index_name).
		Query(query).
		Size(size).
		Scroll(scrollDuration).
		Do(ctx); err != nil {return}

	if response.ScrollId_ == nil {
		err = errors.New("nil scroll id")
		return
	}

	var scroll_id string
	if scroll_id = *response.ScrollId_; scroll_id == "" {return errors.New("scroll_id is empty")}
	defer func() {
		var clear_resp *clearscroll.Response
		if clear_resp, err = client.ClearScroll().ScrollId(scroll_id).Do(ctx); err != nil {log.Error().Msgf("Failed to clear scroll %s", "err", err)}
		log.Debug().Msgf("(scroll) >> cleared %s %s %s %s %s %s", "index", index_name, "freed", clear_resp.NumFreed, "succ", clear_resp.Succeeded)
	}()

	if response.Hits.Total == nil {
		err = errors.New("nil meta hits info")
		return
	}

	var total_hits int
	if total_hits = int(response.Hits.Total.Value); total_hits == 0 {return}

	var fetched_hits int
	fetched_hits += len(response.Hits.Hits)

	cb(total_hits, fetched_hits, response.Hits.Hits)

	var scr = scroll.NewRequest()
	scr.Scroll		= "1m"
	scr.ScrollId	= scroll_id

	//var fullproof int
	for ; fetched_hits < total_hits ; {

		//if fullproof++; fullproof > 1000 {
		//	log.Warn("FULLPROOF TRIGGER")
		//	break
		//}

		var scr_resp *scroll.Response
		if scr_resp, err = client.Scroll().Request(scr).Do(ctx); err != nil {return}

		if len(scr_resp.Hits.Hits) == 0 {
			log.Warn().Msg("Zero hits arived, break")
			break
		}

		fetched_hits += len(scr_resp.Hits.Hits)
		cb(total_hits, fetched_hits, scr_resp.Hits.Hits)
	}

	return
}

