package gpt

import (
	"context"
	"net/http"
	"time"

	"github.com/D0K-ich/KanopyService/logs"
)

type DocGPT struct {
	context context.Context
	cancel 	context.CancelFunc

	client *http.Client
	user 	any		//todo

	message chan string
	response chan string

	store any				//todo
	config *Config
}

var log = logs.NewLog()

func NewDefaultGPT(ctx context.Context, user any) (gpt *DocGPT, err error) {
	log.Info("Start create default gpt...")

	gpt = &DocGPT{
		client	: &http.Client{Timeout: defaultTimeout * time.Minute},
		message	: make(chan string),
		response: make(chan string),
		store	: nil,
		config	: defaultConfig(),
		user	: user,
	}
	gpt.context, gpt.cancel = context.WithCancel(ctx)

	go gpt.run()

	gpt.newMessage("Hello")

	log.Info("Finish create default gpt")
	return
}

func NewGptByConfig(config *Config, user any) (gpt *DocGPT, err error) {
	log.Info("Start create custom gpt...")
	if err = config.Validate(); err != nil {return}

	gpt = &DocGPT{
		client	: &http.Client{Timeout: config.TimeoutResponseMin * time.Minute},
		message	: make(chan string),
		response: make(chan string),
		store	: nil,
		config	: config,
		user	: user,
	}

	go gpt.run()

	log.Info("Finish create custom gpt")
	return
}

