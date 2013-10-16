#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

int Get_wxFLEX_GROWMODE_NONE();
int Get_wxFLEX_GROWMODE_SPECIFIED();
int Get_wxFLEX_GROWMODE_ALL();

WxObjectPtr	wxFlexGridSizer_New(int rows, int cols, int vgap, int hgap);
void		wxFlexGridSizer_AddGrowableCol(WxObjectPtr p, long index, int proportion);
void		wxFlexGridSizer_AddGrowableRow(WxObjectPtr p, long index, int proportion);
int			wxFlexGridSizer_GetFlexibleDirection(WxObjectPtr p);
int			wxFlexGridSizer_GetNonFlexibleGrowMode(WxObjectPtr p);
BOOL		wxFlexGridSizer_IsColGrowable(WxObjectPtr p, long index);
BOOL		wxFlexGridSizer_IsRowGrowable(WxObjectPtr p, long index);
void		wxFlexGridSizer_RemoveGrowableCol(WxObjectPtr p, long index);
void		wxFlexGridSizer_RemoveGrowableRow(WxObjectPtr p, long index);
void		wxFlexGridSizer_SetFlexibleDirection(WxObjectPtr p, int direction);
void		wxFlexGridSizer_SetNonFlexibleGrowMode(WxObjectPtr p, int mode);

#ifdef __cplusplus
}
#endif