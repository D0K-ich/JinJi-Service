package dialogs

import (
	"fmt"
	"time"
	"errors"
	"strings"
	"encoding/json"

	"github.com/D0K-ich/types/uuid"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/get"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

func(s *Storage) New(header, subheader, user_name string, user_message *Message) (dialog *Dialog, err error) {
	if header 		= strings.TrimSpace(header)		; header == "" 		{err = errors.New("empty header for new"); return}
	if subheader 	= strings.TrimSpace(subheader)	; subheader == "" 	{err = errors.New("empty subheader for new"); return}
	if user_name 	= strings.TrimSpace(user_name)	; user_name == "" 	{err = errors.New("empty user_name for new"); return}

	if user_message == nil {err = errors.New("nil user message");return}

	dialog = &Dialog{
		Uuid			: uuid.NewDialogUuid(),
		UserName		: user_name,
		Header			: header,
		SubHeader		: subheader,
		JinJiVersion	: 0,			//todo
		CreatedAt		: time.Now(),
		UpdatedAt		: time.Now(),
		Messages		: Messages{user_message},
	}

	if _, err = s.elastic.Client.Index(s.indexDialogs).
		Id(dialog.Uuid.String()).
		Request(dialog).
		Refresh(refresh.Waitfor).
		Do(s.context); err != nil {return}
	return
}

func(s *Storage) Update(dialog_uuid string, messages Messages) (dialog *Dialog, err error) {
	if dialog_uuid = strings.TrimSpace(dialog_uuid); dialog_uuid == "" {err = errors.New("empty pool_uuid in ProxyPoolByUuid");return}

	var response *get.Response
	if response, err = s.elastic.Client.Get(s.indexDialogs, dialog_uuid).Do(s.context); err != nil {return}
	if !response.Found {err = errors.New("dialog not found for uuid: " + dialog_uuid);return}

	if err = json.Unmarshal(response.Source_, &dialog); err != nil {return}

	dialog.Messages = append(dialog.Messages, messages...)

	if _, err = s.elastic.Client.Index(s.indexDialogs).
		Id(dialog.Uuid.String()).
		Request(dialog).
		Refresh(refresh.Waitfor).
		Do(s.context); err != nil {return}

	if len(messages) == 0 {err = fmt.Errorf("nil dialog_uuid by uuid %s", dialog_uuid);return}
	return
}

func(s *Storage) Delete(dialog_uuid string) (err error) {
	if _, err = s.elastic.Client.Delete(s.indexDialogs, dialog_uuid).Refresh(refresh.Waitfor).Do(s.context); err != nil {return}
	return
}