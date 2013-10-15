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
WxObjectPtr wxSizer_AddSizerItem(WxObjectPtr p, WxObjectPtr item) {
	return __PTR->Add((wxSizerItem*)item);
}