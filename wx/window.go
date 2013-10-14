package wx

import (
	"reflect"
)

//#include "window.h"
import "C"

func init() {
	mapType("wxWindow", reflect.TypeOf(window{}))
}

const (
	SIZE_AUTO_HEIGHT int = 0x0002
	SIZE_AUTO_WIDTH  int = 0x0001
	SIZE_AUTO        int = SIZE_AUTO_WIDTH | SIZE_AUTO_HEIGHT
)

const (
	PanelNameStr = "panel"
)

type window struct {
	evtHandler
}

func (w *window) AcceptsFocus() bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	return C.wxWindow_AcceptsFocus(p) != 0
}

func (w *window) AcceptsFocusFromKeyboard() bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	return C.wxWindow_AcceptsFocusFromKeyboard(p) != 0
}

func (w *window) AcceptsFocusRecursively() bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	return C.wxWindow_AcceptsFocusRecursively(p) != 0
}

func (w *window) IsFocusable() bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	return C.wxWindow_IsFocusable(p) != 0
}

func (w *window) CanAcceptFocus() bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	return C.wxWindow_CanAcceptFocus(p) != 0
}

func (w *window) CanAcceptFocusFromKeyboard() bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	return C.wxWindow_CanAcceptFocusFromKeyboard(p) != 0
}

func (w *window) HasFocus() bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	return goBool(C.wxWindow_HasFocus(p))
}

func (w *window) SetCanFocus(can bool) {
	p := wxPtr(w)
	if p == nil {
		return
	}
	C.wxWindow_SetCanFocus(p, cBool(can))
}

func (w *window) SetFocus() {
	p := wxPtr(w)
	if p == nil {
		return
	}
	C.wxWindow_SetFocus(p)
}

func (w *window) SetFocusFromKbd() {
	p := wxPtr(w)
	if p == nil {
		return
	}
	C.wxWindow_SetFocusFromKbd(p)
}

func (w *window) AddChild(child Window) {
	p := wxPtr(w)
	if p == nil {
		return
	}
	child.unhold()
	C.wxWindow_AddChild(p, wxPtr(child))
}

func (w *window) DestroyChildren() bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	return goBool(C.wxWindow_DestroyChildren(p))
}

func (w *window) FindWindow(id int) Window {
	p := wxPtr(w)
	if p == nil {
		return nil
	}
	if concrete := NewObjectFromPtr(C.wxWindow_FindWindow(p, C.long(id)), false); concrete != nil { // nil can't convert to any interface.
		return concrete.(Window)
	}
	return nil
}

func (w *window) FindWindowByName(name string) Window {
	p := wxPtr(w)
	if p == nil {
		return nil
	}
	if concrete := NewObjectFromPtr(C.wxWindow_FindWindowByName(p, cString(&name)), false); concrete != nil {
		return concrete.(Window)
	}
	return nil
}

func (w *window) GetChildrenCount() int {
	p := wxPtr(w)
	if p == nil {
		return 0
	}
	return int(C.wxWindow_GetChildrenCount(p))
}

//BUG: Check the index in range.
func (w *window) GetChild(index int) Window {
	p := wxPtr(w)
	if p == nil {
		return nil
	}
	if concrete := NewObjectFromPtr(C.wxWindow_GetChild(p, C.int(index)), false); concrete != nil {
		return concrete.(Window)
	}
	return nil
}

//BUG: If the child is really a child of w, we need to hold it again.
//Notice that this function is mostly internal to wxWidgets and shouldn't be called by the user code
//func (w *window) RemoveChild(child Window) {
//	p := wxPtr(w)
//	if p == nil {
//		return
//	}
//	C.wxWindow_RemoveChild(p, pChild)
//}

func (w *window) GetGrandParent() Window {
	p := wxPtr(w)
	if p == nil {
		return nil
	}
	if concrete := NewObjectFromPtr(C.wxWindow_GetGrandParent(p), false); concrete != nil {
		return concrete.(Window)
	}
	return nil
}

func (w *window) GetNextSibling() Window {
	p := wxPtr(w)
	if p == nil {
		return nil
	}
	if concrete := NewObjectFromPtr(C.wxWindow_GetNextSibling(p), false); concrete != nil {
		return concrete.(Window)
	}
	return nil
}

func (w *window) GetParent() Window {
	p := wxPtr(w)
	if p == nil {
		return nil
	}
	if concrete := NewObjectFromPtr(C.wxWindow_GetParent(p), false); concrete != nil {
		return concrete.(Window)
	}
	return nil
}

func (w *window) GetPrevSibling() Window {
	p := wxPtr(w)
	if p == nil {
		return nil
	}
	if concrete := NewObjectFromPtr(C.wxWindow_GetPrevSibling(p), false); concrete != nil {
		return concrete.(Window)
	}
	return nil
}

func (w *window) IsDescendant(win Window) bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	return C.wxWindow_IsDescendant(p, wxPtr(win)) != 0
}

func (w *window) Reparent(newParent Window) bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	return C.wxWindow_Reparent(p, wxPtr(newParent)) != 0
}

func (w *window) SetSizeFlag(x, y, width, height, sizeFlags int) {
	p := wxPtr(w)
	if p == nil {
		return
	}
	C.wxWindow_SetSizeFlag(p, C.int(x), C.int(y), C.int(width), C.int(height), C.int(sizeFlags))
}

func (w *window) SetSize(size Size) {
	p := wxPtr(w)
	if p == nil {
		return
	}
	C.wxWindow_SetSize(p, C.Size(size))
}

func (w *window) GetSize() Size {
	p := wxPtr(w)
	if p == nil {
		return Size{}
	}
	return Size(C.wxWindow_GetSize(p))
}

func (w *window) Show(show bool) {
	p := wxPtr(w)
	if p == nil {
		return
	}
	C.wxWindow_Show(p, cBool(show))
}

func (w *window) PopupMenu(menu Menu) bool {
	return w.PopupMenuAtPos(menu, DefaultPosition)
}

func (w *window) PopupMenuAtPos(menu Menu, pos Point) bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	// Doc: The menu does not get deleted by the window.
	// Do not unhold menu.
	return goBool(C.wxWindow_PopupMenu(p, wxPtr(menu), C.Point(pos)))
}

type Window interface {
	EvtHandler
	AcceptsFocus() bool
	AcceptsFocusFromKeyboard() bool
	AcceptsFocusRecursively() bool
	IsFocusable() bool
	CanAcceptFocus() bool
	CanAcceptFocusFromKeyboard() bool
	HasFocus() bool
	SetCanFocus(can bool)
	SetFocus()
	SetFocusFromKbd()
	AddChild(child Window)
	DestroyChildren() bool
	FindWindow(id int) Window
	FindWindowByName(name string) (win Window)
	GetChildrenCount() int
	GetChild(index int) Window
	//RemoveChild(child Window)
	GetGrandParent() Window
	GetNextSibling() Window
	GetParent() Window
	GetPrevSibling() Window
	IsDescendant(win Window) bool
	Reparent(newParent Window) bool
	SetSizeFlag(x, y, width, height, sizeFlags int)
	SetSize(Size)
	GetSize() Size
	Show(show bool)
	PopupMenuAtPos(menu Menu, pos Point) bool
	PopupMenu(menu Menu) bool
}

func NewWindow(parent Window, id int) Window {
	return NewWindow2(parent, id, DefaultPosition, DefaultSize, 0, PanelNameStr)
}

func NewWindow2(parent Window, id int, pos Point, size Size, style uint, name string) Window {
	w := &window{}
	w.bindWxPtr(C.wxWindow_New(wxPtr(parent), C.int(id), C.Point(pos), C.Size(size), C.long(style), cString(&name)), true)
	return w
}
