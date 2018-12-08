package qtwebenginewidgets

// /usr/include/qt/QtWebEngineWidgets/qwebenginescriptcollection.h
// #include <qwebenginescriptcollection.h>
// #include <QtWebEngineWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QWebEngineScriptCollection struct {
	*qtrt.CObject
}
type QWebEngineScriptCollection_ITF interface {
	QWebEngineScriptCollection_PTR() *QWebEngineScriptCollection
}

func (ptr *QWebEngineScriptCollection) QWebEngineScriptCollection_PTR() *QWebEngineScriptCollection {
	return ptr
}

func (this *QWebEngineScriptCollection) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineScriptCollection) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineScriptCollectionFromPointer(cthis unsafe.Pointer) *QWebEngineScriptCollection {
	return &QWebEngineScriptCollection{&qtrt.CObject{cthis}}
}
func (*QWebEngineScriptCollection) NewFromPointer(cthis unsafe.Pointer) *QWebEngineScriptCollection {
	return NewQWebEngineScriptCollectionFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescriptcollection.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QWebEngineScriptCollection()

/*

 */
func DeleteQWebEngineScriptCollection(this *QWebEngineScriptCollection) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QWebEngineScriptCollectionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescriptcollection.h:56
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*

 */
func (this *QWebEngineScriptCollection) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QWebEngineScriptCollection7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescriptcollection.h:57
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*

 */
func (this *QWebEngineScriptCollection) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QWebEngineScriptCollection5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescriptcollection.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int size() const

/*

 */
func (this *QWebEngineScriptCollection) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QWebEngineScriptCollection4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescriptcollection.h:59
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QWebEngineScript &) const

/*

 */
func (this *QWebEngineScriptCollection) Contains(value QWebEngineScript_ITF) bool {
	var convArg0 unsafe.Pointer
	if value != nil && value.QWebEngineScript_PTR() != nil {
		convArg0 = value.QWebEngineScript_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QWebEngineScriptCollection8containsERK16QWebEngineScript", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescriptcollection.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineScript findScript(const QString &) const

/*

 */
func (this *QWebEngineScriptCollection) FindScript(name string) *QWebEngineScript /*123*/ {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QWebEngineScriptCollection10findScriptERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineScriptFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineScript)
	return rv2
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescriptcollection.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool remove(const QWebEngineScript &)

/*

 */
func (this *QWebEngineScriptCollection) Remove(arg0 QWebEngineScript_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWebEngineScript_PTR() != nil {
		convArg0 = arg0.QWebEngineScript_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN26QWebEngineScriptCollection6removeERK16QWebEngineScript", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginescriptcollection.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*

 */
func (this *QWebEngineScriptCollection) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QWebEngineScriptCollection5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
