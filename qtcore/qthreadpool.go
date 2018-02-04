package qtcore

// /usr/include/qt/QtCore/qthreadpool.h
// #include <qthreadpool.h>
// #include <QtCore>

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
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin

type QThreadPool struct {
	*QObject
}

func (this *QThreadPool) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QThreadPool) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQThreadPoolFromPointer(cthis unsafe.Pointer) *QThreadPool {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QThreadPool{bcthis0}
}
func (*QThreadPool) NewFromPointer(cthis unsafe.Pointer) *QThreadPool {
	return NewQThreadPoolFromPointer(cthis)
}

// /usr/include/qt/QtCore/qthreadpool.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QThreadPool) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QThreadPool10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qthreadpool.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QThreadPool(QObject *)
func NewQThreadPool(parent *QObject /*777 QObject **/) *QThreadPool {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPoolC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQThreadPoolFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qthreadpool.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QThreadPool()
func DeleteQThreadPool(this *QThreadPool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPoolD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qthreadpool.h:68
// index:0
// Public static Visibility=Default Availability=Available
// [8] QThreadPool * globalInstance()
func (this *QThreadPool) GlobalInstance() *QThreadPool /*777 QThreadPool **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool14globalInstanceEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQThreadPoolFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QThreadPool_GlobalInstance() *QThreadPool /*777 QThreadPool **/ {
	var nilthis *QThreadPool
	rv := nilthis.GlobalInstance()
	return rv
}

// /usr/include/qt/QtCore/qthreadpool.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(QRunnable *, int)
func (this *QThreadPool) Start(runnable *QRunnable /*777 QRunnable **/, priority int) {
	var convArg0 = runnable.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool5startEP9QRunnablei", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, priority)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryStart(QRunnable *)
func (this *QThreadPool) TryStart(runnable *QRunnable /*777 QRunnable **/) bool {
	var convArg0 = runnable.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool8tryStartEP9QRunnable", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qthreadpool.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] int expiryTimeout()
func (this *QThreadPool) ExpiryTimeout() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QThreadPool13expiryTimeoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qthreadpool.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpiryTimeout(int)
func (this *QThreadPool) SetExpiryTimeout(expiryTimeout int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool16setExpiryTimeoutEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), expiryTimeout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxThreadCount()
func (this *QThreadPool) MaxThreadCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QThreadPool14maxThreadCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qthreadpool.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxThreadCount(int)
func (this *QThreadPool) SetMaxThreadCount(maxThreadCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool17setMaxThreadCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxThreadCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] int activeThreadCount()
func (this *QThreadPool) ActiveThreadCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QThreadPool17activeThreadCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qthreadpool.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStackSize(uint)
func (this *QThreadPool) SetStackSize(stackSize uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool12setStackSizeEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stackSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] uint stackSize()
func (this *QThreadPool) StackSize() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QThreadPool9stackSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qthreadpool.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reserveThread()
func (this *QThreadPool) ReserveThread() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool13reserveThreadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseThread()
func (this *QThreadPool) ReleaseThread() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool13releaseThreadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForDone(int)
func (this *QThreadPool) WaitForDone(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool11waitForDoneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qthreadpool.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QThreadPool) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancel(QRunnable *)
func (this *QThreadPool) Cancel(runnable *QRunnable /*777 QRunnable **/) {
	var convArg0 = runnable.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool6cancelEP9QRunnable", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryTake(QRunnable *)
func (this *QThreadPool) TryTake(runnable *QRunnable /*777 QRunnable **/) bool {
	var convArg0 = runnable.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool7tryTakeEP9QRunnable", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
