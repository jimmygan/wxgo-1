#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr wxObject_New();
void wxObject_Delete(WxObjectPtr);
String wxObject_GetClassName(WxObjectPtr p);

String wxObject_GetBaseClassName1(String className);
String wxObject_GetBaseClassName2(String className);

#ifdef __cplusplus
}
#endif