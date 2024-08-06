package dialogs

import (
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/get"

	"github.com/D0K-ich/types/uuid"

	ET "github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func(s *Storage) GetAllByUserName(user_name string) (dialogs Dialogs, err error) {
	var query = &ET.Query{
		Bool: &ET.BoolQuery{
			Must: []ET.Query{
				{Terms: &ET.TermsQuery{TermsQuery: map[string]ET.TermsQueryField{"user_name" : user_name}}},
			},
		},
	}


	if err = s.elastic.Scroll(s.context, s.indexDialogs, query, 1000, func(total, fetched int, hits_batch []ET.Hit) {
		for _, hit := range hits_batch {
			var dialog *Dialog
			if err = json.Unmarshal(hit.Source_, &dialog); err != nil {return}
			dialogs = append(dialogs, dialog)
		}	//if prg_cb != nil {prg_cb(total, fetched, "(pages) >> Читаем данные...")}
	}); err != nil {return}

	return
}

func(s *Storage) GetById(dialog_id uuid.DialogUuid) (dialog *Dialog, err error) {
	if !dialog_id.IsValid() {err = errors.New("invalid user id");return}

	var response *get.Response
	if response, err = s.elastic.Client.Get(s.indexDialogs, dialog_id.String()).Do(s.context); err != nil {return}
	if err = json.Unmarshal(response.Source_, &dialog); err != nil {return}
	return

}