package mixins

import (
	"github.com/D0K-ich/JinJi-Service/handlers/mixins"
	"github.com/D0K-ich/JinJi-Service/store/dialogs"
	"github.com/D0K-ich/JinJi-Service/store/users"
)

func NewMixins(user_id int) *Mixins {
	return &Mixins{
		userId: user_id,
		Users: mixins.NewMixins().Store().Users,
		Dialogs: mixins.NewMixins().Store().Dialogs,
	}
}

type Mixins struct {
	userId 		int
	Users  		*users.Storage
	Dialogs  	*dialogs.Storage
}

func (m *Mixins) UserId() int           { return m.userId }
func (m *Mixins) SetUserId(user_id int) { m.userId = user_id }
