package users

import (
	"github.com/D0K-ich/JinJi-Service/store/models"
)

func(s *Storage) New(user *models.User) (err error) {
	if err = s.db.Create(user).Error; err != nil {return}
	return
}

func(s *Storage) Ban(user_name string) (err error) {
	if err = s.db.Where("name = ?", user_name).Delete(&models.User{}).Error; err != nil {return}
	return
}