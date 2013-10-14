package wx

//#include "wx.h"
import "C"

//export x_onWxAppInit
func x_onWxAppInit(app C.WxObjectPtr) bool {
	if onAppInit != nil {
		// app is created and will deleted by wxWidgets, do not hold it.
		return onAppInit(NewObjectFromPtr(app, false).(App))
	}
	return true
}

//export x_onWxAppExit
func x_onWxAppExit() {
	if onAppExit != nil {
		onAppExit()
	}
}

type OnAppInitFunc func(app App) bool
type OnAppExitFunc func()

var onAppInit OnAppInitFunc
var onAppExit OnAppExitFunc

func Entry(onInit OnAppInitFunc, onExit OnAppExitFunc) int {
	onAppInit = onInit
	onAppExit = onExit
	return int(C.wxMain())
}
