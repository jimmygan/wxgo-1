package wx

import (
	"reflect"
)

//#include "boxsizer.h"
import "C"

func init() {
	mapType("wxBoxSizer", reflect.TypeOf(boxSizer{}))
}

type boxSizer struct {
	sizer
}

func (s *boxSizer) Orientation() Orientation {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return Orientation(C.wxBoxSizer_GetOrientation(p))
}

func (s *boxSizer) SetOrientation(orient Orientation) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxBoxSizer_SetOrientation(p, C.int(orient))
}

type BoxSizer interface {
	Sizer
	Orientation() Orientation
	SetOrientation(orient Orientation)
}

func NewBoxSizer(orient Orientation) BoxSizer {
	s := &boxSizer{}
	s.bindWxPtr(C.wxBoxSizer_New(C.int(orient)), true)
	return s
}
