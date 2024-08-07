package profile

import (
	"time"
	"errors"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/D0K-ich/types/uuid"
	"github.com/D0K-ich/JinJi-Service/store/models"
)

func(h *Handler) NewUser(name, password, email string) (payload any, err error) {
	if name 	= strings.TrimSpace(name); 		name 		== "" {err = errors.New("empty name for create user");return}
	if email 	= strings.TrimSpace(email); 	email 		== "" {err = errors.New("empty email for create user");return}
	if password = strings.TrimSpace(password); 	password 	== "" {err = errors.New("empty password for create user");return}

	var exist_user *models.User
	if exist_user, err = h.Mixins.Users.GetByName(name); exist_user.Name != "" {
		err = errors.New("user with this nick already exist")
		return
	}

	var new_user = &models.User{
		Uuid			: uuid.NewUserUuid().String(),
		Name			: name,
		Email			: email,
		State			: models.StateUnconfirmed,
		Phone			: "",
		Password		: password,
		Level			: &models.Level{
			Name			: "Новичок",
			PointsCurrent	: 0,
			PointsTotal		: 100,
		},
		TariffId		: 0,
		Balance			: 100,
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
		LastOnline		: time.Now(),
		TariffExpiration: nil,
		Friends			: &models.Friends{Friends: []*models.Friend{}},
		Achievements	: &models.Achievements{Achievements: []*models.Achievement{}},
	}

	if err = new_user.NewAchievement(&models.Achievement{Name: "Привет!", DateGet: time.Now()}); err != nil {return}

	if err = h.Mixins.Users.Save(new_user); err != nil {return}
	log.Info().Msgf("New user created", "name", name)
	return
}

func(h *Handler) UpdateUser(user *models.User) (err error) {
	if user == nil {return errors.New("nil user for update")}

	if err = h.Mixins.Users.Save(user); err != nil {return}
	log.Info().Msgf("New user created", "name", user)

	return
}
