#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr		wxMenuItem_New(WxObjectPtr parentMenu, int id, StringHandle text, StringHandle helpString, int kind, WxObjectPtr subMenu);
StringHandle	wxMenuItem_GetItemLabelText(WxObjectPtr p);

#ifdef __cplusplus
}
#endif