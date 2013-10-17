package wx

import (
	"reflect"
)

//#include "control.h"
import "C"

func init() {
	mapType("wxControl", reflect.TypeOf(control{}))
}

type control struct {
	window
}

func (c *control) Command(event Event) {
	p := wxPtr(c)
	if p == nil {
		return
	}
	C.wxControl_Command(p, event.wxPtr())
}
func (c *control) Label() string {
	p := wxPtr(c)
	if p == nil {
		return ""
	}
	return goString(C.wxControl_GetLabel(p))
}
func (c *control) LabelText() string {
	p := wxPtr(c)
	if p == nil {
		return ""
	}
	return goString(C.wxControl_GetLabelText(p))
}
func (c *control) SizeFromTextSize(size *Size) Size {
	p := wxPtr(c)
	if p == nil {
		return DefaultSize
	}
	return Size(C.wxControl_GetSizeFromTextSize(p, (*C.Size)(size)))
}
func (c *control) SetLabelText(text string) {
	p := wxPtr(c)
	if p == nil {
		return
	}
	C.wxControl_SetLabelText(p, cString(&text))
}
func (c *control) SetLabelMarkup(markup string) {
	p := wxPtr(c)
	if p == nil {
		return
	}
	C.wxControl_SetLabelMarkup(p, cString(&markup))
}

type Control interface {
	Window
	Command(event Event)
	Label() string
	LabelText() string
	SizeFromTextSize(size *Size) Size
	SetLabelText(text string)
	SetLabelMarkup(markup string)
}

func New_Control(parent Window, id int, pos *Point, size *Size, style int, validator Validator, name string) Control {
	c := &control{}
	c.bindWxPtr(C.wxControl_New(parent.wxPtr(), C.int(id), (*C.Point)(pos), (*C.Size)(size), C.long(style), validator.wxPtr(), cString(&name)), true)
	return c
}
