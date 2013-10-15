package wx

import (
	"reflect"
)

//#include "variant.h"
import "C"

func init() {
	mapType("wxVariant", reflect.TypeOf(variant{}))
}

type variant struct {
	object
}

func (v *variant) Int64() int64 {
	p := wxPtr(v)
	if p == nil {
		return 0
	}
	return int64(C.wxVariant_GetInt64(p))
}

func (v *variant) String() string {
	p := wxPtr(v)
	if p == nil {
		return ""
	}
	return goString(C.wxVariant_GetString(p))
}

type Variant interface {
	Object
	Int64() int64
	String() string
}

func NewVariantInt64(n int64) Variant {
	v := &variant{}
	v.bindWxPtr(C.wxVariant_NewInt64(C.longlong(n)), true)
	return v
}

func NewVariantString(str string) Variant {
	v := &variant{}
	v.bindWxPtr(C.wxVariant_NewString(cString(&str)), true)
	return v
}
