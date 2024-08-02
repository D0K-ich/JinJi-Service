package chat

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func(h *Handler) NewMessage(message, dialog_uuid string) (payload any, err error) {
	if message = strings.TrimSpace(message); message == "" {err = errors.New("empty message"); return}

	payload = map[string]interface{}{
		"date_time": time.Now(),
		"context" : fmt.Sprintf("Your message is %s. Dialog uuid: %s", message, dialog_uuid),
	}
	return
}

func(h *Handler) NewDialog() (payload any, err error) {

	return
}