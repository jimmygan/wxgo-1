package wx

//#include "cgoutil.h"
import "C"

var (
	DefaultPosition = Point{-1, -1}
	DefaultSize     = Size{-1, -1}
)

//type Point struct {
//	X, Y int
//}
type Point C.Point

//type Size struct {
//	Width, Height int
//}
type Size C.Size

type Rect C.Rect

func (rc *Rect) Size() Size {
	return Size{rc.Width, rc.Height}
}

func (rc *Rect) Top() int {
	return int(rc.Y)
}

func (rc *Rect) Left() int {
	return int(rc.X)
}

func (rc *Rect) Right() int {
	return int(rc.X + rc.Width)
}

func (rc *Rect) Bottom() int {
	return int(rc.Y + rc.Height)
}
