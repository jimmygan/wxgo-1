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

type SizerItem interface {
	Object
}
