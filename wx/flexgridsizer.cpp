#include <wx/sizer.h>
#include "flexgridsizer.h"

int Get_wxFLEX_GROWMODE_NONE() {
	return wxFLEX_GROWMODE_NONE;
}
int Get_wxFLEX_GROWMODE_SPECIFIED() {
	return wxFLEX_GROWMODE_SPECIFIED;
}
int Get_wxFLEX_GROWMODE_ALL() {
	return wxFLEX_GROWMODE_ALL;
}

#define __PTR ((wxFlexGridSizer*)(p))

WxObjectPtr	wxFlexGridSizer_New(int rows, int cols, int vgap, int hgap) {
	return new wxFlexGridSizer(rows, cols, vgap, hgap);
}
void		wxFlexGridSizer_AddGrowableCol(WxObjectPtr p, long index, int proportion) {
	__PTR->AddGrowableCol(index, proportion);
}
void		wxFlexGridSizer_AddGrowableRow(WxObjectPtr p, long index, int proportion) {
	__PTR->AddGrowableRow(index, proportion);
}
int			wxFlexGridSizer_GetFlexibleDirection(WxObjectPtr p) {
	return __PTR->GetFlexibleDirection();
}
int			wxFlexGridSizer_GetNonFlexibleGrowMode(WxObjectPtr p) {
	return __PTR->GetNonFlexibleGrowMode();
}
BOOL		wxFlexGridSizer_IsColGrowable(WxObjectPtr p, long index) {
	return __PTR->IsColGrowable(index);
}
BOOL		wxFlexGridSizer_IsRowGrowable(WxObjectPtr p, long index) {
	return __PTR->IsRowGrowable(index);
}
void		wxFlexGridSizer_RemoveGrowableCol(WxObjectPtr p, long index) {
	__PTR->RemoveGrowableCol(index);
}
void		wxFlexGridSizer_RemoveGrowableRow(WxObjectPtr p, long index) {
	__PTR->RemoveGrowableRow(index);
}
void		wxFlexGridSizer_SetFlexibleDirection(WxObjectPtr p, int direction) {
	__PTR->SetFlexibleDirection(direction);
}
void		wxFlexGridSizer_SetNonFlexibleGrowMode(WxObjectPtr p, int mode) {
	__PTR->SetNonFlexibleGrowMode((wxFlexSizerGrowMode)mode);
}