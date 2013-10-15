#include <wx/validate.h>
#include "wxutil.h"
#include "validator.h"

#define __PTR ((wxValidator*)(p))

WxObjectPtr wxValidator_New() {
	return new wxValidator;
}