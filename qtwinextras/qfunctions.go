package qtwinextras

import "unsafe"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

func init() {
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:112
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool isCompositionOpaque()
func IsCompositionOpaque() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin19isCompositionOpaqueEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

//  body block end
