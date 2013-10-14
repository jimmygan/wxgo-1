#include <wx/event.h>
#include "event.h"

#define __PTR ((wxEvent*)(p))


WxObjectPtr wxEvent_GetEventObject(WxObjectPtr p) {
	return __PTR->GetEventObject();
}

int wxEvent_GetEventType(WxObjectPtr p) {
	return __PTR->GetEventType();
}

int wxEvent_GetEventCategory(WxObjectPtr p) {
	return __PTR->GetEventCategory();
}

int wxEvent_GetId(WxObjectPtr p) {
	return __PTR->GetId();
}

WxObjectPtr wxEvent_GetEventUserData(WxObjectPtr p) {
	return __PTR->GetEventUserData();
}

int wxEvent_GetSkipped(WxObjectPtr p) {
	return __PTR->GetSkipped() ? 1 : 0;
}

long wxEvent_GetTimestamp(WxObjectPtr p) {
	return __PTR->GetTimestamp();
}

BOOL wxEvent_IsCommandEvent(WxObjectPtr p) {
	return __PTR->IsCommandEvent();
}

void wxEvent_ResumePropagation(WxObjectPtr p, int propagationLevel) {
	__PTR->ResumePropagation(propagationLevel);
}

void wxEvent_SetEventObject(WxObjectPtr p, WxObjectPtr object) {
	__PTR->SetEventObject((wxObject*)object);
}

void wxEvent_SetEventType(WxObjectPtr p, int type) {
	__PTR->SetEventType(type);
}

void wxEvent_SetId(WxObjectPtr p, int id) {
	__PTR->SetId(id);
}

void wxEvent_SetTimestamp(WxObjectPtr p, long timestamp) {
	__PTR->SetTimestamp(timestamp);
}

BOOL wxEvent_ShouldPropagate(WxObjectPtr p) {
	return __PTR->ShouldPropagate();
}

void wxEvent_Skip(WxObjectPtr p, BOOL skip) {
	__PTR->Skip(skip);
}

void wxEvent_StopPropagation(WxObjectPtr p) {
	__PTR->StopPropagation();
}