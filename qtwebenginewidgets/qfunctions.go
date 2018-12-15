package qtwebenginewidgets

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtprintsupport"
import "github.com/kitech/qt.go/qtquick"
import "github.com/kitech/qt.go/qtwebenginecore"

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
		qtquickwidgets.KeepMe()
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
		qtwidgets.KeepMe()
	}
	if false {
		qtprintsupport.KeepMe()
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
// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:83
// index:81
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QWebEngineHistoryItem &, QWebEngineHistoryItem &)

/*

 */
func Swap81(value1 QWebEngineHistoryItem_ITF, value2 QWebEngineHistoryItem_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QWebEngineHistoryItem_PTR() != nil {
		convArg0 = value1.QWebEngineHistoryItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QWebEngineHistoryItem_PTR() != nil {
		convArg1 = value2.QWebEngineHistoryItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR21QWebEngineHistoryItemS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:106
// index:82
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QWebEngineScript &, QWebEngineScript &)

/*

 */
func Swap82(value1 QWebEngineScript_ITF, value2 QWebEngineScript_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QWebEngineScript_PTR() != nil {
		convArg0 = value1.QWebEngineScript_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QWebEngineScript_PTR() != nil {
		convArg1 = value2.QWebEngineScript_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR16QWebEngineScriptS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

//  body block end
