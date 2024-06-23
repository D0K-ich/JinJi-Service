package admins

import "github.com/D0K-ich/KanopyService/handlers/mixins"

func NewHandler() (handler *Handler) {
	handler = &Handler{
		Mixins: mixins.NewMixins(),
	}
	return
}

type Handler struct {
	*mixins.Mixins
}

