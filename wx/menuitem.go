package wx

import (
	"reflect"
)

//#include "menuitem.h"
import "C"

func init() {
	mapType("wxMenuItem", reflect.TypeOf(menuItem{}))
}

type menuItem struct {
	object
}

func (i *menuItem) LabelText() string {
	p := wxPtr(i)
	if p == nil {
		return ""
	}
	return goString(C.wxMenuItem_GetItemLabelText(p))
}

type MenuItem interface {
	Object
	LabelText() string
}

func NewMenuItem(parentMenu Menu, id int, text string, helpString string, kind ItemKind, subManu Menu) MenuItem {
	i := &menuItem{}
	if subManu != nil {
		subManu.unhold()
	}
	i.bindWxPtr(C.wxMenuItem_New(wxPtr(parentMenu), C.int(id), cString(&text), cString(&helpString), C.int(kind), wxPtr(subManu)), true)
	return i
}
