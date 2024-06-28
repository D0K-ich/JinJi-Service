package gpt

func(d *DocGPT) Route(request any, path string) (payload any, err error) {
	switch path {
	case "message/new": payload, err = d.newMessage(request.(string), "chat")
	}

	return
}
