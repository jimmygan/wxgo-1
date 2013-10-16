package wx

import (
	"reflect"
)

//#include "sizer.h"
import "C"

func init() {
	mapType("wxSizer", reflect.TypeOf(sizer{}))
}

type sizer struct {
	object
}

func (s *sizer) AddWindow(win Window, flags *SizerFlags) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if item := NewObjectFromPtr(C.wxSizer_AddWindow(p, win.wxPtr(), (*C.SizerFlags)(flags)), false); item != nil {
		return item.(SizerItem)
	}
	return nil
}

func (s *sizer) AddWindowWithUserData(win Window, proportion int, flag int, border int, userData Object) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	var pUserData C.WxObjectPtr
	if userData != nil {
		userData.unhold()
		pUserData = userData.wxPtr()
	}
	if item := NewObjectFromPtr(C.wxSizer_AddWindowWithUserData(p, win.wxPtr(), C.int(proportion), C.int(flag), C.int(border), pUserData), false); item != nil {
		return item.(SizerItem)
	}
	return nil
}

func (s *sizer) AddSizer(child Sizer, flags *SizerFlags) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	child.unhold()
	if item := NewObjectFromPtr(C.wxSizer_AddSizer(p, child.wxPtr(), (*C.SizerFlags)(flags)), false); item != nil {
		return item.(SizerItem)
	}
	return nil
}

func (s *sizer) AddSizerWithUserData(child Sizer, proportion int, flag int, border int, userData Object) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	var pUserData C.WxObjectPtr
	if userData != nil {
		userData.unhold()
		pUserData = userData.wxPtr()
	}
	child.unhold()
	if item := NewObjectFromPtr(C.wxSizer_AddSizerWithUserData(p, child.wxPtr(), C.int(proportion), C.int(flag), C.int(border), pUserData), false); item != nil {
		return item.(SizerItem)
	}
	return nil
}

func (s *sizer) AddNoStrechSpacer(size int) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if item := NewObjectFromPtr(C.wxSizer_AddNoStretchSpacer(p, C.int(size)), false); item != nil {
		return item.(SizerItem)
	}
	return nil
}

func (s *sizer) AddSpacerWithUserData(width int, height int, proportion int, flag int, border int, userData Object) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	var pUserData C.WxObjectPtr
	if userData != nil {
		userData.unhold()
		pUserData = userData.wxPtr()
	}
	if item := NewObjectFromPtr(C.wxSizer_AddSpacerWithUserData(p, C.int(width), C.int(height), C.int(proportion), C.int(flag), C.int(border), pUserData), false); item != nil {
		return item.(SizerItem)
	}
	return nil
}

func (s *sizer) AddStretchSpace(proportion int) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if item := NewObjectFromPtr(C.wxSizer_AddStretchSpacer(p, C.int(proportion)), false); item != nil {
		return item.(SizerItem)
	}
	return nil
}

func (s *sizer) AddSizerItem(item SizerItem) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	item.unhold()
	C.wxSizer_Add(p, item.wxPtr())
	return item
}

func (s *sizer) Clear(deleteWindows bool) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	itemCount := s.ItemCount()
	for i := 0; i < itemCount; i++ {
		//BUG: hold the items
		//item := s.Item(i)
	}
	C.wxSizer_Clear(p, cBool(deleteWindows))
}

//func (s *sizer) ComputeFittingClientSize(window Window) Size {
//	p := s.wxPtr()
//	if p == nil {
//		return DefaultSize
//	}
//	return Size(C.wxSizer_ComputeFittingClientSize(p, window.wxPtr()))
//}
//func (s *sizer) ComputeFittingWindowSize(window Window) Size {
//	p := s.wxPtr()
//	if p == nil {
//		return DefaultSize
//	}
//	return Size(C.wxSizer_ComputeFittingWindowSize(p, window.wxPtr()))
//}
func (s *sizer) DetachWindow(window Window) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	// The window is managed by parent window.
	return goBool(C.wxSizer_DetachWindow(p, window.wxPtr()))
}
func (s *sizer) DetachSizer(sizer Sizer) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	result := goBool(C.wxSizer_DetachSizer(p, sizer.wxPtr()))
	if result {
		sizer.hold()
	}
	return result
}
func (s *sizer) Detach(index int) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	//BUG: get the Sizer from item and hold it.
	//item := s.Item(index)
	return goBool(C.wxSizer_Detach(p, C.int(index)))

}
func (s *sizer) Fit(window Window) Size {
	p := s.wxPtr()
	if p == nil {
		return DefaultSize
	}
	return Size(C.wxSizer_Fit(p, window.wxPtr()))
}
func (s *sizer) FitInside(window Window) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizer_FitInside(p, window.wxPtr())
}

//func (s *sizer) InformFirstDirection(direction Direction, size int, availOtherDir Direction) bool {
//	p := s.wxPtr()
//	if p == nil {
//		return false
//	}
//	return goBool(C.wxSizer_InformFirstDirection(p, C.int(direction), C.int(size), C.int(availOtherDir)))
//}
func (s *sizer) ContainingWindow() Window {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if obj := NewObjectFromPtr(C.wxSizer_GetContainingWindow(p), false); obj != nil {
		return obj.(Window)
	}
	return nil
}
func (s *sizer) SetContainingWindow(window Window) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizer_SetContainingWindow(p, window.wxPtr())
}
func (s *sizer) ItemCount() int {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return int(C.wxSizer_GetItemCount(p))
}
func (s *sizer) Item(index int) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if obj := NewObjectFromPtr(C.wxSizer_GetItem(p, C.long(index)), false); obj != nil {
		return obj.(SizerItem)
	}
	return nil
}
func (s *sizer) ItemByWindow(window Window, recursive bool) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if obj := NewObjectFromPtr(C.wxSizer_GetItemByWindow(p, window.wxPtr(), cBool(recursive)), false); obj != nil {
		return obj.(SizerItem)
	}
	return nil
}
func (s *sizer) ItemBySizer(sizer Sizer, recursive bool) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if obj := NewObjectFromPtr(C.wxSizer_GetItemBySizer(p, sizer.wxPtr(), cBool(recursive)), false); obj != nil {
		return obj.(SizerItem)
	}
	return nil
}
func (s *sizer) ItemById(id int, recursive bool) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if obj := NewObjectFromPtr(C.wxSizer_GetItemById(p, C.int(id), cBool(recursive)), false); obj != nil {
		return obj.(SizerItem)
	}
	return nil
}
func (s *sizer) MinSize() Size {
	p := s.wxPtr()
	if p == nil {
		return DefaultSize
	}
	return Size(C.wxSizer_GetMinSize(p))
}
func (s *sizer) Position() Point {
	p := s.wxPtr()
	if p == nil {
		return DefaultPosition
	}
	return Point(C.wxSizer_GetPosition(p))
}
func (s *sizer) Size() Size {
	p := s.wxPtr()
	if p == nil {
		return DefaultSize
	}
	return Size(C.wxSizer_GetSize(p))
}
func (s *sizer) HideWindow(window Window, recursive bool) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizer_HideWindow(p, window.wxPtr(), cBool(recursive)))
}
func (s *sizer) HideSizer(sizer Sizer, recursive bool) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizer_HideSizer(p, sizer.wxPtr(), cBool(recursive)))
}
func (s *sizer) HideItem(index int) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizer_Hide(p, C.long(index)))
}
func (s *sizer) InsertWindow(index int, window Window, flags *SizerFlags) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if obj := NewObjectFromPtr(C.wxSizer_InsertWindow(p, C.long(index), window.wxPtr(), (*C.SizerFlags)(flags)), false); obj != nil {
		return obj.(SizerItem)
	}
	return nil
}
func (s *sizer) InsertWindowWithUserData(index int, window Window, proportion int, flag int, border int, userData Object) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	var pUserData C.WxObjectPtr
	if userData != nil {
		userData.unhold()
		pUserData = userData.wxPtr()
	}
	if obj := NewObjectFromPtr(C.wxSizer_InsertWindowWithUserData(p, C.long(index), window.wxPtr(), C.int(proportion), C.int(flag), C.int(border), pUserData), false); obj != nil {
		return obj.(SizerItem)
	}
	return nil
}
func (s *sizer) InsertSizer(index int, sizer Sizer, flags *SizerFlags) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if obj := NewObjectFromPtr(C.wxSizer_InsertSizer(p, C.long(index), sizer.wxPtr(), (*C.SizerFlags)(flags)), false); obj != nil {
		sizer.unhold()
		return obj.(SizerItem)
	}
	return nil
}
func (s *sizer) InsertSizerWithUserData(index int, sizer Sizer, proportion int, flag int, border int, userData Object) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	var pUserData C.WxObjectPtr
	if userData != nil {
		userData.unhold()
		pUserData = userData.wxPtr()
	}
	if obj := NewObjectFromPtr(C.wxSizer_InsertSizerWithUserData(p, C.long(index), sizer.wxPtr(), C.int(proportion), C.int(flag), C.int(border), pUserData), false); obj != nil {
		sizer.unhold()
		return obj.(SizerItem)
	}
	return nil
}
func (s *sizer) InsertNoStretchSpacer(index int, size int) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if obj := NewObjectFromPtr(C.wxSizer_InsertNoStretchSpacer(p, C.long(index), C.int(size)), false); obj != nil {
		return obj.(SizerItem)
	}
	return nil
}
func (s *sizer) InsertSpacerWithUserData(index int, width int, height int, proportion int, flag int, border int, userData Object) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	var pUserData C.WxObjectPtr
	if userData != nil {
		userData.unhold()
		pUserData = userData.wxPtr()
	}
	if obj := NewObjectFromPtr(C.wxSizer_InsertSpacerWithUserData(p, C.long(index), C.int(width), C.int(height), C.int(proportion), C.int(flag), C.int(border), pUserData), false); obj != nil {
		return obj.(SizerItem)
	}
	return nil
}
func (s *sizer) InsertStretchSpacer(index int, proportion int) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	if obj := NewObjectFromPtr(C.wxSizer_InsertStretchSpacer(p, C.long(index), C.int(proportion)), false); obj != nil {
		return obj.(SizerItem)
	}
	return nil
}
func (s *sizer) Insert(index int, item SizerItem) SizerItem {
	p := s.wxPtr()
	if p == nil {
		return nil
	}
	C.wxSizer_Insert(p, C.long(index), item.wxPtr())
	item.unhold()
	return item
}
func (s *sizer) IsEmpty() bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizer_IsEmpty(p))
}
func (s *sizer) IsWindowShown(window Window) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizer_IsWindowShown(p, window.wxPtr()))
}
func (s *sizer) IsSizerShown(sizer Sizer) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizer_IsSizerShown(p, sizer.wxPtr()))
}
func (s *sizer) Layout() {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizer_Layout(p)
}
func (s *sizer) RemoveSizer(sizer Sizer) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizer_RemoveSizer(p, sizer.wxPtr()))
}
func (s *sizer) Remove(index int) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizer_Remove(p, C.int(index)))
}
func (s *sizer) ReplaceWindow(oldWindow Window, newWindow Window, recursive bool) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizer_ReplaceWindow(p, oldWindow.wxPtr(), newWindow.wxPtr(), cBool(recursive)))
}
func (s *sizer) ReplaceSizer(oldSizer Sizer, newSizer Sizer, recursive bool) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	result := goBool(C.wxSizer_ReplaceSizer(p, oldSizer.wxPtr(), newSizer.wxPtr(), cBool(recursive)))
	if result {
		oldSizer.hold()
		newSizer.unhold()
	}
	return result
}
func (s *sizer) Replace(index int, newItem SizerItem) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	oldItem := s.Item(index)
	result := goBool(C.wxSizer_Replace(p, C.long(index), newItem.wxPtr()))
	if result {
		oldItem.hold()
		newItem.unhold()
	}
	return result
}
func (s *sizer) SetDimension(pos *Point, size *Size) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizer_SetDimension(p, (*C.Point)(pos), (*C.Size)(size))
}
func (s *sizer) SetMinSize(size *Size) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizer_SetMinSize(p, (*C.Size)(size))
}
func (s *sizer) SetSizeHints(window Window) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizer_SetSizeHints(p, window.wxPtr())
}
func (s *sizer) ShowWindow(window Window, show bool, recursive bool) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizer_ShowWindow(p, window.wxPtr(), cBool(show), cBool(recursive)))
}
func (s *sizer) ShowSizer(sizer Sizer, show bool, recursive bool) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizer_ShowSizer(p, sizer.wxPtr(), cBool(show), cBool(recursive)))
}
func (s *sizer) Show(index int, show bool) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxSizer_Show(p, C.long(index), cBool(show)))
}
func (s *sizer) ShowAll(show bool) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxSizer_ShowAll(p, cBool(show))
}

type Sizer interface {
	Object
	AddWindow(win Window, flags *SizerFlags) SizerItem
	AddWindowWithUserData(win Window, proportion int, flag int, border int, userData Object) SizerItem
	AddSizer(child Sizer, flags *SizerFlags) SizerItem
	AddSizerWithUserData(child Sizer, proportion int, flag int, border int, userData Object) SizerItem
	AddNoStrechSpacer(size int) SizerItem
	AddSpacerWithUserData(width int, height int, proportion int, flag int, border int, userData Object) SizerItem
	AddStretchSpace(proportion int) SizerItem
	Clear(deleteWindows bool)
	//ComputeFittingClientSize(window Window) Size
	//ComputeFittingWindowSize(window Window) Size
	DetachWindow(window Window) bool
	DetachSizer(sizer Sizer) bool
	Detach(index int) bool
	Fit(window Window) Size
	FitInside(window Window)
	//InformFirstDirection(direction Direction, size int, availOtherDir Direction) bool
	ContainingWindow() Window
	SetContainingWindow(window Window)
	ItemCount() int
	Item(index int) SizerItem
	ItemByWindow(window Window, recursive bool) SizerItem
	ItemBySizer(sizer Sizer, recursive bool) SizerItem
	ItemById(id int, recursive bool) SizerItem
	MinSize() Size
	Position() Point
	Size() Size
	HideWindow(window Window, recursive bool) bool
	HideSizer(sizer Sizer, recursive bool) bool
	HideItem(index int) bool
	InsertWindow(index int, window Window, flags *SizerFlags) SizerItem
	InsertWindowWithUserData(index int, window Window, proportion int, flag int, border int, userData Object) SizerItem
	InsertSizer(index int, sizer Sizer, flags *SizerFlags) SizerItem
	InsertSizerWithUserData(index int, sizer Sizer, proportion int, flag int, border int, userData Object) SizerItem
	InsertNoStretchSpacer(index int, size int) SizerItem
	InsertSpacerWithUserData(index int, width int, height int, proportion int, flag int, border int, userData Object) SizerItem
	InsertStretchSpacer(index int, proportion int) SizerItem
	Insert(index int, item SizerItem) SizerItem
	IsEmpty() bool
	IsWindowShown(window Window) bool
	IsSizerShown(sizer Sizer) bool
	Layout()
	RemoveSizer(sizer Sizer) bool
	Remove(index int) bool
	ReplaceWindow(oldWindow Window, newWindow Window, recursive bool) bool
	ReplaceSizer(oldSizer Sizer, newSizer Sizer, recursive bool) bool
	Replace(index int, newItem SizerItem) bool
	SetDimension(pos *Point, size *Size)
	SetMinSize(size *Size)
	SetSizeHints(window Window)
	ShowWindow(window Window, show bool, recursive bool) bool
	ShowSizer(sizer Sizer, show bool, recursive bool) bool
	Show(index int, show bool) bool
	ShowAll(show bool)
}
