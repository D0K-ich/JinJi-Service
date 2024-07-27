package friends

import (
	"github.com/D0K-ich/JinJi-Service/network/messages"
)

func(h *Handler) Route(message *messages.Message) (payload any, err error) {

	switch message.SubjectAction() {
	case "add/new"				: payload, err = h.New(message.String("friend_name"), message.String("user_name"))
	case "drop/by-name"			: payload, err = h.Drop(message.String("friend_name"), message.String("user_name"))

	case "list/all"				: payload, err = h.GetAllFriends(message.String("user_name"))
	}

	return
}
