package users

import (
	"errors"
	"github.com/D0K-ich/JinJi-Service/store/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"strings"
	"time"
)

func (h *Handler) newUser(name, password, email string) (payload any, err error) {
	if name 	= strings.TrimSpace(name); 		name 		== "" {err = errors.New("empty name for create user");return}
	if email 	= strings.TrimSpace(email); 	email 		== "" {err = errors.New("empty email for create user");return}
	if password = strings.TrimSpace(password); 	password 	== "" {err = errors.New("empty password for create user");return}

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
		Friends			: nil,
	}

	if err = h.Store().Users.Save(new_user); err != nil {return}
	log.Info("New user created", zap.Any("name", name))
	return
}

