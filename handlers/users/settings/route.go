package settings

import (
	"github.com/D0K-ich/JinJi-Service/network/messages"
)

func(h *Handler) Route(message *messages.Message) (payload any, err error) {

	switch message.SubjectAction() {


	}

	return
}
