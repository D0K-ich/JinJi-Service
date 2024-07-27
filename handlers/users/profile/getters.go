package profile

import (
	"errors"
	"strings"

	"github.com/D0K-ich/JinJi-Service/store/models"
)

func(h *Handler) GetByName(name string) (payload any, err error) {
	if name = strings.TrimSpace(name); name == "" {err = errors.New("empty name for get"); return}

	var user *models.User
	if user, err = h.Mixins.Store.GetByName(name); err != nil {return}

	payload = map[string]interface{}{
		"user" : user,
	}
	return
}
