package wx

import (
	"reflect"
	"unsafe"
)

//#include "app.h"
import "C"

func init() {
	mapType("MyApp", reflect.TypeOf(app{}))
}

type app struct {
	evtHandler
}

func (a *app) bindWxPtr(p C.WxObjectPtr, hold bool) {
	a.evtHandler.bindWxPtr(p, hold)
	a.Bind(appRunInUIThreadEventHandler, "[App.RunInUI Event Handler]", evt_RUN_IN_UI)
}

func (a *app) SetTopWindow(w Window) {
	p := wxPtr(a)
	if p == nil {
		return
	}
	w.unhold()
	C.wxApp_SetTopWindow(p, wxPtr(w))
}

func (a *app) ExitMainLoop() {
	p := wxPtr(a)
	if p == nil {
		return
	}
	C.wxApp_ExitMainLoop(p)
}

var evt_RUN_IN_UI = EVT_USER_FIRST + 0x0FFFFFFF

func appRunInUIThreadEventHandler(event Event) {
	payload := (*runInUIPayload)(event.(ThreadEvent).Payload())
	payload.r <- payload.f()
}

type runInUIPayload struct {
	f func() interface{}
	r chan<- interface{}
}

// Run runs f in UI goroutine. Do not call this method in the UI goroutine.
// This function blocks until f has been executed and then returns the return
// value of f.
func (a *app) RunInUI(f func() interface{}) interface{} {
	e := NewThreadEvent2(evt_RUN_IN_UI, ID_ANY)
	r := make(chan interface{})
	defer close(r)
	// This function blocks until appRunInUIThreadEventHandler() finished using
	// payload. Payload will survive GC, it is safe.
	payload := runInUIPayload{f: f, r: r}
	e.SetPayload(unsafe.Pointer(&payload))
	a.QueueEvent(e)
	// return <-r doesn't work.
	ret := <-r
	// Must reference payload after chan receive to keep it alive.
	//BUG: Does this safe?? A REALLY cleaver compiler can optimize this line out.
	payload.r = nil
	return ret
}

type App interface {
	EvtHandler
	SetTopWindow(Window)
	ExitMainLoop()
	RunInUI(f func() interface{}) interface{}
}
