package wx

import (
	"reflect"
)

//#include "flexgridsizer.h"
import "C"

func init() {
	mapType("wxFlexGridSizer", reflect.TypeOf(flexGridSizer{}))
}

type FlexSizerGrowMode int

var (
	FLEX_GROWMODE_NONE      FlexSizerGrowMode = FlexSizerGrowMode(C.Get_wxFLEX_GROWMODE_NONE())
	FLEX_GROWMODE_SPECIFIED FlexSizerGrowMode = FlexSizerGrowMode(C.Get_wxFLEX_GROWMODE_SPECIFIED())
	FLEX_GROWMODE_ALL       FlexSizerGrowMode = FlexSizerGrowMode(C.Get_wxFLEX_GROWMODE_ALL())
)

type flexGridSizer struct {
	gridSizer
}

func (s *flexGridSizer) AddGrowableCol(index int, proportion int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxFlexGridSizer_AddGrowableCol(p, C.long(index), C.int(proportion))
}
func (s *flexGridSizer) AddGrowableRow(index int, proportion int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxFlexGridSizer_AddGrowableRow(p, C.long(index), C.int(proportion))
}
func (s *flexGridSizer) GetFlexibleDirection() int {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return int(C.wxFlexGridSizer_GetFlexibleDirection(p))
}
func (s *flexGridSizer) GetNonFlexibleGrowMode() int {
	p := s.wxPtr()
	if p == nil {
		return 0
	}
	return int(C.wxFlexGridSizer_GetNonFlexibleGrowMode(p))
}
func (s *flexGridSizer) IsColGrowable(index int) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxFlexGridSizer_IsColGrowable(p, C.long(index)))
}
func (s *flexGridSizer) IsRowGrowable(index int) bool {
	p := s.wxPtr()
	if p == nil {
		return false
	}
	return goBool(C.wxFlexGridSizer_IsRowGrowable(p, C.long(index)))
}
func (s *flexGridSizer) RemoveGrowableCol(index int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxFlexGridSizer_RemoveGrowableCol(p, C.long(index))
}
func (s *flexGridSizer) RemoveGrowableRow(index int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxFlexGridSizer_RemoveGrowableRow(p, C.long(index))
}
func (s *flexGridSizer) SetFlexibleDirection(direction int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxFlexGridSizer_SetFlexibleDirection(p, C.int(direction))
}
func (s *flexGridSizer) SetNonFlexibleGrowMode(mode int) {
	p := s.wxPtr()
	if p == nil {
		return
	}
	C.wxFlexGridSizer_SetNonFlexibleGrowMode(p, C.int(mode))
}

type FlexGridSizer interface {
	GridSizer
	AddGrowableCol(index int, proportion int)
	AddGrowableRow(index int, proportion int)
	GetFlexibleDirection() int
	GetNonFlexibleGrowMode() int
	IsColGrowable(index int) bool
	IsRowGrowable(index int) bool
	RemoveGrowableCol(index int)
	RemoveGrowableRow(index int)
	SetFlexibleDirection(direction int)
	SetNonFlexibleGrowMode(mode int)
}

func NewFlexGridSizer(rows int, cols int, vgap int, hgap int) FlexGridSizer {
	s := &flexGridSizer{}
	s.bindWxPtr(C.wxFlexGridSizer_New(C.int(rows), C.int(cols), C.int(vgap), C.int(hgap)), true)
	return s
}
