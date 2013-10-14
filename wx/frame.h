#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr wxFrame_New(WxObjectPtr parent, int id, StringHandle title);
void 		wxFrame_Center(WxObjectPtr p);

void 		wxFrame_SetMenuBar(WxObjectPtr p, WxObjectPtr menubar);
WxObjectPtr wxFrame_GetMenuBar(WxObjectPtr p);

#ifdef __cplusplus
}
#endif