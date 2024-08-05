package chat

import (
	"github.com/D0K-ich/JinJi-Service/handlers/users/mixins"
)

func NewHandler(mix *mixins.Mixins) (handler *Handler) {
	handler = &Handler{
		Mixins: mix,
	}
	return
}

type Handler struct {
	*mixins.Mixins
}
