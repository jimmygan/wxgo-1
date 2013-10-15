#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr wxBoxSizer_New(int orient);
int wxBoxSizer_GetOrientation(WxObjectPtr p);
void wxBoxSizer_SetOrientation(WxObjectPtr p, int orient);



#ifdef __cplusplus
}
#endif