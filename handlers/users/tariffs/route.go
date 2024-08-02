package tariffs

import (
	"github.com/D0K-ich/types/message"
)

func(h *Handler) Route(message *message.Message) (payload any, err error) {

	switch message.SubjectAction() {
	//case "profile/new"				: payload, err 	= h.NewUser(message.String("name"), message.String("password"), message.String("email"))
	//case "profile/get"			    : payload, err 	= h.GetByName(message.String("name"))
	//case "profile/update" 			: err 			= h.UpdateUser(messages.ReMarshalMust[*models.User](message.ToMap()))

	}

	return
}
