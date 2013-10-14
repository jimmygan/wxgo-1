package wx

import (
	"reflect"
)

//#include "menubar.h"
import "C"

func init() {
	mapType("wxMenuBar", reflect.TypeOf(menuBar{}))
}

const (
	// style	If wxMB_DOCKABLE the menu bar can be detached (wxGTK only).
	MENU_TEAROFF int = 0x0001
)

type menuBar struct {
	window
}

func (b *menuBar) Append(menu Menu, title string) {
	p := wxPtr(b)
	if p == nil {
		return
	}
	menu.unhold()
	C.wxMenuBar_Append(p, wxPtr(menu), cString(&title))
}

func (b *menuBar) Remove(pos int) Menu {
	p := wxPtr(b)
	if p == nil {
		return nil
	}
	if menu := NewObjectFromPtr(C.wxMenuBar_Remove(p, C.long(pos)), true); menu != nil {
		return menu.(Menu)
	}
	return nil
}

type MenuBar interface {
	Window
	Append(menu Menu, title string)
	Remove(pos int) Menu
}

func NewMenuBar() MenuBar {
	return NewMenuBarWithStyle(0)
}

func NewMenuBarWithStyle(style int) MenuBar {
	b := &menuBar{}
	b.bindWxPtr(C.wxMenuBar_New(C.long(style)), true)
	return b
}
