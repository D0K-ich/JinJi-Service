package messages

import (
	"github.com/google/uuid"
	"strings"
	"sync"
)

const (
	FieldType       = "_type"			// messagfe type request | response | error | event
	FieldKind       = "_kind"
	FieldAccessToken= "_access_token"

	// uuids
	FieldUuid       = "_uuid"
	FieldUserUuid	= "_user_uuid"
	FieldAdminUuid	= "_admin_uuid"
	FieldLinkUuid   = "_link_uuid"
	FieldClientUuid = "_client_uuid"
	FieldWorkerUuid = "_worker_uuid"
	FieldTaskUuid   = "_task_uuid"

	// paths
	FieldModule 	= "_module"
	FieldSubject    = "_subject"
	FieldAction		= "_action"

	// some meta info
	FieldIpAddr		= "_ip_addr"
	FieldAppVer		= "_app_ver"
	FieldTimeout	= "_timeout"		// in seconds
	FieldStatusCode	= "_status_code"

	// data
	FieldBytes		= "_bytes"

	FieldMessage	= "message"
	//FieldAttachment	= "_attachment"
)

const (
	TypeRequest		= Type("request")	// outgoing message
	TypeResponse	= Type("response")	// response of request
	TypeError		= Type("error")		// error response
	TypeEvent		= Type("event")		// event - usually async message between systems ()
	TypeConfirm		= Type("confirm")	// confirm delivery responsibility transfer

	KindJson		= "json"
	KindBinary		= "binary"
	// for notice or pinf use event with Module = notice or ping
	//TypeNotice	= Type("notice")	//
	//TypePing		= Type("ping")
)

type Type string
func(t Type) String() string {return string(t)}

type Message struct {
	lock			sync.RWMutex

	_type			Type
	_kind			string
	_accessToken	string

	_uuid			string
	_userUuid		string
	//_adminUuid		uuid.AdminUuid
	//_linkUuid		uuid.LinkUuid
	//_clientUuid		uuid.ClientUuid
	_workerUuid		uuid.UUID
	//_taskUuid		uuid.TaskUuid

	_module			Module
	_subject		Subject
	_action			Action

	_ipAddr			string
	_appVer			string
	_timeout		int					// in seconds
	_statusCode		int
	_bytes			[]byte

	message			string
	payload			map[string]any
}

type Module string
func(m Module) String() string {return string(m)}
func(m Module) IsEmpty() bool {return strings.TrimSpace(string(m)) == ""}

type Subject string
func(s Subject) String() string {return string(s)}
func(s Subject) IsEmpty() bool {return strings.TrimSpace(string(s)) == ""}

type Action string
func(a Action) String() string {return string(a)}
func(a Action) IsEmpty() bool {return strings.TrimSpace(string(a)) == ""}