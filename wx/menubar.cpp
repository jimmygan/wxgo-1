#include <wx/menu.h>
#include "wxutil.h"
#include "menubar.h"

#define __PTR ((wxMenuBar*)(p))

WxObjectPtr	wxMenuBar_New(long style) {
	return new wxMenuBar(style);
}
void 		wxMenuBar_Append(WxObjectPtr p, WxObjectPtr menu, String title) {
	__PTR->Append((wxMenu*)menu, NewWxString(title));
}
WxObjectPtr	wxMenuBar_Remove(WxObjectPtr p, long pos) {
	return __PTR->Remove(pos);
}