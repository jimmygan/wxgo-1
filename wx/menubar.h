#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr	wxMenuBar_New(long style);
void 		wxMenuBar_Append(WxObjectPtr p, WxObjectPtr menu, StringHandle title);
WxObjectPtr	wxMenuBar_Remove(WxObjectPtr p, long pos);

#ifdef __cplusplus
}
#endif