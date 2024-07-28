package users

import (
	"time"

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

//arch
func(s *Storage) AddArch(arch_name, user_name string) (err error) {
	var user *models.User
	if err = s.db.Where("name = ?", user_name).First(&user).Error; err != nil {return}

	user.Achievements.Achievements = append(user.Achievements.Achievements, &models.Achievement{
		Name	:    arch_name,
		DateGet	: time.Now(),
	})

	if err = s.db.Save(&user).Error; err != nil {return}
	return
}

//Friends
func(s *Storage) AddFriend(user_name, friend_name string) (err error) {
	var user 	*models.User
	var friend 	*models.User
	if user, err 	= s.GetByName(user_name); err != nil {return}

	for _, exist_friend := range user.Friends.Friends {if exist_friend.Name == friend_name {return}}

	if friend, err 	= s.GetByName(friend_name); err != nil {return}

	if friend != nil && friend.Name != "" {
		user.Friends.Friends = append(user.Friends.Friends, &models.Friend{
			Name		: friend.Name,
			LastOnline	: friend.LastOnline,
			Level		: friend.Level,
		})
	}

	if err = s.db.Save(&user).Error; err != nil {return}
	return
}

func(s *Storage) DropFriend(user_name, friend_name string) (err error) {
	var user 	*models.User
	var friend 	*models.User
	if user, err 	= s.GetByName(user_name); err != nil {return}
	if friend, err 	= s.GetByName(friend_name); err != nil {return}

	if friend != nil && friend.Name != "" {
		var new_friends = &models.Friends{Friends: []*models.Friend{}}
		for _, exist_friend := range user.Friends.Friends {
			if exist_friend.Name == friend_name {continue}
			new_friends.Friends = append(new_friends.Friends, exist_friend)
		}
		user.Friends = new_friends
	}

	if err = s.db.Save(&user).Error; err != nil {return}
	return
}