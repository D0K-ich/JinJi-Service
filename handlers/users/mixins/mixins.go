package mixins

import "github.com/D0K-ich/KanopyService/handlers/mixins"

func NewMixins(user_id int) *Mixins {
	return &Mixins{
		userId: user_id,
		Mixins: mixins.NewMixins(),
	}
}

type Mixins struct {
	userId int
	*mixins.Mixins
}

func (m *Mixins) UserId() int           { return m.userId }
func (m *Mixins) SetUserId(user_id int) { m.userId = user_id }
