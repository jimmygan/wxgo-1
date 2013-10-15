#include <wx/menu.h>
#include "wxutil.h"
#include "menu.h"

#define __PTR ((wxMenu*)(p))

WxObjectPtr	wxMenu_New(String title, long style) {
	return new wxMenu(NewWxString(title), style);	
}
WxObjectPtr wxMenu_Append (WxObjectPtr p, int id, String item, String helpString, int kind) {
	return __PTR->Append(id, NewWxString(item), NewWxString(helpString), (wxItemKind)kind);
}