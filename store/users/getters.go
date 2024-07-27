package users

import (
	"errors"
	"strings"

	"github.com/D0K-ich/JinJi-Service/store/models"
)

func(s *Storage) GetByName(name string) (user *models.User, err error) {
	if name = strings.TrimSpace(name); name == "" {err = errors.New("get nil user name for get"); return}

	if err = s.db.Where("name = ?", name).First(&user).Error;  err != nil {return}
	return
}
