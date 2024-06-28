package gpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/kr/pretty"
	"go.uber.org/zap"

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

	var answer string
	if answer, err = gpt.newMessage("Про что жанр иссекай?", "chat"); err != nil {return}
	pretty.Println(answer)

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

func(d *DocGPT) run() {
	var err error
	var answer string

	for {
		select {
		case <- d.context.Done() : return

		case message := <- d.message :
			if answer, err = d.newMessage(message, "chat"); err != nil {log.Error("Get err on gpt communicate", zap.Any("err", err))}
			pretty.Println(answer)
		}
	}
}

func(d *DocGPT) newRequest(message *Message, endpoint, method string) (request *http.Request, err error) {
	var request_byte []byte
	if request_byte, err = json.Marshal(message); err != nil {return}

	if request, err = http.NewRequest(method, fmt.Sprintf("%s%s", d.config.RootPath, endpoint), bytes.NewReader(request_byte)); err != nil {return}
	request.Header.Add("Authorization", d.config.Token)
	request.Header.Add("Content-Type", defaultContentType)

	return
}

