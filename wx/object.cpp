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
StringHandle wxObject_GetClassName(WxObjectPtr p) {
	const wxClassInfo* info = ((wxObject*)p)->GetClassInfo();
	if(info == NULL) {
		return NewGoString("");
	}
	return NewGoString(wxString(info->GetClassName()));
}

StringHandle wxObject_GetBaseClassName1(StringHandle className){
	wxClassInfo* base = wxClassInfo::FindClass(NewWxString(className));
	if(!base) {
		return NULL;
	}
	return NewGoString(base->GetBaseClassName1());
}
StringHandle wxObject_GetBaseClassName2(StringHandle className) {
	wxClassInfo* base = wxClassInfo::FindClass(NewWxString(className));
	if(!base) {
		return NULL;
	}
	return NewGoString(base->GetBaseClassName2());
}