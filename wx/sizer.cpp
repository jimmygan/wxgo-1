#include <wx/sizer.h>
#include "sizer.h"

#define __PTR ((wxSizer*)(p))


WxObjectPtr wxSizer_AddWindow(WxObjectPtr p, WxObjectPtr window, SizerFlags* sizerFlags) {
	return __PTR->Add((wxWindow*)window, *(wxSizerFlags*)sizerFlags);
}
WxObjectPtr wxSizer_AddWindowWithUserData(WxObjectPtr p, WxObjectPtr window, int proportion, int flag, int border, WxObjectPtr userData) {
	return __PTR->Add((wxWindow*)window, proportion, flag, border, (wxObject*)userData);
}
WxObjectPtr wxSizer_AddSizer(WxObjectPtr p, WxObjectPtr sizer, SizerFlags* flags) {
	return __PTR->Add((wxSizer*)sizer, *(wxSizerFlags*)flags);
}
WxObjectPtr wxSizer_AddSizerWithUserData(WxObjectPtr p, WxObjectPtr sizer, int proportion, int flag, int border, WxObjectPtr userData) {
	return __PTR->Add((wxSizer*)sizer, proportion, flag, border, (wxObject*)userData);
}
WxObjectPtr wxSizer_AddNoStretchSpacer(WxObjectPtr p, int size) {
	return __PTR->AddSpacer(size);
}
WxObjectPtr wxSizer_AddSpacerWithUserData(WxObjectPtr p, int width, int height, int proportion, int flag, int border, WxObjectPtr userData) {
	return __PTR->Add(width, height, proportion, flag, border, (wxObject*)userData);
}
WxObjectPtr wxSizer_AddStretchSpacer(WxObjectPtr p, int proportion) {
	return __PTR->AddStretchSpacer(proportion);
}
WxObjectPtr wxSizer_Add(WxObjectPtr p, WxObjectPtr item) {
	return __PTR->Add((wxSizerItem*)item);
}

void		wxSizer_Clear(WxObjectPtr p, BOOL deleteWindows) {
	__PTR->Clear(deleteWindows);
}
Size		wxSizer_ComputeFittingClientSize(WxObjectPtr p, WxObjectPtr window) {
	wxSize size = __PTR->ComputeFittingClientSize((wxWindow*)window);
	return *((Size*)&size);
}
Size		wxSizer_ComputeFittingWindowSize(WxObjectPtr p, WxObjectPtr window) {
	wxSize size = __PTR->ComputeFittingWindowSize((wxWindow*)window);
	return *((Size*)&size);
}
BOOL		wxSizer_DetachWindow(WxObjectPtr p, WxObjectPtr window){
	return __PTR->Detach((wxWindow*)window);
}
BOOL		wxSizer_DetachSizer(WxObjectPtr p, WxObjectPtr sizer) {
	return __PTR->Detach((wxSizer*)sizer);
}
BOOL		wxSizer_Detach(WxObjectPtr p, int index) {
	return __PTR->Detach(index);
}
Size		wxSizer_Fit(WxObjectPtr p, WxObjectPtr window) {
	wxSize size = __PTR->Fit((wxWindow*)window);
	return *((Size*)&size);
}
void		wxSizer_FitInside(WxObjectPtr p, WxObjectPtr window) {
	__PTR->FitInside((wxWindow*)window);
}
BOOL		wxSizer_InformFirstDirection(WxObjectPtr p, int direction, int size, int availOtherDir) {
	return __PTR->InformFirstDirection(direction, size, availOtherDir);
}
WxObjectPtr	wxSizer_GetContainingWindow(WxObjectPtr p) {
	return __PTR->GetContainingWindow();
}
void		wxSizer_SetContainingWindow(WxObjectPtr p, WxObjectPtr window) {
	__PTR->SetContainingWindow((wxWindow*)window);
}
long		wxSizer_GetItemCount(WxObjectPtr p) {
	return __PTR->GetItemCount();
}
WxObjectPtr	wxSizer_GetItem(WxObjectPtr p, long index) {
	return __PTR->GetItem(index);
}
WxObjectPtr	wxSizer_GetItemByWindow(WxObjectPtr p, WxObjectPtr window, BOOL recursive) {
	return __PTR->GetItem((wxWindow*)window, recursive);
}
WxObjectPtr	wxSizer_GetItemBySizer(WxObjectPtr p, WxObjectPtr sizer, BOOL recursive) {
	return __PTR->GetItem((wxSizer*)sizer, recursive);
}
WxObjectPtr	wxSizer_GetItemById(WxObjectPtr p, int id, BOOL recursive) {
	return __PTR->GetItemById(id, (bool)recursive);
}
Size		wxSizer_GetMinSize(WxObjectPtr p) {
	wxSize size = __PTR->GetMinSize();
	return *((Size*)&size);
}
Point		wxSizer_GetPosition(WxObjectPtr p) {
	wxPoint pt = __PTR->GetPosition();
	return *((Point*)&pt);
}
Size		wxSizer_GetSize(WxObjectPtr p) {
	wxSize size = __PTR->GetSize();
	return *((Size*)&size);
}
BOOL		wxSizer_HideWindow(WxObjectPtr p, WxObjectPtr window, BOOL recursive) {
	return __PTR->Hide((wxWindow*)window, recursive);
}
BOOL		wxSizer_HideSizer(WxObjectPtr p, WxObjectPtr sizer, BOOL recursive) {
	return __PTR->Hide((wxSizer*)sizer, recursive);
}
BOOL		wxSizer_Hide(WxObjectPtr p, long index) {
	return __PTR->Hide(index);
}
WxObjectPtr	wxSizer_InsertWindow(WxObjectPtr p, long index, WxObjectPtr window, SizerFlags* flags) {
	return __PTR->Insert(index, (wxWindow*)window, *(wxSizerFlags*)flags);
}
WxObjectPtr	wxSizer_InsertWindowWithUserData(WxObjectPtr p, long index, WxObjectPtr window, int proportion, int flag, int border, WxObjectPtr userData) {
	return __PTR->Insert(index, (wxWindow*)window, proportion, flag, border, (wxObject*)userData);
}
WxObjectPtr	wxSizer_InsertSizer(WxObjectPtr p, long index, WxObjectPtr sizer, SizerFlags* flags) {
	return __PTR->Insert(index, (wxSizer*)sizer, *(wxSizerFlags*)flags);
}
WxObjectPtr	wxSizer_InsertSizerWithUserData(WxObjectPtr p, long index, WxObjectPtr sizer, int proportion, int flag, int border, WxObjectPtr userData) {
	return __PTR->Insert(index, (wxSizer*)sizer, proportion, flag, border, (wxObject*)userData);
}
WxObjectPtr	wxSizer_InsertNoStretchSpacer(WxObjectPtr p, long index, int size) {
	return __PTR->InsertSpacer(index, size);
}
WxObjectPtr	wxSizer_InsertSpacerWithUserData(WxObjectPtr p, long index, int width, int height, int proportion, int flag, int border, WxObjectPtr userData) {
	return __PTR->Insert(index, width, height, proportion, flag, border, (wxObject*)userData);
}
WxObjectPtr	wxSizer_InsertStretchSpacer(WxObjectPtr p, long index, int proportion) {
	return __PTR->InsertStretchSpacer(proportion);
}
WxObjectPtr	wxSizer_Insert(WxObjectPtr p, long index, WxObjectPtr item) {
	return __PTR->Insert(index, (wxSizerItem*)item);
}
BOOL		wxSizer_IsEmpty(WxObjectPtr p) {
	return __PTR->IsEmpty();
}
BOOL		wxSizer_IsWindowShown(WxObjectPtr p, WxObjectPtr window) {
	return ((const wxSizer*)__PTR)->IsShown((wxWindow*)window);
}
BOOL		wxSizer_IsSizerShown(WxObjectPtr p, WxObjectPtr sizer) {
	return __PTR->IsShown((wxSizer*)sizer);
}
void		wxSizer_Layout(WxObjectPtr p) {
	__PTR->Layout();
}
BOOL		wxSizer_RemoveSizer(WxObjectPtr p, WxObjectPtr sizer) {
	return __PTR->Remove((wxSizer*)sizer);
}
BOOL		wxSizer_Remove(WxObjectPtr p, int index) {
	return __PTR->Remove(index);
}
BOOL		wxSizer_ReplaceWindow(WxObjectPtr p, WxObjectPtr oldWindow, WxObjectPtr newWindow, BOOL recursive) {
	return __PTR->Replace((wxWindow*)oldWindow, (wxWindow*)newWindow, recursive);
}
BOOL		wxSizer_ReplaceSizer(WxObjectPtr p, WxObjectPtr oldSizer, WxObjectPtr newSizer, BOOL recursive) {
	return __PTR->Replace((wxSizer*)oldSizer, (wxSizer*)newSizer, recursive);
}
BOOL		wxSizer_Replace(WxObjectPtr p, long index, WxObjectPtr newItem) {
	return __PTR->Replace(index, (wxSizerItem*)newItem);
}
void		wxSizer_SetDimension(WxObjectPtr p, Point* pos, Size* size) {
	return __PTR->SetDimension(*((wxPoint*)pos), *((wxSize*)size));
}
void		wxSizer_SetMinSize(WxObjectPtr p, Size* size) {
	return __PTR->SetMinSize(*((wxSize*)size));
}
void		wxSizer_SetSizeHints(WxObjectPtr p, WxObjectPtr window) {
	__PTR->SetSizeHints((wxWindow*)window);
}
BOOL		wxSizer_ShowWindow(WxObjectPtr p, WxObjectPtr window, BOOL show, BOOL recursive)  {
	return __PTR->Show((wxWindow*)window, show, recursive);
}
BOOL		wxSizer_ShowSizer(WxObjectPtr p, WxObjectPtr sizer, BOOL show, BOOL recursive) {
	return __PTR->Show((wxSizer*)sizer, show, recursive);
}
BOOL		wxSizer_Show(WxObjectPtr p, long index, BOOL show) {
	return __PTR->Show(index, show);
}
void		wxSizer_ShowAll(WxObjectPtr p, BOOL show) {
	return __PTR->ShowItems(show);
}
