#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr	wxMenu_New(StringHandle title, long style);
WxObjectPtr wxMenu_Append (WxObjectPtr p, int id, StringHandle item, StringHandle helpString, int kind);


#ifdef __cplusplus
}
#endif