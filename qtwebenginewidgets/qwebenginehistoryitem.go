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
// extern C begin: 4
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
type QWebEngineHistoryItem struct {
	*qtrt.CObject
}
type QWebEngineHistoryItem_ITF interface {
	QWebEngineHistoryItem_PTR() *QWebEngineHistoryItem
}

func (ptr *QWebEngineHistoryItem) QWebEngineHistoryItem_PTR() *QWebEngineHistoryItem { return ptr }

func (this *QWebEngineHistoryItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineHistoryItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineHistoryItemFromPointer(cthis unsafe.Pointer) *QWebEngineHistoryItem {
	return &QWebEngineHistoryItem{&qtrt.CObject{cthis}}
}
func (*QWebEngineHistoryItem) NewFromPointer(cthis unsafe.Pointer) *QWebEngineHistoryItem {
	return NewQWebEngineHistoryItemFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:60
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineHistoryItem & operator=(const QWebEngineHistoryItem &)

/*

 */
func (this *QWebEngineHistoryItem) Operator_equal(other QWebEngineHistoryItem_ITF) *QWebEngineHistoryItem {
	var convArg0 unsafe.Pointer
	if other != nil && other.QWebEngineHistoryItem_PTR() != nil {
		convArg0 = other.QWebEngineHistoryItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHistoryItemaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineHistoryItemFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineHistoryItem)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QWebEngineHistoryItem()

/*

 */
func DeleteQWebEngineHistoryItem(this *QWebEngineHistoryItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHistoryItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl originalUrl() const

/*

 */
func (this *QWebEngineHistoryItem) OriginalUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineHistoryItem11originalUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:64
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url() const

/*

 */
func (this *QWebEngineHistoryItem) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineHistoryItem3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title() const

/*

 */
func (this *QWebEngineHistoryItem) Title() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineHistoryItem5titleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime lastVisited() const

/*

 */
func (this *QWebEngineHistoryItem) LastVisited() *qtcore.QDateTime /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineHistoryItem11lastVisitedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl iconUrl() const

/*

 */
func (this *QWebEngineHistoryItem) IconUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineHistoryItem7iconUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*

 */
func (this *QWebEngineHistoryItem) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QWebEngineHistoryItem7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginehistory.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QWebEngineHistoryItem &)

/*

 */
func (this *QWebEngineHistoryItem) Swap(other QWebEngineHistoryItem_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QWebEngineHistoryItem_PTR() != nil {
		convArg0 = other.QWebEngineHistoryItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QWebEngineHistoryItem4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
