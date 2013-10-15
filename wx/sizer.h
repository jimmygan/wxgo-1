#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"
#include "sizerflags.h"

WxObjectPtr wxSizer_AddWindow(WxObjectPtr p, WxObjectPtr window, SizerFlags* sizerFlags);
WxObjectPtr wxSizer_AddWindowWithUserData(WxObjectPtr p, WxObjectPtr window, int proportion, int flag, int border, WxObjectPtr userData);
WxObjectPtr wxSizer_AddSizer(WxObjectPtr p, WxObjectPtr sizer, SizerFlags* flags);
WxObjectPtr wxSizer_AddSizerWithUserData(WxObjectPtr p, WxObjectPtr sizer, int proportion, int flag, int border, WxObjectPtr userData);
WxObjectPtr wxSizer_AddNoStretchSpacer(WxObjectPtr p, int size);
WxObjectPtr wxSizer_AddSpacerWithUserData(WxObjectPtr p, int width, int height, int proportion, int flag, int border, WxObjectPtr userData);
WxObjectPtr wxSizer_AddStretchSpacer(WxObjectPtr p, int proportion);
WxObjectPtr wxSizer_AddSizerItem(WxObjectPtr p, WxObjectPtr item);

#ifdef __cplusplus
}
#endif