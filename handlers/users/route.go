package users

import (
	"errors"
	"github.com/D0K-ich/JinJi-Service/network/messages"
)

func(h *Handler) Route(message *messages.Message) (payload any, err error) {
	switch message.Module() {
	default				: err = errors.New("unknown path")

	case "profile"					: payload, err = h.Profile().Route(message)
	case "billing"					: payload, err = h.Transactions().Route(message)
	case "friends"					: payload, err = h.Friends().Route(message)
	case "settings"					: payload, err = h.Settings().Route(message)
	case "achievements"				: payload, err = h.Achievements().Route(message)
	}

	return
}
