package service

import "github.com/D0K-ich/JinJi-Service/service/shikimory"

type Config struct {
	Shikimory *shikimory.Config	`yaml:"shikimory"`
}
