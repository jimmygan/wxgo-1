package wx

import (
	"reflect"
	"unsafe"
)

//#include "object.h"
import "C"

var globalTypeMap = make(map[string]reflect.Type)

func mapType(name string, t reflect.Type) {
	globalTypeMap[name] = t
}

// NewObjectFromPtr creates an Object wrapper of the wxPtr.
// The concrete type of returned Object corresponds to the concrete type or
// parent type of *wxPtr.
func NewObjectFromPtr(wxPtr C.WxObjectPtr, hold bool) Object {
	if wxPtr == nil {
		return nil
	}
	className := goString(C.wxObject_GetClassName(wxPtr))
	t := findWxTypeFromTypeMap(className)
	if t == nil {
		return nil
	}
	obj := reflect.New(t).Interface().(Object)
	obj.bindWxPtr(wxPtr, hold)
	return obj
}

func findWxTypeFromTypeMap(className string) reflect.Type {
	var t reflect.Type
	var base string
	if t = globalTypeMap[className]; t != nil {
		return t
	}

	base = goString(C.wxObject_GetBaseClassName1(cString(&className)))
	if base != "" {
		if t = findWxTypeFromTypeMap(base); t != nil {
			return t
		}
	}
	base = goString(C.wxObject_GetBaseClassName2(cString(&className)))
	if base != "" {
		if t = findWxTypeFromTypeMap(base); t != nil {
			return t
		}
	}
	return nil
}

func wxPtrTypeName(wxPtr unsafe.Pointer) string {
	t := findWxTypeFromTypeMap(goString(C.wxObject_GetClassName(C.WxObjectPtr(wxPtr))))
	if t != nil {
		return t.String()
	}
	return "<??>"
}
