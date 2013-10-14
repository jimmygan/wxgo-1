package wx

import (
	"reflect"
	"unsafe"
)

//#include "threadevent.h"
import "C"

func init() {
	mapType("wxThreadEvent", reflect.TypeOf(threadEvent{}))
}

type threadEvent struct {
	event
}

func (e *threadEvent) GetEventCategory() EventCategory {
	p := wxPtr(e)
	if p == nil {
		return 0
	}
	return EventCategory(C.wxThreadEvent_GetEventCategory(p))
}

func (e *threadEvent) SetPayload(payload unsafe.Pointer) {
	p := wxPtr(e)
	if p == nil {
		return
	}
	C.wxThreadEvent_SetPayload(p, payload)
}

func (e *threadEvent) GetPayload() unsafe.Pointer {
	p := wxPtr(e)
	if p == nil {
		return nil
	}
	return C.wxThreadEvent_GetPayload(p)
}

func (e *threadEvent) GetExtraLong() uint {
	p := wxPtr(e)
	if p == nil {
		return 0
	}
	return uint(C.wxThreadEvent_GetExtraLong(p))
}

func (e *threadEvent) GetInt() int {
	p := wxPtr(e)
	if p == nil {
		return 0
	}
	return int(C.wxThreadEvent_GetInt(p))
}

func (e *threadEvent) GetString() string {
	p := wxPtr(e)
	if p == nil {
		return ""
	}
	return goString(C.wxThreadEvent_GetString(p))
}

func (e *threadEvent) SetLong(n uint) {
	p := wxPtr(e)
	if p == nil {
		return
	}
	C.wxThreadEvent_SetExtraLong(p, C.long(n))
}

func (e *threadEvent) SetInt(n int) {
	p := wxPtr(e)
	if p == nil {
		return
	}
	C.wxThreadEvent_SetInt(p, C.int(n))
}

func (e *threadEvent) SetString(str string) {
	p := wxPtr(e)
	if p == nil {
		return
	}
	C.wxThreadEvent_SetString(p, cString(&str))
}

type ThreadEvent interface {
	Event
	SetPayload(payload unsafe.Pointer)
	GetPayload() unsafe.Pointer
	GetExtraLong() uint
	GetInt() int
	GetString() string
	SetLong(n uint)
	SetInt(n int)
	SetString(str string)
}

func NewThreadEvent() ThreadEvent {
	e := &threadEvent{}
	e.bindWxPtr(C.wxThreadEvent_New(), true)
	return e
}

func NewThreadEvent2(eventType EventType, id int) ThreadEvent {
	e := &threadEvent{}
	e.bindWxPtr(C.wxThreadEvent_New2(C.int(eventType), C.int(id)), true)
	return e
}
