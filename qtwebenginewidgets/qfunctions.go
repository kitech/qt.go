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
// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:128
// index:105
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QWebEngineContextMenuData::MediaFlags::enum_type, int)

/*

 */
func Operator_or105(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN25QWebEngineContextMenuData9MediaFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginecontextmenudata.h:129
// index:106
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QWebEngineContextMenuData::EditFlags::enum_type, int)

/*

 */
func Operator_or106(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN25QWebEngineContextMenuData8EditFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginepage.h:375
// index:107
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QWebEnginePage::FindFlags::enum_type, int)

/*

 */
func Operator_or107(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN14QWebEnginePage8FindFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:82
// index:89
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QWebEngineHistoryItem &, QWebEngineHistoryItem &)

/*

 */
func Swap89(value1 QWebEngineHistoryItem_ITF, value2 QWebEngineHistoryItem_ITF) {
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

// /usr/include/qt/QtWebEngineWidgets/qwebenginescript.h:105
// index:90
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QWebEngineScript &, QWebEngineScript &)

/*

 */
func Swap90(value1 QWebEngineScript_ITF, value2 QWebEngineScript_ITF) {
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
