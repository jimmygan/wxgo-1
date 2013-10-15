#include <wx/sizer.h>
#include "boxsizer.h"
#include "wxutil.h"

#define __PTR ((wxBoxSizer*)(p))

WxObjectPtr wxBoxSizer_New(int orient) {
	return new wxBoxSizer(orient);
}
int wxBoxSizer_GetOrientation(WxObjectPtr p) {
	return __PTR->GetOrientation();
}
void wxBoxSizer_SetOrientation(WxObjectPtr p, int orient) {
	__PTR->SetOrientation(orient);
}