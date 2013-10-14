package wx

import (
	"reflect"
)

//#include "menu.h"
import "C"

func init() {
	mapType("wxMenu", reflect.TypeOf(menu{}))
}

type menu struct {
	evtHandler
}

func (m *menu) Append(id int, item string, helpString string, kind ItemKind) MenuItem {
	p := wxPtr(m)
	if p == nil {
		return nil
	}
	itemStr := item
	if item := NewObjectFromPtr(C.wxMenu_Append(p, C.int(id), cString(&itemStr), cString(&helpString), C.int(kind)), false); item != nil {
		return item.(MenuItem)
	}
	return nil
}

type Menu interface {
	EvtHandler
	Append(id int, item string, helpString string, kind ItemKind) MenuItem
}

func NewMenu(title string) Menu {
	return NewMenuWithStyle(title, 0)
}

func NewMenuWithStyle(title string, style int) Menu {
	m := &menu{}
	m.bindWxPtr(C.wxMenu_New(cString(&title), C.long(style)), true)
	return m
}
