#include "wxutil.h"

wxString NewWxString(String str) {
	// The String is pointer to GoString (in _cgo_export.h) which is the
	// internal representation of go string.
	GoString* pGoString = (GoString*)str;
	return wxString::FromUTF8Unchecked(pGoString->p, pGoString->n);
}

String NewGoString(const wxString& str) {
	// Call exported go function to create a go string.
	const  wxCharBuffer utf8 = str.utf8_str();
	return x_newGoString((char*)utf8.data(), utf8.length());
}
