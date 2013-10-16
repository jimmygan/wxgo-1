#include <wx/sizer.h>
#include "sizeritem.h"

#define __PTR ((wxSizerItem*)(p))

WxObjectPtr	wxSizerItem_NewSpacer(int width, int height, int proportion, int flag, int border, WxObjectPtr userData) {
	return new wxSizerItem(width, height, proportion, flag, border, (wxObject*)userData);
}
WxObjectPtr	wxSizerItem_NewWindow(WxObjectPtr window, SizerFlags* flags) {
	return new wxSizerItem((wxWindow*)window, *((wxSizerFlags*)flags));
}
WxObjectPtr	wxSizerItem_NewSizer(WxObjectPtr sizer, SizerFlags* flags) {
	return new wxSizerItem((wxSizer*)sizer, *((wxSizerFlags*)flags));
}
void		wxSizerItem_AssignWindow(WxObjectPtr p, WxObjectPtr window) {
	__PTR->AssignWindow((wxWindow*)window);
}
void		wxSizerItem_AssignSizer(WxObjectPtr p, WxObjectPtr sizer) {
	__PTR->AssignSizer((wxSizer*)sizer);
}
void		wxSizerItem_AssignSpacer(WxObjectPtr p, Size* size) {
	__PTR->AssignSpacer(size->Width, size->Height);
}
void		wxSizerItem_DeleteWindows(WxObjectPtr p) {
	__PTR->DeleteWindows();
}
void		wxSizerItem_DetachSizer(WxObjectPtr p) {
	__PTR->DetachSizer();
}
int			wxSizerItem_GetBorder(WxObjectPtr p) {
	__PTR->GetBorder();
}
int			wxSizerItem_GetFlag(WxObjectPtr p) {
	__PTR->GetFlag();
}
int			wxSizerItem_GetId(WxObjectPtr p) {
	__PTR->GetId();
}
Size		wxSizerItem_GetMinSize(WxObjectPtr p) {
	wxSize size = __PTR->GetMinSize();
	return *((Size*)&size);
}
void		wxSizerItem_SetMinSize(WxObjectPtr p, Size* size) {
	__PTR->SetMinSize(*((wxSize*)size));
}
Point		wxSizerItem_GetPosition(WxObjectPtr p) {
	wxPoint point = __PTR->GetPosition();
	return *((Point*)&point);
}
int			wxSizerItem_GetProportion(WxObjectPtr p) {
	return __PTR->GetProportion();
}
float		wxSizerItem_GetRatio(WxObjectPtr p) {
	return __PTR->GetRatio();
}
Rect		wxSizerItem_GetRect(WxObjectPtr p) {
	wxRect rect = __PTR->GetRect();
	return *((Rect*)&rect);
}
Size		wxSizerItem_GetSize(WxObjectPtr p) {
	wxSize size = __PTR->GetSize();
	return *((Size*)&size);
}
WxObjectPtr	wxSizerItem_GetSizer(WxObjectPtr p) {
	return __PTR->GetSizer();
}
Size		wxSizerItem_GetSpacer(WxObjectPtr p) {
	wxSize size = __PTR->GetSpacer();
	return *((Size*)&size);
}
WxObjectPtr	wxSizerItem_GetUserData(WxObjectPtr p) {
	return __PTR->GetUserData();
}
WxObjectPtr wxSizerItem_GetWindow(WxObjectPtr p) {
	return __PTR->GetWindow();
}
BOOL		wxSizerItem_IsShown(WxObjectPtr p) {
	return __PTR->IsShown();
}
BOOL		wxSizerItem_IsSizer(WxObjectPtr p) {
	return __PTR->IsSizer();
}
BOOL		wxSizerItem_IsSpacer(WxObjectPtr p) {
	return __PTR->IsSpacer();
}
BOOL		wxSizerItem_IsWindow(WxObjectPtr p) {
	return __PTR->IsWindow();
}
void		wxSizerItem_SetBorder(WxObjectPtr p, int border) {
	__PTR->SetBorder(border);
}
void		wxSizerItem_SetDimension(WxObjectPtr p, Point* pos, Size* size) {
	__PTR->SetDimension(*((wxPoint*)pos), *((wxSize*)size));
}
void		wxSizerItem_SetFlag(WxObjectPtr p, int flag) {
	__PTR->SetFlag(flag);
}
void		wxSizerItem_SetId(WxObjectPtr p, int id) {
	__PTR->SetId(id);
}
void		wxSizerItem_SetInitSize(WxObjectPtr p, Size* size) {
	__PTR->SetInitSize(size->Width, size->Height);
}
void		wxSizerItem_SetProportion(WxObjectPtr p, int proportion) {
	__PTR->SetProportion(proportion);
}
void		wxSizerItem_SetUserData(WxObjectPtr p, WxObjectPtr userData) {
	__PTR->SetUserData((wxObject*)userData);
}
void		wxSizerItem_SetRatio(WxObjectPtr p, float ratio) {
	__PTR->SetRatio(ratio);
}
void		wxSizerItem_SetRatioSize(WxObjectPtr p, Size* size) {
	__PTR->SetRatio(*((wxSize*)size));
}
