package users

import "errors"

func(h *Handler) Route(path string) (payload any, err error) {
	switch path {
	default				: err = errors.New("unknown path")
	case "create/new" 	: payload, err = h.NewUser()
	}

	return
}
