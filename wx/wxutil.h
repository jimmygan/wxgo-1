#include <wx/string.h>
#include <wx/gdicmn.h>

extern "C" {
#include "_cgo_export.h"
#include "cgoutil.h"
}

// Check sizeof(A) equals sizeof(B) at compile time.
#define COMPILE_CHECK_SIZE_EQUAL(A,B) char ___sizeof_##A_does_not_equal_to_sizeof_##B___[1/(sizeof(A)==sizeof(B))]

wxString NewWxString(String handle);
String NewGoString(const wxString& str);

wxSize ToWxSize(const Size& size);
wxPoint ToWxPoint(const Point& point);
Size ToSize(const wxSize& size);
Point ToPoint(const wxPoint& point);
wxRect ToWxRect(const Rect& rect);
Rect ToRect(const wxRect& rect);


