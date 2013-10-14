#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr wxObject_New();
void wxObject_Delete(WxObjectPtr);
StringHandle wxObject_GetClassName(WxObjectPtr p);

StringHandle wxObject_GetBaseClassName1(StringHandle className);
StringHandle wxObject_GetBaseClassName2(StringHandle className);

#ifdef __cplusplus
}
#endif