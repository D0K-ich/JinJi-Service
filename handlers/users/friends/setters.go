package friends

import (
	"errors"
	"strings"

	"github.com/rs/zerolog/log"
)

func(h *Handler) New(friend_name, user_name string) (payload any, err error) {
	if friend_name = strings.TrimSpace(friend_name); friend_name == "" {err = errors.New("empty name for create user");return}
	if user_name = strings.TrimSpace(user_name); user_name == "" {err = errors.New("empty name main user");return}

	if user_name == friend_name {err = errors.New("this your nick;)"); return}

	if err = h.Mixins.Store.AddFriend(user_name, friend_name); err != nil {err = errors.New("err on add frient user " + err.Error());return}

	log.Info().Msg("New friend added")
	return
}

func(h *Handler) Drop(friend_name, user_name string) (payload any, err error) {
	if friend_name = strings.TrimSpace(friend_name); friend_name == "" {err = errors.New("empty name for create user");return}
	if user_name = strings.TrimSpace(user_name); user_name == "" {err = errors.New("empty name main user");return}

	if user_name == friend_name {err = errors.New("this your nick;)"); return}

	if err = h.Mixins.Store.DropFriend(user_name, friend_name); err != nil {err = errors.New("err on add frient user " + err.Error());return}

	log.Info().Msg("New friend droppped")
	return
}