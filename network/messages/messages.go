package messages

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/D0K-ich/JinJi-Service/logs"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"unicode/utf8"
)

var log = logs.NewLog()

func FromJson(data json.RawMessage) (m *Message, err error) {
	var mmap = make(map[string]any)
	if err = json.Unmarshal(data, &mmap); err != nil {return}
	m = map2message(mmap)
	return
}

func FromMap(mmap map[string]any) *Message {
	return map2message(mmap)
}

func map2message(in map[string]any) (m *Message) {

	m = newMessage()

	for key, val := range in {
		switch key {
		case FieldUuid			: m._uuid			= ToString(val)
		case FieldType			: m._type			= Type(ToString(val))
		case FieldKind			: m._kind			= ToString(val)
		case FieldAccessToken	: m._accessToken	= ToString(val)
		case FieldUserUuid		: m._userUuid		= ToString(val)
		//case FieldAdminUuid		: m._adminUuid		= uuid.AdminUuid(iface.ToString(val))
		//case FieldLinkUuid		: m._linkUuid		= uuid.LinkUuid(iface.ToString(val))
		//case FieldClientUuid	: m._clientUuid		= uuid.ClientUuid(iface.ToString(val))
		//case FieldWorkerUuid	: m._workerUuid		= uuid.WorkerUuid(iface.ToString(val))
		//case FieldTaskUuid		: m._taskUuid		= uuid.TaskUuid(iface.ToString(val))
		case FieldModule		: m._module			= Module(ToString(val))
		case FieldSubject		: m._subject		= Subject(ToString(val))
		case FieldAction		: m._action			= Action(ToString(val))
		case FieldIpAddr		: m._ipAddr			= ToString(val)
		case FieldAppVer		: m._appVer			= ToString(val)
		case FieldTimeout		: m._timeout		= ToInt(val)
		case FieldStatusCode	: m._statusCode		= ToInt(val)
		//case FieldBytes			: m._
		case FieldMessage		: m.message			= strings.TrimSpace(ToString(val))
		default					: m.payload[key]	= val
		}
	}

	return
}

func ToString(input interface{}) string {
	switch value := input.(type) {
	case nil								: return ""
	case interface{String() string}			: return value.String()
	case string								: return value
	case []string							: return strings.Join(value, "; ")
	case int, int32, int64, float32, float64: return fmt.Sprintf("%v", input)
	case bool	: if value {return "true"} else {return "false"}
	}
	return ""
}

func newMessage() *Message {
	return &Message{
		payload	: make(map[string]any),
	}
}

func ToStringSlice(input interface{}) (output []string) {
	switch values := input.(type) {
	case nil			: return
	case bool			: if values {return []string{"true"}} else {return []string{"false"}}
	case string			: if strings.TrimSpace(values) == "" {return} else {return []string{values}}
	case []string		:
		for _, v := range values {
			if v == "" {continue}
			output = append(output, v)
		}
	case []interface{String() string}	:
		for _, v := range values {
			if v.String() == "" {continue}
			output = append(output, v.String())
		}
	case []interface{}	:
		for _, v := range values {
			if ToString(v) == "" {continue}
			output = append(output, ToString(v))
		}
	}
	return
}


func ToInt(in interface{}) (out int) {
	if in == nil {return}
	switch value := in.(type) {
	case int		: return value
	case int64		: return int(value)
	case int32		: return int(value)
	case float64	: return int(value)
	case float32	: return int(value)
	case string	:
		if value = strings.TrimSpace(value); value == "" {return}
		var err error
		if out, err = strconv.Atoi(value); err != nil {
			return
		}
	}
	return
}

func ToInt64(in interface{}) (out int64) {
	if in == nil {return}
	switch value := in.(type) {
	case int		: return int64(value)
	case int64		: return value
	case int32		: return int64(value)
	case float64	: return int64(value)
	case float32	: return int64(value)
	case string	:
		if value = strings.TrimSpace(value); value == "" {return}
		var err error
		if out, err = strconv.ParseInt(value, 10, 64); err != nil {
			return
		}
	}
	return
}

func ToIntSlice(input interface{}) (out []int) {
	switch values := input.(type) {
	case nil		: return
	case bool		: return
	case []string	:
		for _, item := range values {
			if item = strings.TrimSpace(item); item == "" {continue}
			var val, _ = strconv.Atoi(item)
			out = append(out, val)
		}
	case []interface{} :
		for _, item := range values {
			var val = ToInt(item)
			if val == 0 {continue}
			out = append(out, val)
		}
	}

	return
}

func ToBool(input interface{}) bool {
	switch val := input.(type) {
	case bool	: return val
	case string	: return strings.TrimSpace(strings.ToLower(val)) == "true"
	case nil	: return false
	default		: log.Warn("unwxpected type in ToBool", zap.Any("input", input))
	}
	return false
}

func ToFloat(in any) float64 {
	if in == nil {return 0}
	switch val := in.(type) {
	default			: log.Error("unexpected type in iface to float", zap.Any("type", fmt.Sprintf("%T", val)))
	case float64	: return val
	case float32	: return float64(val)
	case int		: return float64(val)
	case int32		: return float64(val)
	case int64		: return float64(val)
	case uint32		: return float64(val)
	case uint64		: return float64(val)
	case string		:
		var err error
		var float_val float64
		if float_val, err = strconv.ParseFloat(LowerTrim(val), 64); err != nil {
			log.Error("Failed to convert to float, parse error", zap.Any("val", val), zap.Any("err", err))
			return 0
		}
		return float_val
	}
	return 0
}

func Truncate(in string, max_len int) string {
	if utf8.RuneCountInString(in) <= max_len {return in}
	return string([]rune(in)[:max_len])
}

func LowerTrim(in string) string {return strings.ToLower(strings.TrimSpace(in))}

func Hash(in string) string {
	var bytes_hash = md5.Sum([]byte(in))
	return hex.EncodeToString(bytes_hash[:])
}

func(m *Message) ToMap() (out map[string]any) {
	if m == nil {return}

	out = map[string]any{
		FieldType      	: m._type,
		FieldKind       : m._kind,
		FieldAccessToken: m._accessToken,
		// uuids
		FieldUuid       : m._uuid,
		FieldUserUuid	: m._userUuid,
		FieldWorkerUuid : m._workerUuid,
		// paths
		FieldModule 	: m._module,
		FieldSubject    : m._subject,
		FieldAction		: m._action,
		// some meta info
		FieldIpAddr		: m._ipAddr,
		FieldAppVer		: m._appVer,
		FieldTimeout	: m._timeout,
		FieldStatusCode	: m._statusCode,

		FieldMessage	: m.message,
		// data
		//FieldBytes		: m._bytes,
	}
	// attach payload
	m.lock.RLock()
	for key, val := range m.payload {out[key] = val}
	m.lock.RUnlock()

	return
}

func(m *Message) Marshall() (json.RawMessage, error) {return json.Marshal(m.ToMap())}
func(m *Message) Serialize() (body json.RawMessage) {
	var err error
	if body, err = m.Marshall(); err != nil {
		log.Error("Failed to serialize message", zap.Any("m", m), zap.Any("err", err))
		return nil
	}
	return
}