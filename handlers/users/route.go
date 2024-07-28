package users

import (
	"errors"
	"github.com/D0K-ich/JinJi-Service/network/messages"
	"github.com/D0K-ich/JinJi-Service/store/models"
)

func(h *Handler) Route(message *messages.Message) (payload any, err error) {
	switch message.SubjectAction() {
	default				: err = errors.New("unknown path")

	//profile
	case "profile/new"				: payload, err 	= h.Profile().NewUser(message.String("name"), message.String("password"), message.String("email"), h.UserId())
	case "profile/get"			    : payload, err 	= h.Profile().GetById(h.UserId())
	case "profile/update" 			: err 			= h.Profile().UpdateUser(messages.ReMarshalMust[*models.User](message.ToMap()))

	//billing
	//case "transactions/list"        : payload, err 	= h.Transactions().List(h.UserId())
	//case "transaction/status"       : payload, err 	= h.Transactions().StatusByOrderId(h.UserId(), incoming.String("order_id"))
	//case "transaction/create-ecom"  : payload, err 	= h.Transactions().CreateEcom(h.UserId(), incoming.Float("amount"), incoming.String("comment"))
	//
	//case "tariffs/list"			    : payload, err 	= h.Tariffs().List()
	//case "tariff/change"		    : payload, err 	= h.Tariffs().Change(h.UserId(), incoming.Int("tariff_id"))


	//friends
	case "add/new"				: payload, err = h.Friends().New(message.String("friend_name"), message.String("user_name"))
	case "drop/by-name"			: payload, err = h.Friends().Drop(message.String("friend_name"), message.String("user_name"))

	case "list/all"				: payload, err = h.Friends().GetAllFriends(message.String("user_name"))

	//settings
	//case "settings"				: payload, err = h.Settings().Route(message)

	//achievements
	case "achievements/all"				: payload, err 	= h.Achievements().GetAllAchievements()
	case "achievements/new"				: payload, err 	= h.Achievements().NewAchievement(message.String("arch_name"), message.String("user_name"))
	}

	return
}
