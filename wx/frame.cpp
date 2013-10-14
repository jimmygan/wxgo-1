#include <wx/frame.h>
#include <wx/menu.h>
#include "wxutil.h"
#include "frame.h"

#define __PTR ((wxFrame*)(p))

WxObjectPtr wxFrame_New(WxObjectPtr parent, int id, StringHandle title) {
	return new wxFrame((wxWindow*)parent, id, NewWxString(title));
}
void wxFrame_Center(WxObjectPtr p) {
	__PTR->Center();
}
void wxFrame_SetMenuBar(WxObjectPtr p, WxObjectPtr menubar) {
	__PTR->SetMenuBar((wxMenuBar*)menubar);
}
WxObjectPtr wxFrame_GetMenuBar(WxObjectPtr p) {
	return __PTR->GetMenuBar();
}