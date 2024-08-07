package achievements

import (
	"errors"
	"strings"
)

func(h *Handler) NewAchievement(arch_name []string, user_name string) (payload any, err error) {
	if user_name = strings.TrimSpace(user_name); user_name == "" {err = errors.New("empty user name");return}
	if len(arch_name) == 0 {err = errors.New("empty arch name");return}

	if err = h.Users.AddArch(arch_name, user_name); err != nil {return}
	return
}