package users

import (
	"github.com/kr/pretty"
	"github.com/valyala/fasthttp"

	"github.com/D0K-ich/JinJi-Service/logs"
)

var log = logs.NewLog()

func (h *Handler) NewUser() (response *fasthttp.Response, err error) {
	if err = h.Store().Users.CreateUser("email", "phone", "name", 10); err != nil {
		return
	}

	pretty.Println(h.Store().Users.GetByEmail("email"))
	return
}
