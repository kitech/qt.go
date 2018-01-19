//  header block begin
// /usr/include/qt/QtCore/qthreadpool.h
// #include <qthreadpool.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qthreadpool.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QThreadPool) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QThreadPool10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:65
// index:0
// void QThreadPool(class QObject *)
func NewQThreadPool(parent unsafe.Pointer) *QThreadPool {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPoolC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QThreadPool{cthis}
}

// /usr/include/qt/QtCore/qthreadpool.h:66
// index:0
// virtual
// void ~QThreadPool()
func DeleteQThreadPool(*QThreadPool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPoolD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:68
// index:0
// static
// QThreadPool * globalInstance()
func (this *QThreadPool) GlobalInstance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool14globalInstanceEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QThreadPool_GlobalInstance() {
	// 0: (), ()
	var nilthis *QThreadPool
	nilthis.GlobalInstance()
}

// /usr/include/qt/QtCore/qthreadpool.h:70
// index:0
// void start(class QRunnable *, int)
func (this *QThreadPool) Start(runnable unsafe.Pointer, priority int) {
	// 0: (, QRunnable * runnable, int priority), (runnable, &priority)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool5startEP9QRunnablei", ffiqt.FFI_TYPE_VOID, this.cthis, runnable, &priority)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:71
// index:0
// bool tryStart(class QRunnable *)
func (this *QThreadPool) TryStart(runnable unsafe.Pointer) {
	// 0: (, QRunnable * runnable), (runnable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool8tryStartEP9QRunnable", ffiqt.FFI_TYPE_VOID, this.cthis, runnable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:73
// index:0
// int expiryTimeout()
func (this *QThreadPool) ExpiryTimeout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QThreadPool13expiryTimeoutEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:74
// index:0
// void setExpiryTimeout(int)
func (this *QThreadPool) SetExpiryTimeout(expiryTimeout int) {
	// 0: (, int expiryTimeout), (&expiryTimeout)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool16setExpiryTimeoutEi", ffiqt.FFI_TYPE_VOID, this.cthis, &expiryTimeout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:76
// index:0
// int maxThreadCount()
func (this *QThreadPool) MaxThreadCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QThreadPool14maxThreadCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:77
// index:0
// void setMaxThreadCount(int)
func (this *QThreadPool) SetMaxThreadCount(maxThreadCount int) {
	// 0: (, int maxThreadCount), (&maxThreadCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool17setMaxThreadCountEi", ffiqt.FFI_TYPE_VOID, this.cthis, &maxThreadCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:79
// index:0
// int activeThreadCount()
func (this *QThreadPool) ActiveThreadCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QThreadPool17activeThreadCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:81
// index:0
// void setStackSize(uint)
func (this *QThreadPool) SetStackSize(stackSize uint) {
	// 0: (, uint stackSize), (&stackSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool12setStackSizeEj", ffiqt.FFI_TYPE_VOID, this.cthis, &stackSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:82
// index:0
// uint stackSize()
func (this *QThreadPool) StackSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QThreadPool9stackSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:84
// index:0
// void reserveThread()
func (this *QThreadPool) ReserveThread() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool13reserveThreadEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:85
// index:0
// void releaseThread()
func (this *QThreadPool) ReleaseThread() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool13releaseThreadEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:87
// index:0
// bool waitForDone(int)
func (this *QThreadPool) WaitForDone(msecs int) {
	// 0: (, int msecs), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool11waitForDoneEi", ffiqt.FFI_TYPE_VOID, this.cthis, &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:89
// index:0
// void clear()
func (this *QThreadPool) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:93
// index:0
// void cancel(class QRunnable *)
func (this *QThreadPool) Cancel(runnable unsafe.Pointer) {
	// 0: (, QRunnable * runnable), (runnable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool6cancelEP9QRunnable", ffiqt.FFI_TYPE_VOID, this.cthis, runnable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:95
// index:0
// bool tryTake(class QRunnable *)
func (this *QThreadPool) TryTake(runnable unsafe.Pointer) {
	// 0: (, QRunnable * runnable), (runnable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QThreadPool7tryTakeEP9QRunnable", ffiqt.FFI_TYPE_VOID, this.cthis, runnable)
	gopp.ErrPrint(err, rv)
}

//  body block end
