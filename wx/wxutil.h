#include <wx/string.h>
#include <wx/gdicmn.h>

extern "C" {
#include "_cgo_export.h"
#include "cgoutil.h"
}

// Check sizeof(A) equals sizeof(B) at compile time.
#define COMPILE_CHECK_SIZE_EQUAL(A,B) char ___sizeof_##A##_does_not_equal_to_sizeof_##B##___[1/(sizeof(A)==sizeof(B))]

wxString NewWxString(String handle);
String NewGoString(const wxString& str);

