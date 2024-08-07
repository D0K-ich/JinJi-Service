package chat

import (
	"fmt"
	"time"
	"errors"
	"strings"

	"github.com/D0K-ich/types/uuid"
	"github.com/D0K-ich/JinJi-Service/store/dialogs"
)

func(h *Handler) NewMessage(message string, user_name string, dialog_uuid uuid.DialogUuid) (payload any, err error) {
	if !dialog_uuid.IsValid() {err = errors.New("invalid dialog uuid");return}
	if message 		= strings.TrimSpace(message); 	message 	== "" 	{err = errors.New("empty message"); return}
	if user_name 	= strings.TrimSpace(user_name); user_name 	== "" 	{err = errors.New("empty user_name"); return}

	var dialog *dialogs.Dialog
	if dialog, err = h.Mixins.Dialogs.Update(dialog_uuid.String(), dialogs.Messages{{
		UserName	: user_name,
		FromUser	: message,
		FromJinJi	: "FRROM DHIN",
		CreatedAt	: time.Now(),
	}}); err != nil {return}


	payload = map[string]interface{}{
		"date_time"	: time.Now(),
		"context" 	: fmt.Sprintf("Your message is %s. Dialog uuid: %s", message, dialog_uuid),
		"dialog"	: dialog,
	}
	return
}

func(h *Handler) NewDialog(user_name, user_message string) (payload any, err error) {
	if user_name 	= strings.TrimSpace(user_name); 	user_name == "" 	{err = errors.New("empty user name");return}
	if user_message = strings.TrimSpace(user_message); 	user_message == "" 	{err = errors.New("empty user name");return}

	//todo create request and responce to djinNeo
	var header = "Header"
	var subheader = "SUBHeader"

	var dialog *dialogs.Dialog
	if dialog, err = h.Mixins.Dialogs.New(header, subheader, user_name, &dialogs.Message{
		UserName:  user_name,
		FromUser:  user_message,
		FromJinJi: "FROM JINJUI", //todo
		CreatedAt: time.Now(),
	}); err != nil {return}

	payload = map[string]interface{}{
		"dialog" : dialog,
	}

	return
}

func(h *Handler) DeleteDialog(dialog_uuid uuid.DialogUuid) (payload any, err error) {
	if !dialog_uuid.IsValid() {err = errors.New("empty dialog_uuid")}

	if err = h.Dialogs.Delete(dialog_uuid.String()); err != nil {return}
	return
}