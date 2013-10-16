package wx

import (
	"reflect"
)

//#include "sizeritem.h"
import "C"

func init() {
	mapType("wxSizerItem", reflect.TypeOf(sizerItem{}))
}

type sizerItem struct {
	object
}

func (s *sizerItem) SetWindow(window Window) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizerItem_AssignWindow(p, window.wxPtr())
}
func (s *sizerItem) SetSizer(sizer Sizer) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizerItem_AssignSizer(p, sizer.wxPtr())
}
func (s *sizerItem) SetSpacer(size *Size) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizerItem_AssignSpacer(p, (*C.Size)(size))
}
func (s *sizerItem) DeleteWindows() {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizerItem_DeleteWindows(p)
}
func (s *sizerItem) DetachSizer() {
	p := s.wxPtr()
	if p == nil {
		return
	}
	if s.IsSizer() {
		s.Sizer().hold()
	}
	C.wxSizerItem_DetachSizer(p)
}
func (s *sizerItem) Border() int {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return int(C.wxSizerItem_GetBorder(p))
}
func (s *sizerItem) Flag() int {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return int(C.wxSizerItem_GetFlag(p))
}
func (s *sizerItem) Id() int {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return int(C.wxSizerItem_GetId(p))
}
func (s *sizerItem) MinSize() Size {
	p := s.wxPtr()
	if p == nil {
		return DefaultSize
	}
	return Size(C.wxSizerItem_GetMinSize(p))
}
func (s *sizerItem) SetMinSize(size *Size) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizerItem_SetMinSize(p, (*C.Size)(size))
}
func (s *sizerItem) Position() Point {
	p := s.wxPtr()
	if p == nil {
		return DefaultPosition
	}
	return Point(C.wxSizerItem_GetPosition(p))
}
func (s *sizerItem) Proportion() int {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return int(C.wxSizerItem_GetProportion(p))
}
func (s *sizerItem) Ratio() float32 {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return float32(C.wxSizerItem_GetRatio(p))
}
func (s *sizerItem) Rect() Rect {
	p := s.wxPtr()
	if p == nil {
		return Rect{-1, -1, -1, -1}
	}
	return Rect(C.wxSizerItem_GetRect(p))
}
func (s *sizerItem) Size() Size {
	p := s.wxPtr()
	if p == nil {
		return DefaultSize
	}
	return Size(C.wxSizerItem_GetSize(p))
}
func (s *sizerItem) Sizer() Sizer {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if obj := NewObjectFromPtr(C.wxSizerItem_GetSizer(p), false); obj != nil {
		return obj.(Sizer)
	}
	return nil
}
func (s *sizerItem) Spacer() Size {
	p := s.wxPtr()
	if p == nil {
		return DefaultSize
	}
	return Size(C.wxSizerItem_GetSpacer(p))
}
func (s *sizerItem) UserData() Object {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	return NewObjectFromPtr(C.wxSizerItem_GetUserData(p), false)
}
func (s *sizerItem) Window() Window {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if obj := NewObjectFromPtr(C.wxSizerItem_GetWindow(p), false); obj != nil {
		return obj.(Window)
	}
	return nil
}
func (s *sizerItem) IsShown() bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizerItem_IsShown(p))
}
func (s *sizerItem) IsSizer() bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizerItem_IsSizer(p))
}
func (s *sizerItem) IsSpacer() bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizerItem_IsSpacer(p))
}
func (s *sizerItem) IsWindow() bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizerItem_IsWindow(p))
}
func (s *sizerItem) SetBorder(border int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizerItem_SetBorder(p, C.int(border))
}
func (s *sizerItem) SetDimension(pos *Point, size *Size) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizerItem_SetDimension(p, (*C.Point)(pos), (*C.Size)(size))
}
func (s *sizerItem) SetFlag(flag int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizerItem_SetFlag(p, C.int(flag))
}
func (s *sizerItem) SetId(id int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizerItem_SetId(p, C.int(id))
}
func (s *sizerItem) SetInitSize(size *Size) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizerItem_SetInitSize(p, (*C.Size)(size))
}
func (s *sizerItem) SetProportion(proportion int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizerItem_SetProportion(p, C.int(proportion))
}
func (s *sizerItem) SetUserData(userData Object) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	var pUserData C.WxObjectPtr
	if userData != nil {
		userData.unhold()
		pUserData = userData.wxPtr()
	}
	C.wxSizerItem_SetUserData(p, pUserData)
}
func (s *sizerItem) SetRatio(ratio float32) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizerItem_SetRatio(p, C.float(ratio))
}
func (s *sizerItem) SetRatioSize(size *Size) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizerItem_SetRatioSize(p, (*C.Size)(size))
}

type SizerItem interface {
	Object
	SetWindow(window Window)
	SetSizer(sizer Sizer)
	SetSpacer(size *Size)
	DeleteWindows()
	DetachSizer()
	Border() int
	Flag() int
	Id() int
	MinSize() Size
	SetMinSize(size *Size)
	Position() Point
	Proportion() int
	Ratio() float32
	Rect() Rect
	Size() Size
	Sizer() Sizer
	Spacer() Size
	UserData() Object
	Window() Window
	IsShown() bool
	IsSizer() bool
	IsSpacer() bool
	IsWindow() bool
	SetBorder(border int)
	SetDimension(pos *Point, size *Size)
	SetFlag(flag int)
	SetId(id int)
	SetInitSize(size *Size)
	SetProportion(proportion int)
	SetUserData(userData Object)
	SetRatio(ratio float32)
	SetRatioSize(size *Size)
}

func NewSpacerSizeItem(width int, height int, proportion int, flag int, border int, userData Object) SizerItem {
	s := &sizerItem{}
	var pUserData C.WxObjectPtr
	if userData != nil {
		userData.unhold()
		pUserData = userData.wxPtr()
	}
	s.bindWxPtr(C.wxSizerItem_NewSpacer(C.int(width), C.int(height), C.int(proportion), C.int(flag), C.int(border), pUserData), true)
	return s
}
func (s *sizerItem) NewWindowSizerItem(window Window, flags *SizerFlags) SizerItem {
	item := &sizerItem{}
	item.bindWxPtr(C.wxSizerItem_NewWindow(window.wxPtr(), (*C.SizerFlags)(flags)), true)
	return item
}
func (s *sizerItem) NewSizerItem(sizer Sizer, flags *SizerFlags) SizerItem {
	item := &sizerItem{}
	item.bindWxPtr(C.wxSizerItem_NewSizer(sizer.wxPtr(), (*C.SizerFlags)(flags)), true)
	return item
}
