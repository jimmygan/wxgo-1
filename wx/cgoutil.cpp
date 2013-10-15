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

wxSize ToWxSize(const Size& size) {
	return wxSize(size.Width, size.Height);
}
wxPoint ToWxPoint(const Point& point) {
	return wxPoint(point.Y, point.X);
}
Point ToPoint(const wxPoint& point) {
	Point pt;
	pt.X = point.x;
	pt.Y = point.y;
	return pt;
}
Size ToSize(const wxSize& size) {
	Size s;
	s.Width = size.GetWidth();
	s.Height = size.GetHeight();
	return s;
}
wxRect ToWxRect(const Rect& rect) {
	return wxRect(rect.X, rect.Y, rect.Width, rect.Height);
}
Rect ToRect(const wxRect& rect) {
	Rect rc;
	rc.X = rect.x;
	rc.Y = rect.y;
	rc.Width = rect.width;
	rc.Height = rect.height;
	return rc;
}
