#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr	wxWindow_New(WxObjectPtr parent, int id, Point pos, Size size, long style, StringHandle name);
void		wxWindow_SetSizeFlag(WxObjectPtr p, int x, int y, int width, int height, int sizeFlags);
void		wxWindow_Show(WxObjectPtr p, BOOL show);
BOOL		wxWindow_AcceptsFocus(WxObjectPtr p);
BOOL		wxWindow_AcceptsFocusFromKeyboard(WxObjectPtr p);
BOOL		wxWindow_AcceptsFocusRecursively(WxObjectPtr p);
BOOL		wxWindow_IsFocusable(WxObjectPtr p);
BOOL		wxWindow_CanAcceptFocus(WxObjectPtr p);
BOOL		wxWindow_CanAcceptFocusFromKeyboard(WxObjectPtr p);
BOOL		wxWindow_HasFocus(WxObjectPtr p);
void		wxWindow_SetCanFocus(WxObjectPtr p, BOOL can);
void		wxWindow_SetFocus(WxObjectPtr p);
void		wxWindow_SetFocusFromKbd(WxObjectPtr p);
void		wxWindow_AddChild(WxObjectPtr p, WxObjectPtr child);
BOOL		wxWindow_DestroyChildren(WxObjectPtr p);
WxObjectPtr	wxWindow_FindWindow(WxObjectPtr p, long id);
WxObjectPtr	wxWindow_FindWindowByName(WxObjectPtr p, StringHandle name);
// GetChildren hack
int			wxWindow_GetChildrenCount(WxObjectPtr p);
// GetChildren hack
WxObjectPtr	wxWindow_GetChild(WxObjectPtr p, int index);
void		wxWindow_RemoveChild(WxObjectPtr p, void* child);
WxObjectPtr wxWindow_GetGrandParent(WxObjectPtr p);
WxObjectPtr	wxWindow_GetNextSibling(WxObjectPtr p);
WxObjectPtr	wxWindow_GetParent(WxObjectPtr p);
WxObjectPtr	wxWindow_GetPrevSibling(WxObjectPtr p);
BOOL		wxWindow_IsDescendant(WxObjectPtr p,WxObjectPtr win);
BOOL		wxWindow_Reparent(WxObjectPtr p, WxObjectPtr newParent);

Size		wxWindow_GetSize(WxObjectPtr p);
void		wxWindow_SetSize(WxObjectPtr p, Size size);

BOOL		wxWindow_PopupMenu(WxObjectPtr p, WxObjectPtr menu, Point pos);



#ifdef __cplusplus
}
#endif