//  header block begin
// /usr/include/qt/QtCore/qthreadpool.h
// #include <qthreadpool.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
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
type QThreadPool struct {
	*QObject
}

func (this *QThreadPool) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQThreadPoolFromPointer(cthis unsafe.Pointer) *QThreadPool {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QThreadPool{bcthis0}
}

// /usr/include/qt/QtCore/qthreadpool.h:56
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QThreadPool) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QThreadPool10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthreadpool.h:65
// index:0
// Public
// void QThreadPool(class QObject *)
func NewQThreadPool(parent unsafe.Pointer) *QThreadPool {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPoolC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQThreadPoolFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qthreadpool.h:66
// index:0
// Public virtual
// void ~QThreadPool()
func DeleteQThreadPool(*QThreadPool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPoolD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:68
// index:0
// Public static
// QThreadPool * globalInstance()
func (this *QThreadPool) GlobalInstance() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool14globalInstanceEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QThreadPool_GlobalInstance() {
	var nilthis *QThreadPool
	nilthis.GlobalInstance()
}

// /usr/include/qt/QtCore/qthreadpool.h:70
// index:0
// Public
// void start(class QRunnable *, int)
func (this *QThreadPool) Start(runnable unsafe.Pointer, priority int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool5startEP9QRunnablei", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), runnable, &priority)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:71
// index:0
// Public
// bool tryStart(class QRunnable *)
func (this *QThreadPool) TryStart(runnable unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool8tryStartEP9QRunnable", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), runnable)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthreadpool.h:73
// index:0
// Public
// int expiryTimeout()
func (this *QThreadPool) ExpiryTimeout() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QThreadPool13expiryTimeoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthreadpool.h:74
// index:0
// Public
// void setExpiryTimeout(int)
func (this *QThreadPool) SetExpiryTimeout(expiryTimeout int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool16setExpiryTimeoutEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &expiryTimeout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:76
// index:0
// Public
// int maxThreadCount()
func (this *QThreadPool) MaxThreadCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QThreadPool14maxThreadCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthreadpool.h:77
// index:0
// Public
// void setMaxThreadCount(int)
func (this *QThreadPool) SetMaxThreadCount(maxThreadCount int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool17setMaxThreadCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &maxThreadCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:79
// index:0
// Public
// int activeThreadCount()
func (this *QThreadPool) ActiveThreadCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QThreadPool17activeThreadCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthreadpool.h:81
// index:0
// Public
// void setStackSize(uint)
func (this *QThreadPool) SetStackSize(stackSize uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool12setStackSizeEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &stackSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:82
// index:0
// Public
// uint stackSize()
func (this *QThreadPool) StackSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QThreadPool9stackSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthreadpool.h:84
// index:0
// Public
// void reserveThread()
func (this *QThreadPool) ReserveThread() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool13reserveThreadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:85
// index:0
// Public
// void releaseThread()
func (this *QThreadPool) ReleaseThread() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool13releaseThreadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:87
// index:0
// Public
// bool waitForDone(int)
func (this *QThreadPool) WaitForDone(msecs int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool11waitForDoneEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qthreadpool.h:89
// index:0
// Public
// void clear()
func (this *QThreadPool) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:93
// index:0
// Public
// void cancel(class QRunnable *)
func (this *QThreadPool) Cancel(runnable unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool6cancelEP9QRunnable", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), runnable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:95
// index:0
// Public
// bool tryTake(class QRunnable *)
func (this *QThreadPool) TryTake(runnable unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool7tryTakeEP9QRunnable", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), runnable)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
