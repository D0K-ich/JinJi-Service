package users

import (
	"errors"
	"github.com/D0K-ich/JinJi-Service/network/messages"
)

func(h *Handler) Route(message *messages.Message) (payload any, err error) {
	switch message.SubjectAction() {
	default				: err = errors.New("unknown path")

	case "create/new"	: payload, err = h.newUser(message.String("name"), message.String("password"), message.String("email"))
	case "get/by-name"	: payload, err = h.getByName(message.String("name"))
	}

	return
}
