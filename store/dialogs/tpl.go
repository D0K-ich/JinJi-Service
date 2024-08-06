package dialogs

import (
	"time"

	"github.com/D0K-ich/types/uuid"
	"github.com/D0K-ich/JinJi-Service/store/adapters/elastic"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type Dialogs []*Dialog
type Dialog struct {
	Uuid 			uuid.DialogUuid 	`json:"uuid"`

	UserName 		string				`json:"user_name"`
	Header 			string				`json:"header"`
	SubHeader 		string				`json:"sub_header"`
	JinJiVersion 	float64				`json:"jinji_version"`
	CreatedAt 		time.Time 			`json:"created_at"`
	UpdatedAt 		time.Time 			`json:"updated_at"`

	Messages 		Messages			`json:"messages"`
}

type Messages []*Message
type Message struct {
	FromUser 		string		`json:"from_user"`
	FromJinJi 		string		`json:"from_jinji"`

	CreatedAt 		time.Time 		`json:"created_at"`
}

var FieldDialogsMap = &types.TypeMapping{
	Properties: map[string]types.Property{
		"uuid"			: elastic.FieldSimpleKeyword,
		"user_name"		: elastic.FieldSimpleKeyword,
		"header"		: elastic.FieldSimpleKeyword,
		"sub_header"	: elastic.FieldSimpleKeyword,
		"created_at"	: elastic.FieldDateMap,
		"updated_at"	: elastic.FieldDateMap,
		"jinji_version" : elastic.FieldFloat64Map,

		"messages"		: FieldMessageMap,
	},
}

var FieldMessageMap = &types.TypeMapping{
	Properties: map[string]types.Property{
		"from_user"		: elastic.FieldTextMap,
		"from_jinji"	: elastic.FieldTextMap,
		"created_at"	: elastic.FieldDateMap,
	},
}