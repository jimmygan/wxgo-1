package wx

import (
	"reflect"
)

//#include "control.h"
import "C"

func init() {
	mapType("wxControl", reflect.TypeOf(control{}))
}

type control struct {
	window
}
