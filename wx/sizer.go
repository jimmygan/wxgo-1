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
		pUserData = userData.wxPtr()
	}
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

type Sizer interface {
	Object
	AddWindow(win Window, flags *SizerFlags) SizerItem
	AddWindowWithUserData(win Window, proportion int, flag int, border int, userData Object) SizerItem
	AddSizer(child Sizer, flags *SizerFlags) SizerItem
	AddSizerWithUserData(child Sizer, proportion int, flag int, border int, userData Object) SizerItem
	AddNoStrechSpacer(size int) SizerItem
	AddSpacerWithUserData(width int, height int, proportion int, flag int, border int, userData Object) SizerItem
	AddStretchSpace(proportion int) SizerItem
}
