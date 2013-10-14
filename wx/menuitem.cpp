#include <wx/menu.h>
#include "menuitem.h"
#include "wxutil.h"

#define __PTR ((wxMenuItem*)(p))

WxObjectPtr		wxMenuItem_New(WxObjectPtr parentMenu, int id, StringHandle text, StringHandle helpString, int kind, WxObjectPtr subMenu) {
	return new wxMenuItem((wxMenu*)parentMenu, id, NewWxString(text), NewWxString(helpString), (wxItemKind)kind, (wxMenu*)subMenu);
}
StringHandle	wxMenuItem_GetItemLabelText(WxObjectPtr p) {
	return NewGoString(__PTR->GetItemLabelText());
}