package achievements

import (
	"errors"
	"strings"
)

func(h *Handler) NewAchievement(arch_name, user_name string) (payload any, err error) {
	if user_name = strings.TrimSpace(user_name); user_name == "" {err = errors.New("empty user name");return}
	if arch_name = strings.TrimSpace(arch_name); arch_name == "" {err = errors.New("empty arch name");return}

	if err = h.Store.AddArch(arch_name, user_name); err != nil {return}
	return
}