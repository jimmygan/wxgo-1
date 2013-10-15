package wx

import (
	"reflect"
)

//#include "event.h"
import "C"

func init() {
	mapType("wxEvent", reflect.TypeOf(event{}))
}

type event struct {
	object
	userData interface{}
}

func (e *event) EventObject() Object {
	p := wxPtr(e)
	if p == nil {
		return nil
	}
	return NewObjectFromPtr(C.wxEvent_GetEventObject(p), false)
}

func (e *event) EventType() EventType {
	p := wxPtr(e)
	if p == nil {
		return EventType(-1)
	}
	return EventType(C.wxEvent_GetEventType(p))
}

func (e *event) EventCategory() EventCategory {
	p := wxPtr(e)
	if p == nil {
		return EventCategory(-1)
	}
	return EventCategory(C.wxEvent_GetEventCategory(p))
}

func (e *event) Id() int {
	p := wxPtr(e)
	if p == nil {
		return -1
	}
	return int(C.wxEvent_GetId(p))
}

func (e *event) UserData() interface{} {
	return e.userData
}

func (e *event) setUserData(ud interface{}) {
	e.userData = ud
}

func (e *event) wxUserData() Object {
	p := wxPtr(e)
	if p == nil {
		return nil
	}
	return NewObjectFromPtr(C.wxEvent_GetEventUserData(p), false)
}

func (e *event) Skipped() bool {
	p := wxPtr(e)
	if p == nil {
		return false
	}
	return C.wxEvent_GetSkipped(p) != 0
}

func (e *event) Timestamp() int64 {
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
	EventObject() Object
	EventType() EventType
	EventCategory() EventCategory
	Id() int
	UserData() interface{}
	Skipped() bool
	Timestamp() int64
	IsCommandEvent() bool
	ResumePropagation(propagationLevel int)
	SetEventObject(obj Object)
	SetEventType(t EventType)
	SetId(id int)
	SetTimestamp(timestamp int64)
	ShouldPropagate() bool
	Skip(skip bool)
	StopPropagation()
	wxUserData() Object
	setUserData(ud interface{})
}
