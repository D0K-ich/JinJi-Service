package users

import (
	"github.com/D0K-ich/JinJi-Service/handlers/users/achievements"
	"github.com/D0K-ich/JinJi-Service/handlers/users/friends"
	"github.com/D0K-ich/JinJi-Service/handlers/users/mixins"
	"github.com/D0K-ich/JinJi-Service/handlers/users/profile"
	"github.com/D0K-ich/JinJi-Service/handlers/users/settings"
	"github.com/D0K-ich/JinJi-Service/handlers/users/tariffs"
	"github.com/D0K-ich/JinJi-Service/handlers/users/transactions"
)

func NewHandler(user_id int) (handler *Handler) {
	handler = &Handler{
		Mixins: mixins.NewMixins(user_id),
	}
	return
}

type Handler struct {
	*mixins.Mixins
}

func(h *Handler) Profile() 		    *profile.Handler           	{return profile.NewHandler(h.Mixins)}
func(h *Handler) Tariffs() 		    *tariffs.Handler           	{return tariffs.NewHandler(h.Mixins)}
func(h *Handler) Settings()         *settings.Handler           {return settings.NewHandler(h.Mixins)}
func(h *Handler) Friends()          *friends.Handler            {return friends.NewHandler(h.Mixins)}
func(h *Handler) Achievements()     *achievements.Handler       {return achievements.NewHandler(h.Mixins)}
func(h *Handler) Transactions() 	*transactions.Handler 		{return transactions.NewHandler(h.Mixins)}