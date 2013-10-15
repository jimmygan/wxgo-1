#include <wx/event.h>
#include "wxutil.h"
#include "threadevent.h"

#define __PTR ((wxThreadEvent*)(p))

WxObjectPtr	wxThreadEvent_New() {
	return new wxThreadEvent;
}
WxObjectPtr	wxThreadEvent_New2(int eventType, int id) {
	return new wxThreadEvent(eventType, id);
}
int		wxThreadEvent_GetEventCategory(WxObjectPtr p) {
	return __PTR->GetEventCategory();
}
void	wxThreadEvent_SetPayload(WxObjectPtr p, void* payload) {
	__PTR->SetPayload(payload);
}
void*	wxThreadEvent_GetPayload(WxObjectPtr p) {
	return __PTR->GetPayload<void*>();
}
long	wxThreadEvent_GetExtraLong(WxObjectPtr p) {
	return __PTR->GetExtraLong();
}
int		wxThreadEvent_GetInt(WxObjectPtr p) {
	return __PTR->GetInt();
}
String	wxThreadEvent_GetString(WxObjectPtr p) {
	return NewGoString(__PTR->GetString());
}
void	wxThreadEvent_SetExtraLong(WxObjectPtr p, long l) {
	__PTR->SetExtraLong(l);
}
void	wxThreadEvent_SetInt(WxObjectPtr p, int n) {
	__PTR->SetInt(n);
}
void 	wxThreadEvent_SetString(WxObjectPtr p, String str) {
	__PTR->SetString(NewWxString(str));
}