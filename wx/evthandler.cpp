#include <wx/wx.h>
#include "evthandler.h"
#include "wxutil.h"
#include <stdio.h>

class eventFunctor {
public:
	eventFunctor(void* pF): m_pGoFunc(pF){}
	const void* m_pGoFunc;
public:
	void operator() (wxEvent& event) {
		callEventFunc((void*)m_pGoFunc, &event);
	}
	
	bool operator==(const eventFunctor& other) {
		return m_pGoFunc == other.m_pGoFunc;
	}
};


#define __PTR ((wxEvtHandler*)(p))

WxObjectPtr	wxEvtHandler_New() {
	return new wxEvtHandler;
}
void	wxEvtHandler_QueueEvent(WxObjectPtr p, void* event) {
	__PTR->QueueEvent((wxEvent*)event);
}
void	wxEvtHandler_AddPendingEvent (WxObjectPtr p, void* event) {
	__PTR->AddPendingEvent(*(wxEvent*)event);
}
BOOL	wxEvtHandler_ProcessEvent(WxObjectPtr p, WxObjectPtr event) {
	return __PTR->ProcessEvent(*(wxEvent*)event);
}
BOOL	wxEvtHandler_ProcessEventLocally(WxObjectPtr p, WxObjectPtr event) {
	return __PTR->ProcessEventLocally(*(wxEvent*)event);
}
BOOL	wxEvtHandler_SafelyProcessEvent(WxObjectPtr p, WxObjectPtr event) {
	return __PTR->SafelyProcessEvent(*(wxEvent*)event);
}
void	wxEvtHandler_ProcessPendingEvents(WxObjectPtr p) {
	__PTR->ProcessPendingEvents();
}
void	wxEvtHandler_DeletePendingEvents(WxObjectPtr p) {
	__PTR->DeletePendingEvents();
}
void	wxEvtHandler_Bind(WxObjectPtr p, void* func, int eventType, int id, int lastId, WxObjectPtr userData) {
	__PTR->Bind(wxEventTypeTag<wxEvent>(eventType), eventFunctor(func), id, lastId, (wxObject*)userData);
}
BOOL	wxEvtHandler_Unbind(WxObjectPtr p, void* func, int eventType, int id, int lastId, WxObjectPtr userData) {
	return __PTR->Unbind(wxEventTypeTag<wxEvent>(eventType), eventFunctor(func), id, lastId, (wxObject*)userData);
}
BOOL	wxEvtHandler_GetEvtHandlerEnabled(WxObjectPtr p) {
	return __PTR->GetEvtHandlerEnabled();
}
WxObjectPtr	wxEvtHandler_GetNextHandler(WxObjectPtr p) {
	return __PTR->GetNextHandler();
}
WxObjectPtr	wxEvtHandler_GetPreviousHandler(WxObjectPtr p) {
	return __PTR->GetPreviousHandler();
}
void 	wxEvtHandler_SetEvtHandlerEnabled(WxObjectPtr p, BOOL enabled) {
	__PTR->SetEvtHandlerEnabled(enabled);
}
void	wxEvtHandler_SetNextHandler(WxObjectPtr p, WxObjectPtr next) {
	__PTR->SetNextHandler((wxEvtHandler*)next);
}
void	wxEvtHandler_SetPrevioustHandler(WxObjectPtr p, WxObjectPtr prev) {
	__PTR->SetPreviousHandler((wxEvtHandler*)prev);
}
void	wxEvtHandler_Unlink(WxObjectPtr p) {
	__PTR->Unlink();
}
BOOL	wxEvtHandler_IsUnlinked(WxObjectPtr p) {
	return __PTR->IsUnlinked();
}