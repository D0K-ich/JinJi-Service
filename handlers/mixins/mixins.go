package mixins

import "github.com/D0K-ich/KanopyService/store"

func NewMixins() *Mixins {
	return &Mixins{
		store		: store.Default,
	}
}

type Mixins struct {
	store		*store.Store
}

func(m *Mixins) Store()		*store.Store		{return m.store}
