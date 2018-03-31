package qtwebenginewidgets

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h
// #include <qwebenginehistory.h>
// #include <QtWebEngineWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
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

//  ext block end

//  body block begin

/*

 */
type QWebEngineHistory struct {
	*qtrt.CObject
}
type QWebEngineHistory_ITF interface {
	QWebEngineHistory_PTR() *QWebEngineHistory
}

func (ptr *QWebEngineHistory) QWebEngineHistory_PTR() *QWebEngineHistory { return ptr }

func (this *QWebEngineHistory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineHistory) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineHistoryFromPointer(cthis unsafe.Pointer) *QWebEngineHistory {
	return &QWebEngineHistory{&qtrt.CObject{cthis}}
}
func (*QWebEngineHistory) NewFromPointer(cthis unsafe.Pointer) *QWebEngineHistory {
	return NewQWebEngineHistoryFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*

 */
func (this *QWebEngineHistory) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineHistory5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:93
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canGoBack() const

/*

 */
func (this *QWebEngineHistory) CanGoBack() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineHistory9canGoBackEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canGoForward() const

/*

 */
func (this *QWebEngineHistory) CanGoForward() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineHistory12canGoForwardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void back()

/*

 */
func (this *QWebEngineHistory) Back() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineHistory4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void forward()

/*

 */
func (this *QWebEngineHistory) Forward() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineHistory7forwardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void goToItem(const QWebEngineHistoryItem &)

/*

 */
func (this *QWebEngineHistory) GoToItem(item QWebEngineHistoryItem_ITF) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QWebEngineHistoryItem_PTR() != nil {
		convArg0 = item.QWebEngineHistoryItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineHistory8goToItemERK21QWebEngineHistoryItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineHistoryItem backItem() const

/*

 */
func (this *QWebEngineHistory) BackItem() *QWebEngineHistoryItem /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineHistory8backItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineHistoryItemFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineHistoryItem)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineHistoryItem currentItem() const

/*

 */
func (this *QWebEngineHistory) CurrentItem() *QWebEngineHistoryItem /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineHistory11currentItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineHistoryItemFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineHistoryItem)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineHistoryItem forwardItem() const

/*

 */
func (this *QWebEngineHistory) ForwardItem() *QWebEngineHistoryItem /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineHistory11forwardItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineHistoryItemFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineHistoryItem)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineHistoryItem itemAt(int) const

/*

 */
func (this *QWebEngineHistory) ItemAt(i int) *QWebEngineHistoryItem /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineHistory6itemAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineHistoryItemFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineHistoryItem)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentItemIndex() const

/*

 */
func (this *QWebEngineHistory) CurrentItemIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineHistory16currentItemIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*

 */
func (this *QWebEngineHistory) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QWebEngineHistory5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQWebEngineHistory(this *QWebEngineHistory) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QWebEngineHistoryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
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

//  keep block end
