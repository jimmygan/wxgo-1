#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

WxObjectPtr	wxMenu_New(String title, long style);
WxObjectPtr wxMenu_Append (WxObjectPtr p, int id, String item, String helpString, int kind);


#ifdef __cplusplus
}
#endif