#include <wx/sizer.h>
#include "gridsizer.h"

#define __PTR ((wxGridSizer*)(p))

WxObjectPtr wxGridSizer_New(int rows, int cols, int vgap, int hgap) {
	return new wxGridSizer(rows, cols, vgap, hgap);
}
int			wxGridSizer_GetCols(WxObjectPtr p) {
	return __PTR->GetCols();
}
int			wxGridSizer_GetRows(WxObjectPtr p) {
	return __PTR->GetRows();
}
int			wxGridSizer_GetEffectiveColsCount(WxObjectPtr p) {
	return __PTR->GetEffectiveColsCount();
}
int			wxGridSizer_GetEffectiveRowsCount(WxObjectPtr p) {
	return __PTR->GetEffectiveRowsCount();
}
int			wxGridSizer_GetHGap(WxObjectPtr p) {
	return __PTR->GetHGap();
}
int			wxGridSizer_GetVGap(WxObjectPtr p) {
	return __PTR->GetVGap();
}
void		wxGridSizer_SetCols(WxObjectPtr p, int cols) {
	__PTR->SetCols(cols);
}
void		wxGridSizer_SetRows(WxObjectPtr p, int rows) {
	__PTR->SetRows(rows);
}
void		wxGridSizer_SetHGap(WxObjectPtr p, int hgap) {
	__PTR->SetHGap(hgap);
}
void		wxGridSizer_SetVGap(WxObjectPtr p, int vgap) {
	__PTR->SetVGap(vgap);
}
