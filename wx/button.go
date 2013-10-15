package wx

import (
	"reflect"
)

//#include "button.h"
import "C"

func init() {
	mapType("wxButton", reflect.TypeOf(button{}))
}

const ButtonNameStr = "button"

type button struct {
	window
}

type Button interface {
	Window
}

func NewButton(parent Window, id int) Button {
	return NewButton2(parent, id, "", nil, nil, 0, nil, ButtonNameStr)
}

func NewButton2(parent Window, id int, label string, pos *Point, size *Size, style int, validator Validator, name string) Button {
	b := &button{}
	var pValidator C.WxObjectPtr
	if validator != nil {
		pValidator = validator.wxPtr()
	}
	// It has a parent, do not hold it.
	b.bindWxPtr(C.wxButton_New(parent.wxPtr(), C.int(id), cString(&label), (*C.Point)(pos), (*C.Size)(size), C.long(style), pValidator, cString(&name)), false)
	return b
}
