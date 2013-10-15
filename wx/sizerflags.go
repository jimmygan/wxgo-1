package wx

//#include "sizerflags.h"
//#include <wx/defs.h>
import "C"

type Aliment int

const (
	ALIGN_INVALID           Aliment = C.wxALIGN_INVALID
	ALIGN_NOT               Aliment = C.wxALIGN_NOT
	ALIGN_CENTER_HORIZONTAL Aliment = C.wxALIGN_CENTER_HORIZONTAL
	ALIGN_CENTRE_HORIZONTAL Aliment = C.wxALIGN_CENTRE_HORIZONTAL
	ALIGN_LEFT              Aliment = C.wxALIGN_LEFT
	ALIGN_TOP               Aliment = C.wxALIGN_TOP
	ALIGN_RIGHT             Aliment = C.wxALIGN_RIGHT
	ALIGN_BOTTOM            Aliment = C.wxALIGN_BOTTOM
	ALIGN_CENTER_VERTICAL   Aliment = C.wxALIGN_CENTER_VERTICAL
	ALIGN_CENTRE_VERTICAL   Aliment = C.wxALIGN_CENTRE_VERTICAL
	ALIGN_CENTER            Aliment = C.wxALIGN_CENTER
	ALIGN_CENTRE            Aliment = C.wxALIGN_CENTRE
	ALIGN_MASK              Aliment = C.wxALIGN_MASK
)

type Orientation int

const (
	HORIZONTAL       Orientation = C.wxHORIZONTAL
	VERTICAL         Orientation = C.wxVERTICAL
	BOTH             Orientation = C.wxBOTH
	ORIENTATION_MASK Orientation = C.wxORIENTATION_MASK
)

type Direction int

var (
	LEFT           Direction = C.wxLEFT
	RIGHT          Direction = C.wxRIGHT
	UP             Direction = C.wxUP
	DOWN           Direction = C.wxDOWN
	TOP            Direction = C.wxTOP
	BOTTOM         Direction = C.wxBOTTOM
	NORTH          Direction = C.wxNORTH
	SOUTH          Direction = C.wxSOUTH
	WEST           Direction = C.wxWEST
	EAST           Direction = C.wxEAST
	ALL            Direction = C.wxALL
	DIRECTION_MASK Direction = C.wxDIRECTION_MASK
)

type Stretch int

var (
	STRETCH_NOT Stretch = C.wxSTRETCH_NOT
	SHRINK      Stretch = C.wxSHRINK
	GROW        Stretch = C.wxGROW
	EXPAND      Stretch = C.wxEXPAND
	SHAPED      Stretch = C.wxSHAPED
	TILE        Stretch = C.wxTILE

	STRETCH_MASK Stretch = C.wxSTRETCH_MASK
)

type SizerFlagBits int

var (
	FIXED_MINSIZE                SizerFlagBits = C.wxFIXED_MINSIZE
	RESERVE_SPACE_EVEN_IF_HIDDEN SizerFlagBits = C.wxRESERVE_SPACE_EVEN_IF_HIDDEN
	SIZER_FLAG_BITS_MASK         SizerFlagBits = C.wxSIZER_FLAG_BITS_MASK
)

var DefaultBorder = C.wxSizerFlag_GetDefaultBorder()

type SizerFlags C.SizerFlags

func NewSizerFlags(proportion int) *SizerFlags {
	f := &SizerFlags{}
	C.wxSizerFlags_Init((*C.SizerFlags)(f), C.int(proportion))
	return f
}

func (f *SizerFlags) Align(alignment Aliment) *SizerFlags {
	C.wxSizerFlags_Align((*C.SizerFlags)(f), C.int(alignment))
	return f
}

func (f *SizerFlags) Border(direction Direction, borderinpixels int) *SizerFlags {
	C.wxSizerFlags_Border((*C.SizerFlags)(f), C.int(direction), C.int(borderinpixels))
	return f
}

func (f *SizerFlags) Bottom() *SizerFlags {
	C.wxSizerFlags_Bottom((*C.SizerFlags)(f))
	return f
}

func (f *SizerFlags) Center() *SizerFlags {
	C.wxSizerFlags_Center((*C.SizerFlags)(f))
	return f
}

func (f *SizerFlags) DoubleBorder(direction Direction) *SizerFlags {
	C.wxSizerFlags_DoubleBorder((*C.SizerFlags)(f), C.int(direction))
	return f
}

func (f *SizerFlags) Expand() *SizerFlags {
	C.wxSizerFlags_Expand((*C.SizerFlags)(f))
	return f
}

func (f *SizerFlags) FixedMinSize() *SizerFlags {
	C.wxSizerFlags_FixedMinSize((*C.SizerFlags)(f))
	return f
}

func (f *SizerFlags) ReserveSpaceEvenIfHidden() *SizerFlags {
	C.wxSizerFlags_ReserveSpaceEvenIfHidden((*C.SizerFlags)(f))
	return f
}

func (f *SizerFlags) Left() *SizerFlags {
	C.wxSizerFlags_Left((*C.SizerFlags)(f))
	return f
}

func (f *SizerFlags) Proportion(proportion int) *SizerFlags {
	C.wxSizerFlags_Proportion((*C.SizerFlags)(f), C.int(proportion))
	return f
}

func (f *SizerFlags) Right() *SizerFlags {
	C.wxSizerFlags_Right((*C.SizerFlags)(f))
	return f
}

func (f *SizerFlags) Shaped() *SizerFlags {
	C.wxSizerFlags_Shaped((*C.SizerFlags)(f))
	return f
}

func (f *SizerFlags) Top() *SizerFlags {
	C.wxSizerFlags_Top((*C.SizerFlags)(f))
	return f
}

func (f *SizerFlags) TripleBorder(direction Direction) *SizerFlags {
	C.wxSizerFlags_TripleBorder((*C.SizerFlags)(f), C.int(direction))
	return f
}
