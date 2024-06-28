package users

import "github.com/D0K-ich/KanopyService/handlers/users/mixins"

func NewHandler(user_id int) (handler *Handler) {
	handler = &Handler{
		Mixins: mixins.NewMixins(user_id),
	}
	return
}

type Handler struct {
	*mixins.Mixins
}
