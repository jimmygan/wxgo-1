#include <wx/variant.h>
#include "wxutil.h"
#include "variant.h"

#define __PTR ((wxVariant*)(p))

WxObjectPtr wxVariant_NewInt64(long long n) {
	return new wxVariant(wxLongLong(n));
}

WxObjectPtr wxVariant_NewString(String str) {
	return new wxVariant(NewWxString(str));
}
long long wxVariant_GetInt64(WxObjectPtr p) {
	return __PTR->GetLongLong().GetValue();
}
String wxVariant_GetString(WxObjectPtr p) {
	return NewGoString(__PTR->GetString());
}
