package wx

//#include "gdidim.h"
import "C"

// int Width, Height
type Size C.Size

// int X, Y
type Point C.Point

// int X, Y, Size, Point
type Rect C.Rect
