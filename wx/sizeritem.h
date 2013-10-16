#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"
#include "gdidim.h"
#include "sizerflags.h"

WxObjectPtr	wxSizerItem_NewSpacer(int width, int height, int proportion, int flag, int border, WxObjectPtr userData);
WxObjectPtr	wxSizerItem_NewWindow(WxObjectPtr window, SizerFlags* flags);
WxObjectPtr	wxSizerItem_NewSizer(WxObjectPtr sizer, SizerFlags* flags);
void		wxSizerItem_AssignWindow(WxObjectPtr p, WxObjectPtr window);
void		wxSizerItem_AssignSizer(WxObjectPtr p, WxObjectPtr sizer);
void		wxSizerItem_AssignSpacer(WxObjectPtr p, Size* size);
void		wxSizerItem_DeleteWindows(WxObjectPtr p);
void		wxSizerItem_DetachSizer(WxObjectPtr p);
int			wxSizerItem_GetBorder(WxObjectPtr p);
int			wxSizerItem_GetFlag(WxObjectPtr p);
int			wxSizerItem_GetId(WxObjectPtr p);
Size		wxSizerItem_GetMinSize(WxObjectPtr p);
void		wxSizerItem_SetMinSize(WxObjectPtr p, Size* size);
Point		wxSizerItem_GetPosition(WxObjectPtr p);
int			wxSizerItem_GetProportion(WxObjectPtr p);
float		wxSizerItem_GetRatio(WxObjectPtr p);
Rect		wxSizerItem_GetRect(WxObjectPtr p);
Size		wxSizerItem_GetSize(WxObjectPtr p);
WxObjectPtr	wxSizerItem_GetSizer(WxObjectPtr p);
Size		wxSizerItem_GetSpacer(WxObjectPtr p);
WxObjectPtr	wxSizerItem_GetUserData(WxObjectPtr p);
WxObjectPtr wxSizerItem_GetWindow(WxObjectPtr p);
BOOL		wxSizerItem_IsShown(WxObjectPtr p);
BOOL		wxSizerItem_IsSizer(WxObjectPtr p);
BOOL		wxSizerItem_IsSpacer(WxObjectPtr p);
BOOL		wxSizerItem_IsWindow(WxObjectPtr p);
void		wxSizerItem_SetBorder(WxObjectPtr p, int border);
void		wxSizerItem_SetDimension(WxObjectPtr, Point* pos, Size* size);
void		wxSizerItem_SetFlag(WxObjectPtr p, int flag);
void		wxSizerItem_SetId(WxObjectPtr p, int id);
void		wxSizerItem_SetInitSize(WxObjectPtr p, Size* size);
void		wxSizerItem_SetProportion(WxObjectPtr p, int proportion);
void		wxSizerItem_SetUserData(WxObjectPtr p, WxObjectPtr userData);
void		wxSizerItem_SetRatio(WxObjectPtr p, float ratio);
void		wxSizerItem_SetRatioSize(WxObjectPtr p, Size* size);




#ifdef __cplusplus
}
#endif