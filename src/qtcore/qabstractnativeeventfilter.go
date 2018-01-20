//  header block begin
// /usr/include/qt/QtCore/qabstractnativeeventfilter.h
// #include <qabstractnativeeventfilter.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QAbstractNativeEventFilter struct {
	*qtrt.CObject
}

func (this *QAbstractNativeEventFilter) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQAbstractNativeEventFilterFromPointer(cthis unsafe.Pointer) *QAbstractNativeEventFilter {
	return &QAbstractNativeEventFilter{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h:52
// index:0
// Public
// void QAbstractNativeEventFilter()
func NewQAbstractNativeEventFilter() *QAbstractNativeEventFilter {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAbstractNativeEventFilterC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractNativeEventFilterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h:53
// index:0
// Public virtual
// void ~QAbstractNativeEventFilter()
func DeleteQAbstractNativeEventFilter(*QAbstractNativeEventFilter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAbstractNativeEventFilterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h:55
// index:0
// Public pure virtual
// bool nativeEventFilter(const class QByteArray &, void *, long *)
func (this *QAbstractNativeEventFilter) NativeEventFilter(eventType *QByteArray, message unsafe.Pointer, result unsafe.Pointer) interface{} {
	var convArg0 = eventType.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAbstractNativeEventFilter17nativeEventFilterERK10QByteArrayPvPl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, message, result)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
