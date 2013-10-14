package wx

import (
	"reflect"
)

//#include "menuevent.h"
import "C"

func init() {
	mapType("wxMenuEvent", reflect.TypeOf(menuEvent{}))
}

type menuEvent struct {
	event
}

//bug: !!!NOT_COMPLETE!!
func (e *menuEvent) GetMenu() C.WxObjectPtr {
	p := wxPtr(e)
	if p == nil {
		return nil
	}
	return C.wxMenuEvent_GetMenu(p)
}

func (e *menuEvent) GetMenuId() int {
	p := wxPtr(e)
	if p == nil {
		return -1
	}
	return int(C.wxMenuEvent_GetMenuId(p))
}

func (e *menuEvent) IsPopup() bool {
	p := wxPtr(e)
	if p == nil {
		return false
	}
	return C.wxMenuEvent_IsPopup(p) != 0
}

type MenuEvent interface {
	Event
	GetMenu() C.WxObjectPtr
	GetMenuId() int
	IsPopup() bool
}

func NewMenuEvent() MenuEvent {
	evt := &menuEvent{}
	evt.bindWxPtr(C.wxMenuEvent_New(C.int(EVT_NULL), 0, nil), true)
	return evt
}

func NewMenuEvent2(eventType EventType, id int, menu C.WxObjectPtr) MenuEvent {
	evt := &menuEvent{}
	evt.bindWxPtr(C.wxMenuEvent_New(C.int(eventType), C.int(id), menu), true)
	return evt
}
