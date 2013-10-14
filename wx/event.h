#include "cgoutil.h"

#ifdef __cplusplus
extern "C" {
#endif


WxObjectPtr wxEvent_GetEventObject(WxObjectPtr p);
int wxEvent_GetEventType(WxObjectPtr p);
int wxEvent_GetEventCategory(WxObjectPtr p);
int wxEvent_GetId(WxObjectPtr p);
WxObjectPtr wxEvent_GetEventUserData(WxObjectPtr p);
int wxEvent_GetSkipped(WxObjectPtr p);
long wxEvent_GetTimestamp(WxObjectPtr p);
BOOL wxEvent_IsCommandEvent(WxObjectPtr p);
void wxEvent_ResumePropagation(WxObjectPtr p, int propagationLevel);
void wxEvent_SetEventObject(WxObjectPtr p, WxObjectPtr object);
void wxEvent_SetEventType(WxObjectPtr p, int type);
void wxEvent_SetId(WxObjectPtr p, int id);
void wxEvent_SetTimestamp(WxObjectPtr p, long timestamp);
BOOL wxEvent_ShouldPropagate(WxObjectPtr p);
void wxEvent_Skip(WxObjectPtr p, BOOL skip);
void wxEvent_StopPropagation(WxObjectPtr p);


#ifdef __cplusplus
}
#endif