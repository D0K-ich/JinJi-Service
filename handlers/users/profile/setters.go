package profile

import (
	"errors"
	"github.com/D0K-ich/JinJi-Service/store/models"
	"github.com/google/uuid"
	"github.com/kr/pretty"
	"go.uber.org/zap"
	"strings"
	"time"
)

func(h *Handler) NewUser(name, password, email string) (payload any, err error) {
	if name 	= strings.TrimSpace(name); 		name 		== "" {err = errors.New("empty name for create user");return}
	if email 	= strings.TrimSpace(email); 	email 		== "" {err = errors.New("empty email for create user");return}
	if password = strings.TrimSpace(password); 	password 	== "" {err = errors.New("empty password for create user");return}

	var exist_user *models.User
	if exist_user, err = h.Mixins.Store.GetByName(name); exist_user.Name != "" {
		err = errors.New("user with this nick already exist")
		pretty.Println(exist_user, err)
		return
	}

	var new_user = &models.User{
		Uuid			: uuid.New(),
		Name			: name,
		Email			: email,
		State			: models.StateUnconfirmed,
		Phone			: "",
		Password		: password,
		Level			: 0,
		TariffId		: 0,
		Balance			: 100,
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
		LastOnline		: time.Now(),
		TariffExpiration: nil,
		Friends			: &models.Friends{Friends: []*models.Friend{}},
	}

	if err = h.Mixins.Store.Save(new_user); err != nil {return}
	log.Info("New user created", zap.Any("name", name))
	return
}

func(h *Handler) UpdateUser(user *models.User) (err error) {
	if user == nil {return errors.New("nil user for update")}

	if err = h.Mixins.Store.Save(user); err != nil {return}
	log.Info("New user created", zap.Any("name", user))

	return
}
