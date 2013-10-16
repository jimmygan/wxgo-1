package wx

//#include "gdidim.h"
//#include <wx/defs.h>
import "C"

// int Width, Height
type Size C.Size

// int X, Y
type Point C.Point

// int X, Y, Size, Point
type Rect C.Rect

var (
	DefaultCoord    int   = int(C.wxDefaultCoord)
	DefaultSize     Size  = Size(C.Get_wxDefaultSize())
	DefaultPosition Point = Point(C.Get_wxDefaultPosition())
)
