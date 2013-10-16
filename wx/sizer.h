#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"
#include "sizerflags.h"
#include "gdidim.h"

WxObjectPtr wxSizer_AddWindow(WxObjectPtr p, WxObjectPtr window, SizerFlags* sizerFlags);
WxObjectPtr wxSizer_AddWindowWithUserData(WxObjectPtr p, WxObjectPtr window, int proportion, int flag, int border, WxObjectPtr userData);
WxObjectPtr wxSizer_AddSizer(WxObjectPtr p, WxObjectPtr sizer, SizerFlags* flags);
WxObjectPtr wxSizer_AddSizerWithUserData(WxObjectPtr p, WxObjectPtr sizer, int proportion, int flag, int border, WxObjectPtr userData);
WxObjectPtr wxSizer_AddNoStretchSpacer(WxObjectPtr p, int size);
WxObjectPtr wxSizer_AddSpacerWithUserData(WxObjectPtr p, int width, int height, int proportion, int flag, int border, WxObjectPtr userData);
WxObjectPtr wxSizer_AddStretchSpacer(WxObjectPtr p, int proportion);
WxObjectPtr wxSizer_Add(WxObjectPtr p, WxObjectPtr item);
void		wxSizer_Clear(WxObjectPtr p, BOOL deleteWindows);
Size		wxSizer_ComputeFittingClientSize(WxObjectPtr p, WxObjectPtr window);
Size		wxSizer_ComputeFittingWindowSize(WxObjectPtr p, WxObjectPtr window);
BOOL		wxSizer_DetachWindow(WxObjectPtr p, WxObjectPtr window);
BOOL		wxSizer_DetachSizer(WxObjectPtr p, WxObjectPtr sizer);
BOOL		wxSizer_Detach(WxObjectPtr p, int index);
Size		wxSizer_Fit(WxObjectPtr p, WxObjectPtr window);
void		wxSizer_FitInside(WxObjectPtr p, WxObjectPtr window);
BOOL		wxSizer_InformFirstDirection(WxObjectPtr p, int direction, int size, int availOtherDir);
WxObjectPtr	wxSizer_GetContainingWindow(WxObjectPtr p);
void		wxSizer_SetContainingWindow(WxObjectPtr p, WxObjectPtr window);
long		wxSizer_GetItemCount(WxObjectPtr p);
WxObjectPtr	wxSizer_GetItem(WxObjectPtr p, long index);
WxObjectPtr	wxSizer_GetItemByWindow(WxObjectPtr p, WxObjectPtr window, BOOL recursive);
WxObjectPtr	wxSizer_GetItemBySizer(WxObjectPtr p, WxObjectPtr sizer, BOOL recursive);
WxObjectPtr	wxSizer_GetItemById(WxObjectPtr p, int id, BOOL recursive);
Size		wxSizer_GetMinSize(WxObjectPtr p);
Point		wxSizer_GetPosition(WxObjectPtr p);
Size		wxSizer_GetSize(WxObjectPtr p);
BOOL		wxSizer_HideWindow(WxObjectPtr p, WxObjectPtr window, BOOL recursive);
BOOL		wxSizer_HideSizer(WxObjectPtr p, WxObjectPtr sizer, BOOL recursive);
BOOL		wxSizer_Hide(WxObjectPtr p, long index);
WxObjectPtr	wxSizer_InsertWindow(WxObjectPtr p, long index, WxObjectPtr window, SizerFlags* flags);
WxObjectPtr	wxSizer_InsertWindowWithUserData(WxObjectPtr p, long index, WxObjectPtr window, int proportion, int flag, int border, WxObjectPtr userdata);
WxObjectPtr	wxSizer_InsertSizer(WxObjectPtr p, long index, WxObjectPtr sizer, SizerFlags* flags);
WxObjectPtr	wxSizer_InsertSizerWithUserData(WxObjectPtr p, long index, WxObjectPtr sizer, int proportion, int flag, int border, WxObjectPtr userdata);
WxObjectPtr	wxSizer_InsertNoStretchSpacer(WxObjectPtr p, long index, int size);
WxObjectPtr	wxSizer_InsertSpacerWithUserData(WxObjectPtr p, long index, int width, int height, int proportion, int flag, int border, WxObjectPtr userData);
WxObjectPtr	wxSizer_InsertStretchSpacer(WxObjectPtr p, long index, int proportion);
WxObjectPtr	wxSizer_Insert(WxObjectPtr p, long index, WxObjectPtr item);
BOOL		wxSizer_IsEmpty(WxObjectPtr p);
BOOL		wxSizer_IsWindowShown(WxObjectPtr p, WxObjectPtr window);
BOOL		wxSizer_IsSizerShown(WxObjectPtr p, WxObjectPtr sizer);
void		wxSizer_Layout(WxObjectPtr p);
BOOL		wxSizer_RemoveSizer(WxObjectPtr p, WxObjectPtr sizer);
BOOL		wxSizer_Remove(WxObjectPtr p, int index);
BOOL		wxSizer_ReplaceWindow(WxObjectPtr p, WxObjectPtr oldWindow, WxObjectPtr newWindow, BOOL recursive);
BOOL		wxSizer_ReplaceSizer(WxObjectPtr p, WxObjectPtr oldSizer, WxObjectPtr newSizer, BOOL recursive);
BOOL		wxSizer_Replace(WxObjectPtr p, long index, WxObjectPtr newItem);
void		wxSizer_SetDimension(WxObjectPtr p, Point* pos, Size* size);
void		wxSizer_SetMinSize(WxObjectPtr p, Size* size);
void		wxSizer_SetSizeHints(WxObjectPtr p, WxObjectPtr window);
BOOL		wxSizer_ShowWindow(WxObjectPtr p, WxObjectPtr window, BOOL show, BOOL recursive);
BOOL		wxSizer_ShowSizer(WxObjectPtr p, WxObjectPtr sizer, BOOL show, BOOL recursive);
BOOL		wxSizer_Show(WxObjectPtr p, long index, BOOL show);
void		wxSizer_ShowAll(WxObjectPtr p, BOOL show);
long		wxSizer_GetChildrenCount(WxObjectPtr p);
WxObjectPtr	wxSizer_GetChild(long index);


#ifdef __cplusplus
}
#endif