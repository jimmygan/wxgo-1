#include <wx/gdicmn.h>
#include "gdidim.h"
#include "wxutil.h"

COMPILE_CHECK_SIZE_EQUAL(wxSize, Size);
COMPILE_CHECK_SIZE_EQUAL(wxPoint, Point);
COMPILE_CHECK_SIZE_EQUAL(wxRect, Rect);

Size Get_wxDefaultSize() {
	return *((Size*)&wxDefaultSize);
}

Point Get_wxDefaultPosition() {
	return *((Point*)&wxDefaultPosition);
}