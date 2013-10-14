package wx

import (
	"reflect"
)

//#include "event.h"
//#include "defs.h"
import "C"

func init() {
	mapType("wxEvent", reflect.TypeOf(event{}))
}

type EventCategory int

type EventType int

const (
	EVT_ANY        EventType = C.wxID_ANY
	EVT_FIRST      EventType = 10000
	EVT_USER_FIRST EventType = EVT_FIRST + 2000
	EVT_NULL       EventType = 0
)

var (
	EVT_MENU   EventType = EventType(C.Get_wxEVT_MENU())
	EVT_THREAD           = C.Get_wxEVT_THREAD()
)

type event struct {
	object
	userData interface{}
}

func (e *event) GetEventObject() Object {
	p := wxPtr(e)
	if p == nil {
		return nil
	}
	return NewObjectFromPtr(C.wxEvent_GetEventObject(p), false)
}

func (e *event) GetEventType() EventType {
	p := wxPtr(e)
	if p == nil {
		return EventType(-1)
	}
	return EventType(C.wxEvent_GetEventType(p))
}

func (e *event) GetEventCategory() EventCategory {
	p := wxPtr(e)
	if p == nil {
		return EventCategory(-1)
	}
	return EventCategory(C.wxEvent_GetEventCategory(p))
}

func (e *event) GetId() int {
	p := wxPtr(e)
	if p == nil {
		return -1
	}
	return int(C.wxEvent_GetId(p))
}

func (e *event) GetEventUserData() interface{} {
	return e.userData
}

func (e *event) setEventUserData(ud interface{}) {
	e.userData = ud
}

func (e *event) wxGetEventUserData() Object {
	p := wxPtr(e)
	if p == nil {
		return nil
	}
	return NewObjectFromPtr(C.wxEvent_GetEventUserData(p), false)
}

func (e *event) GetSkipped() bool {
	p := wxPtr(e)
	if p == nil {
		return false
	}
	return C.wxEvent_GetSkipped(p) != 0
}

func (e *event) GetTimestamp() int64 {
	p := wxPtr(e)
	if p == nil {
		return 0
	}
	return int64(C.wxEvent_GetTimestamp(p))
}

func (e *event) IsCommandEvent() bool {
	p := wxPtr(e)
	if p == nil {
		return false
	}
	return C.wxEvent_IsCommandEvent(p) != 0
}

func (e *event) ResumePropagation(propagationLevel int) {
	p := wxPtr(e)
	if p == nil {
		return
	}
	C.wxEvent_ResumePropagation(p, C.int(propagationLevel))
}

func (e *event) SetEventObject(obj Object) {
	p := wxPtr(e)
	if p == nil {
		return
	}
	obj.unhold()
	C.wxEvent_SetEventObject(p, wxPtr(obj))
}

func (e *event) SetEventType(t EventType) {
	p := wxPtr(e)
	if p == nil {
		return
	}
	C.wxEvent_SetEventType(p, C.int(t))
}

func (e *event) SetId(id int) {
	p := wxPtr(e)
	if p == nil {
		return
	}
	C.wxEvent_SetId(p, C.int(id))
}

func (e *event) SetTimestamp(timestamp int64) {
	p := wxPtr(e)
	if p == nil {
		return
	}
	C.wxEvent_SetTimestamp(p, C.long(timestamp))
}

func (e *event) ShouldPropagate() bool {
	p := wxPtr(e)
	if p == nil {
		return false
	}
	return C.wxEvent_ShouldPropagate(p) != 0
}

func (e *event) Skip(skip bool) {
	p := wxPtr(e)
	if p == nil {
		return
	}
	C.wxEvent_Skip(p, cBool(skip))
}

func (e *event) StopPropagation() {
	p := wxPtr(e)
	if p == nil {
		return
	}
	C.wxEvent_StopPropagation(p)
}

type Event interface {
	Object
	GetEventObject() Object
	GetEventType() EventType
	GetEventCategory() EventCategory
	GetId() int
	GetEventUserData() interface{}
	GetSkipped() bool
	GetTimestamp() int64
	IsCommandEvent() bool
	ResumePropagation(propagationLevel int)
	SetEventObject(obj Object)
	SetEventType(t EventType)
	SetId(id int)
	SetTimestamp(timestamp int64)
	ShouldPropagate() bool
	Skip(skip bool)
	StopPropagation()
	wxGetEventUserData() Object
	setEventUserData(ud interface{})
}
