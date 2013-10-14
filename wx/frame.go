package wx

import (
	"reflect"
)

//#include "frame.h"
import "C"

func init() {
	mapType("wxFrame", reflect.TypeOf(frame{}))
}

type frame struct {
	// Not the case. TopLevelWindow......
	window
}

func (f *frame) bindWxPtr(p C.WxObjectPtr, hold bool) {
	// call bindWxPtr() on other "parent" class.
	f.window.bindWxPtr(p, hold)
}

//void* wxFrame_New();
//void wxFrame_Center(void* p);

func (f *frame) Center() {
	p := wxPtr(f)
	if p == nil {
		return
	}
	C.wxFrame_Center(p)
}

func (f *frame) MenuBar() MenuBar {
	p := wxPtr(f)
	if p == nil {
		return nil
	}
	if b := NewObjectFromPtr(C.wxFrame_GetMenuBar(p), false); b != nil {
		return b.(MenuBar)
	}
	return nil
}

func (f *frame) SetMenuBar(menubar MenuBar) {
	p := wxPtr(f)
	if f == nil {
		return
	}
	menubar.unhold()
	C.wxFrame_SetMenuBar(p, wxPtr(menubar))
}

type Frame interface {
	// Not the case. TopLevelWindow......
	Window
	Center()
	MenuBar() MenuBar
	SetMenuBar(MenuBar)
}

func NewFrame(parent Window, id int, title string) Frame {
	f := &frame{}
	f.bindWxPtr(C.wxFrame_New(wxPtr(parent), C.int(id), cString(&title)), true)
	return f
}
