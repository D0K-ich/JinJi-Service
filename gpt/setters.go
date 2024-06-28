package gpt

import (
	"io"
	"net/http"
)

func(d *DocGPT) newMessage(message, mode string) (answer string, err error) {
	var request *http.Request
	if request, err = d.newRequest(&Message{
		Message		: message,
		Mode		: mode,
	}, newMessage, "POST"); err != nil {return}

	var response *http.Response
	if response, err = d.client.Do(request); err != nil {return}

	var str_resp []byte
	if str_resp, err = io.ReadAll(response.Body); err != nil {return}

	answer = string(str_resp)
	return
}
