package qtcore

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h
// #include <qabstractnativeeventfilter.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QAbstractNativeEventFilter struct {
	*qtrt.CObject
}
type QAbstractNativeEventFilter_ITF interface {
	QAbstractNativeEventFilter_PTR() *QAbstractNativeEventFilter
}

func (ptr *QAbstractNativeEventFilter) QAbstractNativeEventFilter_PTR() *QAbstractNativeEventFilter {
	return ptr
}

func (this *QAbstractNativeEventFilter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAbstractNativeEventFilter) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAbstractNativeEventFilterFromPointer(cthis unsafe.Pointer) *QAbstractNativeEventFilter {
	return &QAbstractNativeEventFilter{&qtrt.CObject{cthis}}
}
func (*QAbstractNativeEventFilter) NewFromPointer(cthis unsafe.Pointer) *QAbstractNativeEventFilter {
	return NewQAbstractNativeEventFilterFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h:52
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractNativeEventFilter()
func NewQAbstractNativeEventFilter() *QAbstractNativeEventFilter {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractNativeEventFilterC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractNativeEventFilterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAbstractNativeEventFilter)
	return gothis
}

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractNativeEventFilter()
func DeleteQAbstractNativeEventFilter(this *QAbstractNativeEventFilter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractNativeEventFilterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractnativeeventfilter.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool nativeEventFilter(const QByteArray &, void *, long *)
func (this *QAbstractNativeEventFilter) NativeEventFilter(eventType QByteArray_ITF, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool {
	var convArg0 = eventType.QByteArray_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN26QAbstractNativeEventFilter17nativeEventFilterERK10QByteArrayPvPl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, message, &result)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
		qtrt.KeepMe()
	}
}

//  keep block end
