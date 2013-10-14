#include <wx/string.h>
#include <wx/gdicmn.h>

extern "C" {
#include "_cgo_export.h"
#include "cgoutil.h"
}

wxString NewWxString(StringHandle handle);
StringHandle NewGoString(const wxString& str);

wxSize toWxSize(const Size& size);
wxPoint ToWxPoint(const Point& point);
Size toSize(const wxSize& size);
Point toPoint(const wxPoint& point);
wxRect toWxRect(const Rect& rect);
Rect toRect(const wxRect& rect);


