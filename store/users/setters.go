package users

import (
	"github.com/D0K-ich/JinJi-Service/store/models"
)

func(s *Storage) Save(user *models.User) (err error) {
	if err = s.db.Save(&user).Error; err != nil {return}
	return
}

func(s *Storage) Ban(user_name string) (err error) {
	var user *models.User
	if err = s.db.Where("name = ?", user_name).First(&user).Error; err != nil {return}

	user.State = models.StateBlocked

	if err = s.db.Save(&user).Error; err != nil {return}
	return
}