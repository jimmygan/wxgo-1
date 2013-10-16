package wx

import (
	"reflect"
)

//#include "gridsizer.h"
import "C"

func init() {
	mapType("wxGridSizer", reflect.TypeOf(gridSizer{}))
}

type gridSizer struct {
	sizer
}

func (s *gridSizer) Cols() int {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return int(C.wxGridSizer_GetCols(p))

}
func (s *gridSizer) Rows() int {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return int(C.wxGridSizer_GetRows(p))

}
func (s *gridSizer) EffectiveColsCount() int {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return int(C.wxGridSizer_GetEffectiveColsCount(p))

}
func (s *gridSizer) EffectiveRowsCount() int {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return int(C.wxGridSizer_GetEffectiveRowsCount(p))

}
func (s *gridSizer) HGap() int {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return int(C.wxGridSizer_GetHGap(p))

}
func (s *gridSizer) VGap() int {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return int(C.wxGridSizer_GetVGap(p))

}
func (s *gridSizer) SetCols(cols int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxGridSizer_SetCols(p, C.int(cols))

}
func (s *gridSizer) SetRows(rows int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxGridSizer_SetRows(p, C.int(rows))

}
func (s *gridSizer) SetHGap(hgap int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxGridSizer_SetHGap(p, C.int(hgap))

}
func (s *gridSizer) SetVGap(vgap int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxGridSizer_SetVGap(p, C.int(vgap))
}

type GridSizer interface {
	Sizer
	Cols() int
	Rows() int
	EffectiveColsCount() int
	EffectiveRowsCount() int
	HGap() int
	VGap() int
	SetCols(cols int)
	SetRows(rows int)
	SetHGap(hgap int)
	SetVGap(vgap int)
}

func NewGridSizer(rows int, cols int, vgap int, hgap int) GridSizer {
	s := &gridSizer{}
	s.bindWxPtr(C.wxGridSizer_New(C.int(rows), C.int(cols), C.int(vgap), C.int(hgap)), true)
	return s
}
