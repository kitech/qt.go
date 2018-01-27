package qtqml

// /usr/include/qt/QtQml/qqmlincubator.h
// #include <qqmlincubator.h>
// #include <QtQml>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQmlIncubator struct {
	*qtrt.CObject
}

func (this *QQmlIncubator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlIncubator) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQQmlIncubatorFromPointer(cthis unsafe.Pointer) *QQmlIncubator {
	return &QQmlIncubator{&qtrt.CObject{cthis}}
}
func (*QQmlIncubator) NewFromPointer(cthis unsafe.Pointer) *QQmlIncubator {
	return NewQQmlIncubatorFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlincubator.h:68
// index:0
// Public
// void QQmlIncubator(QQmlIncubator::IncubationMode)
func NewQQmlIncubator(arg0 int) *QQmlIncubator {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlIncubatorC2ENS_14IncubationModeE", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlIncubatorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQml/qqmlincubator.h:69
// index:0
// Public virtual
// void ~QQmlIncubator()
func DeleteQQmlIncubator(*QQmlIncubator) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlIncubatorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:71
// index:0
// Public
// void clear()
func (this *QQmlIncubator) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlIncubator5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:72
// index:0
// Public
// void forceCompletion()
func (this *QQmlIncubator) ForceCompletion() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlIncubator15forceCompletionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:74
// index:0
// Public
// bool isNull()
func (this *QQmlIncubator) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlIncubator6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlincubator.h:75
// index:0
// Public
// bool isReady()
func (this *QQmlIncubator) IsReady() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlIncubator7isReadyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlincubator.h:76
// index:0
// Public
// bool isError()
func (this *QQmlIncubator) IsError() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlIncubator7isErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlincubator.h:77
// index:0
// Public
// bool isLoading()
func (this *QQmlIncubator) IsLoading() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlIncubator9isLoadingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlincubator.h:81
// index:0
// Public
// QQmlIncubator::IncubationMode incubationMode()
func (this *QQmlIncubator) IncubationMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlIncubator14incubationModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:83
// index:0
// Public
// QQmlIncubator::Status status()
func (this *QQmlIncubator) Status() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlIncubator6statusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:85
// index:0
// Public
// QObject * object()
func (this *QQmlIncubator) Object() *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QQmlIncubator6objectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlincubator.h:88
// index:0
// Protected virtual
// void statusChanged(QQmlIncubator::Status)
func (this *QQmlIncubator) StatusChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlIncubator13statusChangedENS_6StatusE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlincubator.h:89
// index:0
// Protected virtual
// void setInitialState(QObject *)
func (this *QQmlIncubator) SetInitialState(arg0 *qtcore.QObject /*777 QObject **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QQmlIncubator15setInitialStateEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QQmlIncubator__IncubationMode = int

const QQmlIncubator__Asynchronous QQmlIncubator__IncubationMode = 0
const QQmlIncubator__AsynchronousIfNested QQmlIncubator__IncubationMode = 1
const QQmlIncubator__Synchronous QQmlIncubator__IncubationMode = 2

type QQmlIncubator__Status = int

const QQmlIncubator__Null QQmlIncubator__Status = 0
const QQmlIncubator__Ready QQmlIncubator__Status = 1
const QQmlIncubator__Loading QQmlIncubator__Status = 2
const QQmlIncubator__Error QQmlIncubator__Status = 3

//  body block end
