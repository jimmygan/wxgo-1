#include <wx/object.h>
#include <stdlib.h>
#include "object.h"
#include "wxutil.h"

WxObjectPtr wxObject_New() {
	return new wxObject;
}

void wxObject_Delete(WxObjectPtr p) {
	delete (wxObject*)p;
}

// Get the class name of a wxObject*.
String wxObject_GetClassName(WxObjectPtr p) {
	const wxClassInfo* info = ((wxObject*)p)->GetClassInfo();
	if(info == NULL) {
		return NewGoString("");
	}
	return NewGoString(wxString(info->GetClassName()));
}

String wxObject_GetBaseClassName1(String className){
	wxClassInfo* base = wxClassInfo::FindClass(NewWxString(className));
	if(!base) {
		return NULL;
	}
	return NewGoString(base->GetBaseClassName1());
}
String wxObject_GetBaseClassName2(String className) {
	wxClassInfo* base = wxClassInfo::FindClass(NewWxString(className));
	if(!base) {
		return NULL;
	}
	return NewGoString(base->GetBaseClassName2());
}