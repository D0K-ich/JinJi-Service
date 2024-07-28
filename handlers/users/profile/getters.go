package profile

import (

	"github.com/D0K-ich/JinJi-Service/store/models"
)

func(h *Handler) GetById(id int) (payload any, err error) {
	//if id = strings.TrimSpace(id); id == "" {err = errors.New("empty name for get"); return}

	var user *models.User
	if user, err = h.Mixins.Store.GetByID(id); err != nil {return}

	payload = map[string]interface{}{
		"user" : user,
	}
	return
}
