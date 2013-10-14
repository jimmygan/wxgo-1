#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr wxMenuEvent_New(int type, int id, WxObjectPtr menu);
WxObjectPtr wxMenuEvent_GetMenu(WxObjectPtr p);
int wxMenuEvent_GetMenuId(WxObjectPtr p);
BOOL wxMenuEvent_IsPopup(WxObjectPtr p);

#ifdef __cplusplus
}
#endif