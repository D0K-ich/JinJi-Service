package gpt

import (
	"fmt"
	"time"
	"bytes"
	"context"
	"net/http"
	"encoding/json"

	"github.com/kr/pretty"
	"github.com/rs/zerolog/log"
)

type DocGPT struct {
	context context.Context
	cancel  context.CancelFunc

	client *http.Client
	user   any //todo

	message  chan string
	response chan string

	store  any //todo
	config *Config
}

func NewDefaultGPT(ctx context.Context, user any) (gpt *DocGPT, err error) {
	log.Info().Msg("Start create default gpt...")

	gpt = &DocGPT{
		client:   &http.Client{Timeout: defaultTimeout * time.Minute},
		message:  make(chan string),
		response: make(chan string),
		store:    nil,
		config:   defaultConfig(),
		user:     user,
	}
	gpt.context, gpt.cancel = context.WithCancel(ctx)

	go gpt.run()

	var answer string
	if answer, err = gpt.newMessage("Про что жанр иссекай?", "chat"); err != nil {return}

	log.Info().Msgf("Finish create default gpt %s", answer)
	return
}

func NewGptByConfig(config *Config, user any) (gpt *DocGPT, err error) {
	log.Info().Msg("Start create custom gpt...")
	if err = config.Validate(); err != nil {
		return
	}

	gpt = &DocGPT{
		client:   &http.Client{Timeout: config.TimeoutResponseMin * time.Minute},
		message:  make(chan string),
		response: make(chan string),
		store:    nil,
		config:   config,
		user:     user,
	}

	go gpt.run()

	log.Info().Msg("Finish create custom gpt")
	return
}

func (d *DocGPT) run() {
	var err error
	var answer string

	for {
		select {
		case <-d.context.Done():
			return

		case message := <-d.message:
			if answer, err = d.newMessage(message, "chat"); err != nil {
				log.Error().Msgf("Get err on gpt communicate", "err", err)
			}
			pretty.Println(answer)
		}
	}
}

func (d *DocGPT) newRequest(message *Message, endpoint, method string) (request *http.Request, err error) {
	var request_byte []byte
	if request_byte, err = json.Marshal(message); err != nil {
		return
	}

	if request, err = http.NewRequest(method, fmt.Sprintf("%s%s", d.config.RootPath, endpoint), bytes.NewReader(request_byte)); err != nil {
		return
	}
	request.Header.Add("Authorization", d.config.Token)
	request.Header.Add("Content-Type", defaultContentType)

	return
}
