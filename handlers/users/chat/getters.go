package chat

import (
	"errors"
	"github.com/D0K-ich/JinJi-Service/store/models"
	"strings"
)

func(h *Handler) GetAllDialogs(user_name string) (payload any, err error) {
	var user *models.User
	if user, err = h.Store.GetByName(user_name); err != nil {return}

	payload = map[string]interface{}{
		"dialogs" : user.Dialogs.Dialogs,
	}

	return
}

func(h *Handler) GetAllMessages(dialiog_name, user_id string) (payload any, err error) {
	if dialiog_name = strings.TrimSpace(dialiog_name); dialiog_name == "" {err = errors.New("empty dialog name"); return}
	if user_id = strings.TrimSpace(user_id); user_id == "" {err = errors.New("empty dialog name"); return}


	return
}