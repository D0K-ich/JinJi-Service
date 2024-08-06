package dialogs

import (
	"time"

	"github.com/D0K-ich/JinJi-Service/store/dialogs"
)

type ShortDialogs []*ShortDialog
type ShortDialog struct {
	Header 		string		`json:"header"`
	SubHeader 	string		`json:"sub_header"`
	UpdatedAt 	time.Time	`json:"updated_at"`
}

func NewShortDialogs(full_dialogs dialogs.Dialogs) (short_dialogs ShortDialogs, err error) {
	for _, dialog := range full_dialogs {
		short_dialogs = append(short_dialogs, &ShortDialog{
			Header		: dialog.Header,
			SubHeader	: dialog.SubHeader,
			UpdatedAt	: dialog.UpdatedAt,
		})
	}
	return
}
