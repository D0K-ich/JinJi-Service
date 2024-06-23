package mixins

import "github.com/D0K-ich/KanopyService/store"

func NewMixins() (mixin *Mixins) {
	return &Mixins{
		Store : store.Default,
	}
}

type Mixins struct {
	Store *store.Store
}