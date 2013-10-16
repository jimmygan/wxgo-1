#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"
#include "gdidim.h"

WxObjectPtr wxControl_New(WxObjectPtr parent, int id, Point* pos, Size* size, long style, WxObjectPtr validator, String name);
void 		wxControl_Command(WxObjectPtr p, WxObjectPtr event);
String		wxControl_GetLabel(WxObjectPtr p);
String		wxControl_GetLabelText(WxObjectPtr p);
Size		wxControl_GetSizeFromTextSize(WxObjectPtr p, Size* size);
void		wxControl_SetLabelText(WxObjectPtr p, String text);
void		wxControl_SetLabelMarkup(WxObjectPtr p, String markup);

#ifdef __cplusplus
}
#endif