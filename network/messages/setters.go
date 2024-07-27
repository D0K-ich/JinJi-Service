package messages

import (
	"github.com/google/uuid"
	"time"
)

const defaultTimeout = time.Second * 5

func(m *Message) Set(key string, val any) *Message {
	if m == nil {return nil}
	m.lock.Lock()
	m.payload[key] = val
	m.lock.Unlock()
	return m
}

func(m *Message) SetTimeout(seconds int) *Message {
	if seconds < 1 {seconds = 1}
	m._timeout = seconds
	return m
}

func(m *Message) SetModule(module any) *Message {
	if m == nil {return nil}
	m._module = Module(ToString(module))
	return m
}

func(m *Message) SetSubject(subject any) *Message {
	if m == nil {return nil}
	m._subject = Subject(ToString(subject))
	return m
}

func(m *Message) SetAction(action any) *Message {
	if m == nil {return nil}
	m._action = Action(ToString(action))
	return m
}

func(m *Message) SetWorkerUuid(worker_uuid uuid.UUID) *Message {
	if m == nil {return nil}
	m._workerUuid = worker_uuid
	return m
}

func(m *Message) SetIpAddr(ip_addr string) *Message {
	if m == nil {return nil}
	m._ipAddr = ip_addr
	return m
}

func(m *Message) SetAppVer(app_ver string) *Message {
	if m == nil {return nil}
	m._appVer = app_ver
	return m
}

