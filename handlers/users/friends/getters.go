package friends

import "github.com/D0K-ich/JinJi-Service/store/models"

func(h *Handler) GetAllFriends(name string) (payload any, err error) {
	var user *models.User
	if user, err = h.Store.GetByName(name); err != nil {return}

	payload = map[string]interface{}{
		"friends" : user.Friends.Friends,
	}

	return
}
