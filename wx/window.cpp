#include <wx/window.h>
#include "wxutil.h"
#include "window.h"

#define __PTR ((wxWindow*)(p))

WxObjectPtr wxWindow_New(WxObjectPtr parent, int id, Point* pos, Size* size, long style, String name) {
	return new wxWindow((wxWindow*)parent, id, pos ? *((wxPoint*)pos) : wxDefaultPosition, size ? *((wxSize*)size) : wxDefaultSize, style, NewWxString(name));
}

void wxWindow_SetSizeFlag(WxObjectPtr p, int x, int y, int width, int height, int sizeFlags) {
	__PTR->SetSize(x, y, width, height, sizeFlags);
}

void wxWindow_Show(WxObjectPtr p, BOOL show) {
	__PTR->Show(show);
}
BOOL	wxWindow_AcceptsFocus(WxObjectPtr p) {
	return __PTR->AcceptsFocus();
}
BOOL	wxWindow_AcceptsFocusFromKeyboard(WxObjectPtr p) {
	return __PTR->AcceptsFocusFromKeyboard();
}
BOOL	wxWindow_AcceptsFocusRecursively(WxObjectPtr p) {
	return __PTR->AcceptsFocusRecursively();
}
BOOL	wxWindow_IsFocusable(WxObjectPtr p) {
	return __PTR->IsFocusable();
}
BOOL	wxWindow_CanAcceptFocus(WxObjectPtr p) {
	return __PTR->CanAcceptFocus();
}
BOOL	wxWindow_CanAcceptFocusFromKeyboard(WxObjectPtr p) {
	return __PTR->CanAcceptFocusFromKeyboard();
}
BOOL	wxWindow_HasFocus(WxObjectPtr p) {
	return __PTR->HasFocus();
}
void	wxWindow_SetCanFocus(WxObjectPtr p, BOOL can) {
	__PTR->SetCanFocus(can);
}
void	wxWindow_SetFocus(WxObjectPtr p) {
	__PTR->SetFocus();
}
void	wxWindow_SetFocusFromKbd(WxObjectPtr p) {
	__PTR->SetFocusFromKbd();
}
void	wxWindow_AddChild(WxObjectPtr p, WxObjectPtr child) {
	__PTR->AddChild((wxWindow*)child);
}
BOOL	wxWindow_DestroyChildren(WxObjectPtr p) {
	return __PTR->DestroyChildren();
}
WxObjectPtr	wxWindow_FindWindow(WxObjectPtr p, long id) {
	return __PTR->FindWindow(id);
}
WxObjectPtr	wxWindow_FindWindowByName(WxObjectPtr p, String name) {
	return __PTR->FindWindow(NewWxString(name));
}
// GetChildren hack
int		wxWindow_GetChildrenCount(WxObjectPtr p) {
	return __PTR->GetChildren().GetCount();
}
// GetChildren hack
WxObjectPtr	wxWindow_GetChild(WxObjectPtr p, int index) {
	return __PTR->GetChildren().Item(index);
}
void	wxWindow_RemoveChild(WxObjectPtr p, WxObjectPtr child) {
	__PTR->RemoveChild((wxWindow*)child);
}
WxObjectPtr 	wxWindow_GetGrandParent(WxObjectPtr p) {
	return __PTR->GetGrandParent();
}
void*	wxWindow_GetNextSibling(WxObjectPtr p) {
	return __PTR->GetNextSibling();
}
WxObjectPtr	wxWindow_GetParent(WxObjectPtr p) {
	return __PTR->GetParent();
}
WxObjectPtr	wxWindow_GetPrevSibling(WxObjectPtr p) {
	return __PTR->GetPrevSibling();
}
BOOL	wxWindow_IsDescendant(WxObjectPtr p, WxObjectPtr win) {
	return __PTR->IsDescendant((wxWindowBase*)win);
}
BOOL	wxWindow_Reparent(WxObjectPtr p, WxObjectPtr newParent) {
	return __PTR->Reparent((wxWindow*)newParent);
}
Size		wxWindow_GetSize(WxObjectPtr p) {
	wxSize size = __PTR->GetSize();
	return *((Size*)&size);
}
void		wxWindow_SetSize(WxObjectPtr p, Size* size) {
	__PTR->SetSize(*((wxSize*)size));
}
BOOL		wxWindow_PopupMenu(WxObjectPtr p, WxObjectPtr menu, Point* pos) {
	return __PTR->PopupMenu((wxMenu*)menu, pos ? *((wxPoint*)pos) : wxDefaultPosition);
}
BOOL		wxWindow_Close(WxObjectPtr p, BOOL force) {
	return __PTR->Close(force);
}
void		wxWindow_Destroy(WxObjectPtr p) {
	__PTR->Destroy();
}
BOOL		wxWindow_IsBeingDeleted(WxObjectPtr p) {
	return __PTR->IsBeingDeleted();
}
WxObjectPtr	wxWindow_GetSizer(WxObjectPtr p) {
	return __PTR->GetSizer();
}
void		wxWindow_SetSizer(WxObjectPtr p, WxObjectPtr sizer, BOOL deleteOld) {
	__PTR->SetSizer((wxSizer*)sizer, deleteOld);
}