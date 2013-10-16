#include <wx/control.h>
#include "control.h"
#include "wxutil.h"

#define __PTR ((wxControl*)(p))

WxObjectPtr wxControl_New(WxObjectPtr parent, int id, Point* pos, Size* size, long style, WxObjectPtr validator, String name) {
	return new wxControl((wxWindow*)parent, id, *((wxPoint*)pos), *((wxSize*)size), style, *((wxValidator*)validator), NewWxString(name));
}
void wxControl_Command(WxObjectPtr p, WxObjectPtr event) {
	__PTR->Command(*((wxCommandEvent*)event));
}
String wxControl_GetLabel(WxObjectPtr p) {
	return NewGoString(__PTR->GetLabel());
}
String wxControl_GetLabelText(WxObjectPtr p) {
	return NewGoString(__PTR->GetLabelText());
}
Size wxControl_GetSizeFromTextSize(WxObjectPtr p, Size* size) {
	wxSize sizeVar = __PTR->GetSizeFromTextSize(*((wxSize*)size));
	return (*(Size*)&sizeVar);
}
void wxControl_SetLabelText(WxObjectPtr p,String text) {
	__PTR->SetLabelText(NewWxString(text));
}
void wxControl_SetLabelMarkup(WxObjectPtr p, String markup) {
	__PTR->SetLabelMarkup(NewWxString(markup));
}
