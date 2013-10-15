package wx

import (
	"reflect"
)

//#include "validator.h"
import "C"

func init() {
	mapType("wxValidator", reflect.TypeOf(validator{}))
}

type validator struct {
	evtHandler
}

type Validator interface {
	EvtHandler
}

func NewValidator() Validator {
	v := &validator{}
	v.bindWxPtr(C.wxValidator_New(), true)
	return v
}
