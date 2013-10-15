#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr		wxMenuItem_New(WxObjectPtr parentMenu, int id, String text, String helpString, int kind, WxObjectPtr subMenu);
String	wxMenuItem_GetItemLabelText(WxObjectPtr p);

#ifdef __cplusplus
}
#endif