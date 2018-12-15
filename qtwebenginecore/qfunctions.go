package qtwebenginecore

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtquick"

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
		qtqml.KeepMe()
	}
	if false {
		qtpositioning.KeepMe()
	}
	if false {
		qtwebchannel.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:113
// index:104
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QWebEngineUrlScheme::Flags::enum_type, int)

/*

 */
func Operator_or104(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN19QWebEngineUrlScheme4FlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebenginehttprequest.h:101
// index:88
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QWebEngineHttpRequest &, QWebEngineHttpRequest &)

/*

 */
func Swap88(value1 QWebEngineHttpRequest_ITF, value2 QWebEngineHttpRequest_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QWebEngineHttpRequest_PTR() != nil {
		convArg0 = value1.QWebEngineHttpRequest_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QWebEngineHttpRequest_PTR() != nil {
		convArg1 = value2.QWebEngineHttpRequest_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR21QWebEngineHttpRequestS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

//  body block end
