package wx

import (
	"reflect"
	"unsafe"
)

//#include "evthandler.h"
//#include "event.h"
import "C"

func init() {
	mapType("wxEvtHandler", reflect.TypeOf(evtHandler{}))
}

type evtHandler struct {
	object
}

func (h *evtHandler) QueueEvent(event Event) {
	p := wxPtr(h)
	if p == nil {
		return
	}
	event.unhold()
	C.wxEvtHandler_QueueEvent(p, C.WxObjectPtr(wxPtr(event)))
}

func (h *evtHandler) AddPendingEvent(event Event) {
	p := wxPtr(h)
	if p == nil {
		return
	}
	//wxEvtHandler_AddPendingEvent doc: A copy of event is made by the function.
	C.wxEvtHandler_AddPendingEvent(p, wxPtr(event))
}

func (h *evtHandler) ProcessEvent(event Event) {
	p := wxPtr(h)
	if p == nil {
		return
	}
	C.wxEvtHandler_ProcessEvent(p, wxPtr(event))
}

func (h *evtHandler) ProcessEventLocall(event Event) {
	p := wxPtr(h)
	if p == nil {
		return
	}
	C.wxEvtHandler_ProcessEventLocally(p, wxPtr(event))
}

func (h *evtHandler) SafelyProcessEvent(event Event) {
	p := wxPtr(h)
	if p == nil {
		return
	}
	C.wxEvtHandler_SafelyProcessEvent(p, wxPtr(event))
}

func (h *evtHandler) ProcessPendingEvents(event Event) {
	p := wxPtr(h)
	if p == nil {
		return
	}
	C.wxEvtHandler_ProcessPendingEvents(p)
}

func (h *evtHandler) DeletePendingEvents() {
	p := wxPtr(h)
	if p == nil {
		return
	}
	C.wxEvtHandler_DeletePendingEvents(p)
}

type EventFunc func(Event)

//export callEventFunc
func callEventFunc(f unsafe.Pointer, event C.WxObjectPtr) {
	concrete := NewObjectFromPtr(event, false)
	var generic Event
	if concrete != nil { // nil can't convert to any interface.
		generic = concrete.(Event)
	}
	(*(*EventFunc)(f))(generic)
}

const nilEventFunc = "nil EventFunc"

type evtHandlerEntry struct {
	tag       string
	eventType EventType
	id        int
	lastId    int
	f         EventFunc
	seq       Variant
	ud        interface{} // go user data.
}

var evtHandlerTable = make(map[int64]*evtHandlerEntry)
var nextEvtHandlerSeq int64

func evtHandlerEventHubFunc(event Event) {
	v := &variant{}
	v.bindWxPtr(C.wxEvent_GetEventUserData(wxPtr(event)), false)
	seq := v.GetInt64()
	entry := evtHandlerTable[seq]
	if entry == nil {
		return
	}
	event.setEventUserData(entry.ud)
	entry.f(event)
	// The wxEvent object may be created on stack(ie. Menu event).
	event.Release()
	event.setEventUserData(nil)
}

var evtHandlerEventHubFuncVar = EventFunc(evtHandlerEventHubFunc)

func (h *evtHandler) Bind(f EventFunc, funcTag string, eventType EventType, id int, lastId int, userData interface{}) {
	seq := nextEvtHandlerSeq
	nextEvtHandlerSeq++
	v := NewVariantInt64(seq)
	evtHandlerTable[seq] = &evtHandlerEntry{
		tag:       funcTag,
		eventType: eventType,
		id:        id,
		lastId:    lastId,
		f:         f,
		seq:       v,
		ud:        userData,
	}
	h.bind(&evtHandlerEventHubFuncVar, eventType, id, lastId, v)
}

func (h *evtHandler) Unbind(funcTag string, eventType EventType, id int, lastId int, userData interface{}) bool {
	udIsFunc := reflect.TypeOf(userData).Kind() == reflect.Func
	var v Variant
	for _, entry := range evtHandlerTable {
		if entry.tag == funcTag &&
			entry.eventType == eventType &&
			entry.id == id &&
			(entry.lastId == lastId || lastId == ID_ANY) &&
			(udIsFunc || userData == entry.ud) {
			v = entry.seq
			break
		}
	}
	if v == nil {
		return false
	}
	if h.unbind(&evtHandlerEventHubFuncVar, eventType, id, lastId, v) {
		delete(evtHandlerTable, v.GetInt64())
		return true
	}
	return false
}

func (h *evtHandler) bind(f *EventFunc, eventType EventType, id int, lastId int, userData Object) {
	if f == nil {
		panic(nilEventFunc)
	}
	p := wxPtr(h)
	if p == nil {
		return
	}
	userData.unhold()
	C.wxEvtHandler_Bind(p, unsafe.Pointer(f), C.int(eventType), C.int(id), C.int(lastId), wxPtr(userData))
}

func (h *evtHandler) unbind(f *EventFunc, eventType EventType, id int, lastId int, userData Object) bool {
	if f == nil {
		panic(nilEventFunc)
	}
	p := wxPtr(h)
	if p == nil {
		return false
	}
	return goBool(C.wxEvtHandler_Unbind(p, unsafe.Pointer(f), C.int(eventType), C.int(id), C.int(lastId), wxPtr(userData)))
}

func (h *evtHandler) GetEvtHandlerEnabled() bool {
	p := wxPtr(h)
	if p == nil {
		return false
	}
	return C.wxEvtHandler_GetEvtHandlerEnabled(p) != 0
}

func (h *evtHandler) GetNextHandler() EvtHandler {
	p := wxPtr(h)
	if p == nil {
		return nil
	}
	if concrete := NewObjectFromPtr(C.wxEvtHandler_GetNextHandler(p), false); concrete != nil { // nil can't convert to any interface.
		return concrete.(EvtHandler)
	}
	return nil
}

func (h *evtHandler) GetPreviousHandler() EvtHandler {
	p := wxPtr(h)
	if p == nil {
		return nil
	}
	if concrete := NewObjectFromPtr(C.wxEvtHandler_GetPreviousHandler(p), false); concrete != nil {
		return concrete.(EvtHandler)
	}
	return nil
}

func (h *evtHandler) SetEvtHandlerEnabled(enabled bool) {
	p := wxPtr(h)
	if p == nil {
		return
	}
	C.wxEvtHandler_SetEvtHandlerEnabled(p, cBool(enabled))
}

func (h *evtHandler) SetNextHandler(next EvtHandler) {
	p := wxPtr(h)
	next.unhold()
	C.wxEvtHandler_SetNextHandler(p, wxPtr(next))
}

func (h *evtHandler) SetPrevioustHandler(prev EvtHandler) {
	p := wxPtr(h)
	prev.unhold()
	C.wxEvtHandler_SetPrevioustHandler(p, wxPtr(prev))
}

func (h *evtHandler) Unlink() {
	p := wxPtr(h)
	if p == nil {
		return
	}
	C.wxEvtHandler_Unlink(p)
}

func (h *evtHandler) IsUnlinked() bool {
	p := wxPtr(h)
	if p == nil {
		return false
	}
	return C.wxEvtHandler_IsUnlinked(p) != 0
}

type EvtHandler interface {
	Object
	QueueEvent(event Event)
	AddPendingEvent(event Event)
	ProcessEvent(event Event)
	ProcessEventLocall(event Event)
	SafelyProcessEvent(event Event)
	ProcessPendingEvents(event Event)
	DeletePendingEvents()
	Bind(f EventFunc, funTag string, eventType EventType, id int, lastId int, userData interface{})
	Unbind(funTag string, eventType EventType, id int, lastId int, userData interface{}) bool
	GetEvtHandlerEnabled() bool
	GetNextHandler() EvtHandler
	GetPreviousHandler() EvtHandler
	SetEvtHandlerEnabled(enabled bool)
	SetNextHandler(next EvtHandler)
	SetPrevioustHandler(prev EvtHandler)
	Unlink()
	IsUnlinked() bool
}

func NewEvtHandler() EvtHandler {
	h := &evtHandler{}
	h.bindWxPtr(C.wxEvtHandler_New(), true)
	return h
}
