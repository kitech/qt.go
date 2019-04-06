package qtwebengine

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtquick"
import "github.com/kitech/qt.go/qtwebenginecore"

func init_unused_10021() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtpositioning.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
	if false {
		qtwebchannel.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
	if false {
		qtwebenginecore.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtWebEngine/qtwebengineglobal.h:59
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void initialize()

/*

 */
func Initialize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QtWebEngine10initializeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}

//  body block end
