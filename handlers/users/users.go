package users

import (
	"github.com/D0K-ich/JinJi-Service/store/models"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"time"

	"github.com/D0K-ich/JinJi-Service/logs"
)

var log = logs.NewLog()

func (h *Handler) NewUser() (response *fasthttp.Response, err error) {
	var user = &models.User{ //TODO FOR TESTS!
		PrimaryId:        models.PrimaryId{},
		Uuid:             uuid.New(),
		//Avatar:           nil,
		Name:             "qew",
		Email:            "sda",
		State:            "as",
		Phone:            "das",
		Level:            10,
		TariffId:         20,
		Balance:          30,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		LastOnline:       time.Now(),
		TariffExpiration: nil,
		Friends:          &models.Friends{Friends: []*models.Friend{{Name: "fname"}}},
	}

	if err = h.Store().Users.New(user); err != nil {return}
	log.Info("New user created")
	return
}
