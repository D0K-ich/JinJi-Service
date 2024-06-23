package gpt

import (
	"bytes"
	"encoding/json"
	"github.com/kr/pretty"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func(d *DocGPT) run() {
	var err error
	for {
		select {
		case <- d.context.Done() : return

		case message := <- d.message : if err = d.newMessage(message); err != nil {log.Error("Get err on gpt communicate", zap.Any("err", err))}
		}
	}
}

func(d *DocGPT) newMessage(message string) (err error) {
	var new_gpt_message = RequestGPT{
		Model:       d.config.Model,
		Temperature: d.config.Temperature,
		Messages:    []*MessageGPT{{
			Role:    "user",
			UserId:  "123",
			Content: message,
		}},
	}
	new_gpt_message.Messages = append(new_gpt_message.Messages, d.config.PersonRole...)

	var request []byte
	if request, err = json.Marshal(new_gpt_message); err != nil {return}

	var response *http.Response
	if response, err = d.client.Post(d.config.serverAddress(), defaultContentType, bytes.NewReader(request)); err != nil {return}
	defer response.Body.Close()

	var response_bytes []byte
	if response_bytes, err = io.ReadAll(response.Body); err != nil {return}

	var answer_gpt *ResponseGPT
	if err = json.Unmarshal(response_bytes, &answer_gpt); err != nil {return}

	pretty.Println(answer_gpt)

	return
}