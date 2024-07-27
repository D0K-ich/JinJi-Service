package users

import (
	"errors"
	"strings"

	"github.com/D0K-ich/JinJi-Service/store/models"
)

func(h *Handler) getByName(name string) (payload any, err error) {
	if name = strings.TrimSpace(name); name == "" {err = errors.New("empty name for get"); return}

	var user *models.User
	if user, err = h.Store().Users.GetByName(name); err != nil {return}

	payload = map[string]interface{}{
		"user" : user,
	}
	return
}
