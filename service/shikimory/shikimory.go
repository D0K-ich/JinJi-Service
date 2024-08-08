package shikimory
//TODO
type Shikimory struct {

}

func NewShiki(config *Config) (shiki *Shikimory, err error) {
	if err = config.Validate(); err != nil {return}

	shiki = &Shikimory{}

	return
}
