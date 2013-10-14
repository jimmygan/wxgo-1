#include "cgoutil.h"

#ifdef __cplusplus
extern "C" {
#endif

WxObjectPtr	wxEvtHandler_New();
void		wxEvtHandler_QueueEvent(WxObjectPtr p, WxObjectPtr event);
void		wxEvtHandler_AddPendingEvent (WxObjectPtr p, WxObjectPtr event);
BOOL		wxEvtHandler_ProcessEvent(WxObjectPtr, WxObjectPtr event);
BOOL		wxEvtHandler_ProcessEventLocally(WxObjectPtr p, WxObjectPtr event);
BOOL		wxEvtHandler_SafelyProcessEvent(WxObjectPtr p, WxObjectPtr event);
void		wxEvtHandler_ProcessPendingEvents(WxObjectPtr p);
void		wxEvtHandler_DeletePendingEvents(WxObjectPtr p);
void		wxEvtHandler_Bind(WxObjectPtr p, void* func, int eventType, int id, int lastId, WxObjectPtr userData);
BOOL		wxEvtHandler_Unbind(WxObjectPtr p, void* func, int eventType, int id, int lastId, WxObjectPtr userData);
BOOL		wxEvtHandler_GetEvtHandlerEnabled(WxObjectPtr p);
WxObjectPtr	wxEvtHandler_GetNextHandler(WxObjectPtr p);
WxObjectPtr	wxEvtHandler_GetPreviousHandler(WxObjectPtr p);
void 		wxEvtHandler_SetEvtHandlerEnabled(WxObjectPtr p, BOOL enabled);
void		wxEvtHandler_SetNextHandler(WxObjectPtr p, WxObjectPtr next);
void		wxEvtHandler_SetPrevioustHandler(WxObjectPtr p, WxObjectPtr prev);
void		wxEvtHandler_Unlink(WxObjectPtr p);
BOOL		wxEvtHandler_IsUnlinked(WxObjectPtr p);

#ifdef __cplusplus
}
#endif