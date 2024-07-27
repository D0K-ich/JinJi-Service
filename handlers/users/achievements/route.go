package achievements

import (
	"github.com/D0K-ich/JinJi-Service/network/messages"
)

func(h *Handler) Route(message *messages.Message) (payload any, err error) {

	switch message.SubjectAction() {
	case "get/all"				: payload, err 	= h.GetAllAchievements()
	case "set/new"				: payload, err 	= h.NewAchievement(message.String("arch_name"))
	}

	return
}
