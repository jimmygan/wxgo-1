#include <wx/menu.h>
#include "menuitem.h"
#include "wxutil.h"

#define __PTR ((wxMenuItem*)(p))

WxObjectPtr		wxMenuItem_New(WxObjectPtr parentMenu, int id, String text, String helpString, int kind, WxObjectPtr subMenu) {
	return new wxMenuItem((wxMenu*)parentMenu, id, NewWxString(text), NewWxString(helpString), (wxItemKind)kind, (wxMenu*)subMenu);
}
String	wxMenuItem_GetItemLabelText(WxObjectPtr p) {
	return NewGoString(__PTR->GetItemLabelText());
}