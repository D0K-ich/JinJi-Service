package messages

import (
	"fmt"
	"go.uber.org/zap"
	"net/url"
	"strings"
	"time"
)

func(m *Message) Uuid()				(uuid string)			{if m == nil {return} else {return m._uuid}}
func(m *Message) Type()				(mtype Type)			{if m == nil {return} else {return m._type}}
func(m *Message) Kind()				(kind string)			{if m == nil {return} else {return m._kind}}
func(m *Message) Module()			(module Module)			{if m == nil {return} else {return m._module}}
func(m *Message) Subject()			(subject Subject)		{if m == nil {return} else {return m._subject}}
func(m *Message) Action()			(action Action)			{if m == nil {return} else {return m._action}}
func(m *Message) SubjectAction()	(sa string)				{if m == nil {return} else {return strings.Trim(strings.TrimSpace(strings.Join([]string{m.Subject().String(), m.Action().String()}, "/")), "/")}}
func(m *Message) DataBytes()		(data []byte)			{if m == nil {return} else {return m._bytes}}
func(m *Message) AccessToken()		(a_token string)		{if m == nil {return} else {return m._accessToken}}
func(m *Message) MetaString()		(meta string)			{if m == nil {return} else {return fmt.Sprintf("%s:%s:%s:%s", m.Uuid(), m.Type(), m.Module(), m.SubjectAction())}}

func(m *Message) GetIpAddr()		(ip string)				{if m == nil {return} else {return m._ipAddr}}
func(m *Message) GetAppVer()		(app_ver string)		{if m == nil {return} else {return m._appVer}}

func(m *Message) Timeout() (timeout time.Duration) {
	if m == nil {return 0}
	timeout = time.Second * time.Duration(ToInt(m._timeout))
	if timeout > time.Second {return timeout}

	log.Warn("skip to default timeout 5 sec due field timeout is lower than must", zap.Any("val", timeout))
	return defaultTimeout
}

// StatusCode todo improvements needed
func(m *Message) StatusCode() int {
	if m == nil {return 0}
	if code := m._statusCode; code > 0 {return code}
	if m.Type() == TypeError {return 400}
	return 200
}

func(m *Message) Payload() (payload map[string]any) {
	if m == nil {return}
	payload = make(map[string]any)		// todo add a deep copy
	for key, val := range m.payload {payload[key] = val}
	return
}

// ======================
// PAYLOAD TRANSFORMATION

func(m *Message) Get(field string) (value any) {
	if m == nil {return}
	m.lock.RLock()
	value = m.payload[field]
	m.lock.RUnlock()
	return
}

func(m *Message) String(field string) (value string) {
	if m == nil {return}
	m.lock.RLock()
	value = strings.TrimSpace(ToString(m.payload[field]))
	m.lock.RUnlock()
	return
}

func(m *Message) StringSlice(field string) (values []string) {
	if m == nil {return}
	m.lock.RLock()
	values = ToStringSlice(m.payload[field])
	m.lock.RUnlock()
	return
}

func(m *Message) FullUrl(field string) (value *url.URL) {
	if m == nil {return}
	m.lock.RLock()

	var err error
	if value, err = url.Parse(strings.TrimSpace(ToString(m.payload[field]))); err != nil {
		log.Warn("error on parse url", zap.Any("err", err))
		value = &url.URL{}
	}
	m.lock.RUnlock()
	return
}

func(m *Message) Int(field string) (value int) {
	if m == nil {return}
	m.lock.RLock()
	value = ToInt(m.payload[field])
	m.lock.RUnlock()
	return
}

func(m *Message) Bool(field string) (value bool) {
	if m == nil {return}
	m.lock.RLock()
	value = ToBool(m.payload[field])
	m.lock.RUnlock()
	return
}

func(m *Message) Float(field string) (value float64) {
	if m == nil {return}
	m.lock.RLock()
	value = ToFloat(m.payload[field])
	m.lock.RUnlock()
	return
}

func(m *Message) Version() (value int64) {
	if m == nil {return}
	m.lock.RLock()
	value = ToInt64(m.payload["version"])
	m.lock.RUnlock()
	return
}
