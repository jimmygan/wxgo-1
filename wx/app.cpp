#include <wx/app.h>
#include <wx/window.h>
#include "app.h"
#include "wxutil.h"

class MyApp: public wxApp{
    virtual bool OnInit();
	virtual int OnExit();
	
	wxDECLARE_DYNAMIC_CLASS(MyApp);
};

wxIMPLEMENT_DYNAMIC_CLASS(MyApp, wxApp)

IMPLEMENT_APP_NO_MAIN(MyApp)

bool MyApp::OnInit(){
	x_onWxAppInit(this);
    return true;
}

int MyApp::OnExit() {
	x_onWxAppExit();
	return 0;
}

#define __PTR ((MyApp*)(p))

void wxApp_SetTopWindow(WxObjectPtr p, WxObjectPtr window) {
	__PTR->SetTopWindow((wxWindow*)window);
}
void wxApp_ExitMainLoop(WxObjectPtr p) {
	__PTR->ExitMainLoop();
}