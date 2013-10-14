package wx

import (
	"errors"
	"log"
	"reflect"
	"runtime"
	"unsafe"
)

//#include "object.h"
import "C"

func init() {
	mapType("wxObject", reflect.TypeOf(object{}))
}

var PanicOnInvalidObject bool = true
var DebugObjectBind = true

var errInvalidObject = errors.New("Invalid object")

type base struct {
	p  unsafe.Pointer
	id uint16
}

func (b *base) ptr() unsafe.Pointer {
	if !b.IsValid() {
		if PanicOnInvalidObject {
			panic(errInvalidObject)
		}
		return nil
	}
	return b.p
}

func baseFinalizer(b *base) {
	released := b.Release()
	if released {
		if DebugObjectBind {
			log.Printf("[B]  in finalizer.\n")
		}
	}
}

func (b *base) bindPtr(ptr unsafe.Pointer, del func(unsafe.Pointer), hold bool) {
	b.Release()
	b.id = globalObjectTable.Bind(ptr, del, hold)
	b.p = ptr
	runtime.SetFinalizer(b, baseFinalizer)
}

func (b *base) IsValid() bool {
	valid := b.p != nil && globalObjectTable.IsValid(b.p, b.id)
	return valid
}

func (b *base) Release() bool {
	if !b.IsValid() {
		return false
	}
	globalObjectTable.Release(b.p, b.id)
	b.p = nil
	b.id = 0
	runtime.SetFinalizer(b, nil)
	return true
}

func (b *base) unhold() {
	if !b.IsValid() {
		return
	}
	globalObjectTable.Unhold(b.p, b.id)
}

func (b *base) can_not_implement_this_interface(out_side_of_package_wx) {}

type out_side_of_package_wx int

type Base interface {
	IsValid() bool
	Release() bool

	ptr() unsafe.Pointer
	bindPtr(ptr unsafe.Pointer, del func(unsafe.Pointer), hold bool)
	unhold()
	can_not_implement_this_interface(out_side_of_package_wx)
}

type object struct {
	base
}

func wxObjectDelete(p unsafe.Pointer) {
	C.wxObject_Delete(C.WxObjectPtr(p))
}

func (o *object) wxPtr() C.WxObjectPtr {
	return C.WxObjectPtr(o.ptr())
}

func (o *object) bindWxPtr(ptr C.WxObjectPtr, hold bool) {
	o.bindPtr(unsafe.Pointer(ptr), wxObjectDelete, hold)
}

type Object interface {
	Base
	wxPtr() C.WxObjectPtr
	bindWxPtr(ptr C.WxObjectPtr, hold bool)
}

func NewObject() Object {
	o := &object{}
	o.bindWxPtr(C.wxObject_New(), true)
	return o
}

func wxPtr(obj Object) C.WxObjectPtr {
	if obj == nil {
		return nil
	}
	return obj.wxPtr()
}

func ptr(obj Base) unsafe.Pointer {
	if obj == nil {
		return nil
	}
	return obj.ptr()
}
