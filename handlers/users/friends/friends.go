package friends

import (
	"github.com/D0K-ich/JinJi-Service/handlers/users/mixins"
	"github.com/D0K-ich/JinJi-Service/logs"
)

var log = logs.NewLog()

func NewHandler(mix *mixins.Mixins) (handler *Handler) {
	handler = &Handler{
		Mixins: mix,
	}
	return
}

type Handler struct {
	*mixins.Mixins
}
