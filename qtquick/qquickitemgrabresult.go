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
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtnetwork"
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
		ffiqt.KeepMe()
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
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QQuickItemGrabResult) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QQuickItemGrabResult10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:64
// index:0
// Public
// QImage image()
func (this *QQuickItemGrabResult) Image() *qtgui.QImage /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QQuickItemGrabResult5imageEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:65
// index:0
// Public
// QUrl url()
func (this *QQuickItemGrabResult) Url() *qtcore.QUrl /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QQuickItemGrabResult3urlEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:68
// index:0
// Public
// bool saveToFile(const class QString &)
func (this *QQuickItemGrabResult) SaveToFile(fileName *qtcore.QString) bool {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QQuickItemGrabResult10saveToFileERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:70
// index:1
// Public
// bool saveToFile(const class QString &)
func (this *QQuickItemGrabResult) SaveToFile_1(fileName *qtcore.QString) bool {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QQuickItemGrabResult10saveToFileERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:73
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QQuickItemGrabResult) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QQuickItemGrabResult5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitemgrabresult.h:76
// index:0
// Public
// void ready()
func (this *QQuickItemGrabResult) Ready() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QQuickItemGrabResult5readyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
