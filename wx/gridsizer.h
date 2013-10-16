#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr wxGridSizer_New(int rows, int cols, int vgap, int hgap);
int			wxGridSizer_GetCols(WxObjectPtr p);
int			wxGridSizer_GetRows(WxObjectPtr p);
int			wxGridSizer_GetEffectiveColsCount(WxObjectPtr p);
int			wxGridSizer_GetEffectiveRowsCount(WxObjectPtr p);
int			wxGridSizer_GetHGap(WxObjectPtr p);
int			wxGridSizer_GetVGap(WxObjectPtr p);
void		wxGridSizer_SetCols(WxObjectPtr p, int cols);
void		wxGridSizer_SetRows(WxObjectPtr p, int rows);
void		wxGridSizer_SetHGap(WxObjectPtr p, int hgap);
void		wxGridSizer_SetVGap(WxObjectPtr p, int vgap);

#ifdef __cplusplus
}
#endif