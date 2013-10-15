#include <wx/sizer.h>
#include "sizerflags.h"
#include "wxutil.h"

COMPILE_CHECK_SIZE_EQUAL(wxSizerFlags, SizerFlags);

#define __PTR ((wxSizerFlags*)(p))

void wxSizerFlags_Init(SizerFlags* p, int proportion) {
	new(p)wxSizerFlags(proportion); // Placement new.
}

void wxSizerFlags_Align(SizerFlags* p, int alignment) {
	__PTR->Align(alignment);
}
void wxSizerFlags_Border(SizerFlags* p, int direction, int borderinpixels) {
	__PTR->Border(direction, borderinpixels);
}
void wxSizerFlags_Bottom(SizerFlags* p) {
	__PTR->Bottom();
}
void wxSizerFlags_Center(SizerFlags* p) {
	__PTR->Center();
}
void wxSizerFlags_DoubleBorder(SizerFlags* p, int direction) {
	__PTR->DoubleBorder(direction);
}
void wxSizerFlags_Expand(SizerFlags* p) {
	__PTR->Expand();
}
void wxSizerFlags_FixedMinSize(SizerFlags* p) {
	__PTR->FixedMinSize();
}
void wxSizerFlags_ReserveSpaceEvenIfHidden(SizerFlags* p) {
	__PTR->ReserveSpaceEvenIfHidden();
}
void wxSizerFlags_Left(SizerFlags* p) {
	__PTR->Left();
}
void wxSizerFlags_Proportion(SizerFlags* p, int proportion) {
	__PTR->Proportion(proportion);
}
void wxSizerFlags_Right(SizerFlags* p) {
	__PTR->Right();
}
void wxSizerFlags_Shaped(SizerFlags* p) {
	__PTR->Shaped();
}
void wxSizerFlags_Top(SizerFlags* p) {
	__PTR->Top();
}
void wxSizerFlags_TripleBorder(SizerFlags* p, int direction) {
	__PTR->TripleBorder(direction);
}

int wxSizerFlag_GetDefaultBorder() {
	return wxSizerFlags::GetDefaultBorder();
}