#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr		wxThreadEvent_New();
WxObjectPtr		wxThreadEvent_New2(int eventType, int id);
int				wxThreadEvent_GetEventCategory(WxObjectPtr p);
void			wxThreadEvent_SetPayload(WxObjectPtr p, void* payload);
void*			wxThreadEvent_GetPayload(WxObjectPtr p);
long			wxThreadEvent_GetExtraLong(WxObjectPtr p);
int				wxThreadEvent_GetInt(WxObjectPtr p);
String	wxThreadEvent_GetString(WxObjectPtr p);
void			wxThreadEvent_SetExtraLong(WxObjectPtr p, long l);
void			wxThreadEvent_SetInt(WxObjectPtr p, int n);
void 			wxThreadEvent_SetString(WxObjectPtr p, String str);


#ifdef __cplusplus
}
#endif