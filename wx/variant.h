#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr	wxVariant_NewInt64(long long);
WxObjectPtr	wxVariant_NewString(String str);
long long	wxVariant_GetInt64(WxObjectPtr p);
String	wxVariant_GetString(WxObjectPtr p);


#ifdef __cplusplus
}
#endif