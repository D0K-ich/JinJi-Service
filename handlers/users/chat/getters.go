package chat

import (
	"errors"
	"strings"

	"github.com/D0K-ich/JinJi-Service/store/models"
)

func(h *Handler) GetAllDialogs(user_name string) (payload any, err error) {
	var user *models.User
	if user, err = h.Users.GetByName(user_name); err != nil {return}	//todo

	payload = map[string]interface{}{
		"dialogs" : user.Achievements,
	}

	return
}

func(h *Handler) GetAllMessages(dialiog_name, user_id string) (payload any, err error) {
	if dialiog_name = strings.TrimSpace(dialiog_name); dialiog_name == "" {err = errors.New("empty dialog name"); return}
	if user_id = strings.TrimSpace(user_id); user_id == "" {err = errors.New("empty dialog name"); return}


	return
}