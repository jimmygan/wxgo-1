#include <wx/thread.h>
#include "wxutil.h"

BOOL g_CgoReady = FALSE;

void* operator new(size_t n) {
	return malloc(n);
}

// Override global operator delete to be aware of wx object destruction.
// Frankly, it is the destructors which we need to hook, but we can't.
void operator delete(void* p) {
	if(!p) {
		return;
	}
	if(g_CgoReady && wxThread::IsMain()) {
		x_remoteObjectPtrFromeObjectTable(p);
	}
	free(p);
}