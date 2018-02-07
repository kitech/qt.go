package qtquick

// /usr/include/qt/QtQuick/qquickitemgrabresult.h
// #include <qquickitemgrabresult.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"
import "qt.go/qtgui"
import "qt.go/qtqml"

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
		qtrt.KeepMe()
	}
	if false {
		gopp.KeepMe()
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
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool event(class QEvent *)
func (this *QQuickItemGrabResult) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QQuickItemGrabResult struct {
	*qtcore.QObject
}

func (this *QQuickItemGrabResult) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQuickItemGrabResult) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQuickItemGrabResultFromPointer(cthis unsafe.Pointer) *QQuickItemGrabResult {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQuickItemGrabResult{bcthis0}
}
func (*QQuickItemGrabResult) NewFromPointer(cthis unsafe.Pointer) *QQuickItemGrabResult {
	return NewQQuickItemGrabResultFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QQuickItemGrabResult) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QQuickItemGrabResult10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:64
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage image()
func (this *QQuickItemGrabResult) Image() *qtgui.QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QQuickItemGrabResult5imageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQImage)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url()
func (this *QQuickItemGrabResult) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QQuickItemGrabResult3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool saveToFile(const QString &)
func (this *QQuickItemGrabResult) SaveToFile(fileName *qtcore.QString) bool {
	var convArg0 = fileName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQuickItemGrabResult10saveToFileERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:70
// index:1
// Public Visibility=Default Availability=Available
// [1] bool saveToFile(const QString &)
func (this *QQuickItemGrabResult) SaveToFile_1(fileName *qtcore.QString) bool {
	var convArg0 = fileName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QQuickItemGrabResult10saveToFileERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:73
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QQuickItemGrabResult) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQuickItemGrabResult5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ready()
func (this *QQuickItemGrabResult) Ready() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQuickItemGrabResult5readyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

func DeleteQQuickItemGrabResult(this *QQuickItemGrabResult) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQuickItemGrabResultD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
