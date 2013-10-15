#include <wx/button.h>
#include "button.h"
#include "wxutil.h"

#define __PTR ((wxEvent*)(p))


WxObjectPtr wxButton_New(WxObjectPtr parent, int id, String label, Point pos, Size size, long style, WxObjectPtr validator, String name) {
	return new wxButton((wxWindow*)parent, id, NewWxString(label), ToWxPoint(pos), ToWxSize(size), style, validator ? *(wxValidator*)validator : wxDefaultValidator, NewWxString(name));
}