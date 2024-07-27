package profile

import (
	"github.com/D0K-ich/JinJi-Service/network/messages"
	"github.com/D0K-ich/JinJi-Service/store/models"
)

func(h *Handler) Route(message *messages.Message) (payload any, err error) {

	switch message.SubjectAction() {
	case "profile/new"				: payload, err 	= h.NewUser(message.String("name"), message.String("password"), message.String("email"))
	case "profile/get"			    : payload, err 	= h.GetByName(message.String("name"))
	case "profile/update" 			: err 			= h.UpdateUser(messages.ReMarshalMust[*models.User](message.ToMap()))

	//case "profile/send-code"	    : err 			= h.SendCode(incoming.String("email"))
	//case "profile/sign-in"		    : payload, err 	= h.SignIn(incoming.String("email"), incoming.String("code"))
	//case "profile/sign-out"		    : payload, err 	= h.SignOut(h.UserId())
	//
	//case "profile/bill-out"		    : payload, err 	= h.BillOut(iface.ReMarshalMust[*manager.TaskCostMeta](incoming.ToMap()))
	}

	return
}
