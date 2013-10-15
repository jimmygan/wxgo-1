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

func (w *window) ChildrenCount() int {
	p := wxPtr(w)
	if p == nil {
		return 0
	}
	return int(C.wxWindow_GetChildrenCount(p))
}

//BUG: Check the index in range.
func (w *window) Child(index int) Window {
	p := wxPtr(w)
	if p == nil {
		return nil
	}
	if concrete := NewObjectFromPtr(C.wxWindow_GetChild(p, C.int(index)), false); concrete != nil {
		return concrete.(Window)
	}
	return nil
}

func (w *window) GrandParent() Window {
	p := wxPtr(w)
	if p == nil {
		return nil
	}
	if concrete := NewObjectFromPtr(C.wxWindow_GetGrandParent(p), false); concrete != nil {
		return concrete.(Window)
	}
	return nil
}

func (w *window) NextSibling() Window {
	p := wxPtr(w)
	if p == nil {
		return nil
	}
	if concrete := NewObjectFromPtr(C.wxWindow_GetNextSibling(p), false); concrete != nil {
		return concrete.(Window)
	}
	return nil
}

func (w *window) Parent() Window {
	p := wxPtr(w)
	if p == nil {
		return nil
	}
	if concrete := NewObjectFromPtr(C.wxWindow_GetParent(p), false); concrete != nil {
		return concrete.(Window)
	}
	return nil
}

func (w *window) PrevSibling() Window {
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

func (w *window) Close() bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	return goBool(C.wxWindow_Close(p, 0))
}

func (w *window) ForceClose() bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	return goBool(C.wxWindow_Close(p, 1))
}

func (w *window) Destroy() {
	p := wxPtr(w)
	if p == nil {
		return
	}
	C.wxWindow_Destroy(p)
}

func (w *window) IsBeingDeleted() bool {
	p := wxPtr(w)
	if p == nil {
		return false
	}
	return goBool(C.wxWindow_IsBeingDeleted(p))
}

func (w *window) Size() Size {
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

func (w *window) Sizer() Sizer {
	p := w.wxPtr()
	if p == nil {
		return nil
	}
	if s := NewObjectFromPtr(C.wxWindow_GetSizer(p), false); s != nil {
		return s.(Sizer)
	}
	return nil
}

func (w *window) SetSizer(sizer Sizer) {
	w.SetSizer2(sizer, true)
}
func (w *window) SetSizer2(sizer Sizer, deleteOld bool) {
	p := wxPtr(w)
	if p == nil {
		return
	}
	// if deleteOld is false, we need to hold the old sizer, if any.
	var oldSizer C.WxObjectPtr
	if !deleteOld {
		oldSizer = C.wxWindow_GetSizer(p) // Use C++ ptr directly.
	}
	sizer.unhold()
	C.wxWindow_SetSizer(p, sizer.wxPtr(), cBool(deleteOld))
	if oldSizer != nil {
		NewObjectFromPtr(oldSizer, true) // Hold it.
	}
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
	ChildrenCount() int
	Child(index int) Window
	//RemoveChild(child Window)
	GrandParent() Window
	NextSibling() Window
	Parent() Window
	PrevSibling() Window
	IsDescendant(win Window) bool
	Reparent(newParent Window) bool
	SetSizeFlag(x, y, width, height, sizeFlags int)
	SetSize(Size)
	Size() Size
	Show(show bool)
	PopupMenuAtPos(menu Menu, pos Point) bool
	PopupMenu(menu Menu) bool
	Close() bool
	ForceClose() bool
	Destroy()
	IsBeingDeleted() bool
	Sizer() Sizer
	SetSizer(sizer Sizer)
	SetSizer2(sizer Sizer, deleteOld bool)
}

func NewWindow(parent Window, id int) Window {
	return NewWindow2(parent, id, DefaultPosition, DefaultSize, 0, PanelNameStr)
}

func NewWindow2(parent Window, id int, pos Point, size Size, style uint, name string) Window {
	w := &window{}
	w.bindWxPtr(C.wxWindow_New(wxPtr(parent), C.int(id), C.Point(pos), C.Size(size), C.long(style), cString(&name)), true)
	return w
}
