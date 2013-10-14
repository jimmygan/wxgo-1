#include <wx/event.h>
#include "menuevent.h"

#define __PTR ((wxMenuEvent*)(p))

WxObjectPtr wxMenuEvent_New(int type, int id, WxObjectPtr menu) {
	return new wxMenuEvent(type, id, (wxMenu*)menu);
}

WxObjectPtr wxMenuEvent_GetMenu(WxObjectPtr p) {
	return __PTR->GetMenu();
}

int wxMenuEvent_GetMenuId(WxObjectPtr p) {
	return __PTR->GetMenuId();
}

BOOL wxMenuEvent_IsPopup(WxObjectPtr p) {
	return __PTR->IsPopup();
}