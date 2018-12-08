package qtcore

// /usr/include/qt/QtCore/qthreadpool.h
// #include <qthreadpool.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QThreadPool struct {
	*QObject
}
type QThreadPool_ITF interface {
	QObject_ITF
	QThreadPool_PTR() *QThreadPool
}

func (ptr *QThreadPool) QThreadPool_PTR() *QThreadPool { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QThreadPool) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QThreadPool10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qthreadpool.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QThreadPool(QObject *)

/*
Constructs a thread pool with the given parent.
*/
func (*QThreadPool) NewForInherit(parent QObject_ITF /*777 QObject **/) *QThreadPool {
	return NewQThreadPool(parent)
}
func NewQThreadPool(parent QObject_ITF /*777 QObject **/) *QThreadPool {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPoolC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQThreadPoolFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QThreadPool")
	return gothis
}

// /usr/include/qt/QtCore/qthreadpool.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QThreadPool(QObject *)

/*
Constructs a thread pool with the given parent.
*/
func (*QThreadPool) NewForInheritp() *QThreadPool {
	return NewQThreadPoolp()
}
func NewQThreadPoolp() *QThreadPool {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPoolC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQThreadPoolFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QThreadPool")
	return gothis
}

// /usr/include/qt/QtCore/qthreadpool.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QThreadPool()

/*

 */
func DeleteQThreadPool(this *QThreadPool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPoolD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qthreadpool.h:68
// index:0
// Public static Visibility=Default Availability=Available
// [8] QThreadPool * globalInstance()

/*
Returns the global QThreadPool instance.
*/
func (this *QThreadPool) GlobalInstance() *QThreadPool /*777 QThreadPool **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool14globalInstanceEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQThreadPoolFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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

/*
Reserves a thread and uses it to run runnable, unless this thread will make the current thread count exceed maxThreadCount(). In that case, runnable is added to a run queue instead. The priority argument can be used to control the run queue's order of execution.

Note that the thread pool takes ownership of the runnable if runnable->autoDelete() returns true, and the runnable will be deleted automatically by the thread pool after the runnable->run() returns. If runnable->autoDelete() returns false, ownership of runnable remains with the caller. Note that changing the auto-deletion on runnable after calling this functions results in undefined behavior.
*/
func (this *QThreadPool) Start(runnable QRunnable_ITF /*777 QRunnable **/, priority int) {
	var convArg0 unsafe.Pointer
	if runnable != nil && runnable.QRunnable_PTR() != nil {
		convArg0 = runnable.QRunnable_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool5startEP9QRunnablei", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, priority)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(QRunnable *, int)

/*
Reserves a thread and uses it to run runnable, unless this thread will make the current thread count exceed maxThreadCount(). In that case, runnable is added to a run queue instead. The priority argument can be used to control the run queue's order of execution.

Note that the thread pool takes ownership of the runnable if runnable->autoDelete() returns true, and the runnable will be deleted automatically by the thread pool after the runnable->run() returns. If runnable->autoDelete() returns false, ownership of runnable remains with the caller. Note that changing the auto-deletion on runnable after calling this functions results in undefined behavior.
*/
func (this *QThreadPool) Startp(runnable QRunnable_ITF /*777 QRunnable **/) {
	var convArg0 unsafe.Pointer
	if runnable != nil && runnable.QRunnable_PTR() != nil {
		convArg0 = runnable.QRunnable_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	priority := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool5startEP9QRunnablei", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, priority)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryStart(QRunnable *)

/*
Attempts to reserve a thread to run runnable.

If no threads are available at the time of calling, then this function does nothing and returns false. Otherwise, runnable is run immediately using one available thread and this function returns true.

Note that the thread pool takes ownership of the runnable if runnable->autoDelete() returns true, and the runnable will be deleted automatically by the thread pool after the runnable->run() returns. If runnable->autoDelete() returns false, ownership of runnable remains with the caller. Note that changing the auto-deletion on runnable after calling this function results in undefined behavior.
*/
func (this *QThreadPool) TryStart(runnable QRunnable_ITF /*777 QRunnable **/) bool {
	var convArg0 unsafe.Pointer
	if runnable != nil && runnable.QRunnable_PTR() != nil {
		convArg0 = runnable.QRunnable_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool8tryStartEP9QRunnable", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qthreadpool.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] int expiryTimeout() const

/*

 */
func (this *QThreadPool) ExpiryTimeout() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QThreadPool13expiryTimeoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qthreadpool.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpiryTimeout(int)

/*

 */
func (this *QThreadPool) SetExpiryTimeout(expiryTimeout int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool16setExpiryTimeoutEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), expiryTimeout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxThreadCount() const

/*

 */
func (this *QThreadPool) MaxThreadCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QThreadPool14maxThreadCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qthreadpool.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxThreadCount(int)

/*

 */
func (this *QThreadPool) SetMaxThreadCount(maxThreadCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool17setMaxThreadCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxThreadCount)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] int activeThreadCount() const

/*

 */
func (this *QThreadPool) ActiveThreadCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QThreadPool17activeThreadCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qthreadpool.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStackSize(uint)

/*

 */
func (this *QThreadPool) SetStackSize(stackSize uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool12setStackSizeEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stackSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] uint stackSize() const

/*

 */
func (this *QThreadPool) StackSize() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QThreadPool9stackSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qthreadpool.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reserveThread()

/*
Reserves one thread, disregarding activeThreadCount() and maxThreadCount().

Once you are done with the thread, call releaseThread() to allow it to be reused.

Note: This function will always increase the number of active threads. This means that by using this function, it is possible for activeThreadCount() to return a value greater than maxThreadCount() .

See also releaseThread().
*/
func (this *QThreadPool) ReserveThread() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool13reserveThreadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseThread()

/*
Releases a thread previously reserved by a call to reserveThread().

Note: Calling this function without previously reserving a thread temporarily increases maxThreadCount(). This is useful when a thread goes to sleep waiting for more work, allowing other threads to continue. Be sure to call reserveThread() when done waiting, so that the thread pool can correctly maintain the activeThreadCount().

See also reserveThread().
*/
func (this *QThreadPool) ReleaseThread() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool13releaseThreadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForDone(int)

/*
Waits up to msecs milliseconds for all threads to exit and removes all threads from the thread pool. Returns true if all threads were removed; otherwise it returns false. If msecs is -1 (the default), the timeout is ignored (waits for the last thread to exit).
*/
func (this *QThreadPool) WaitForDone(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool11waitForDoneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qthreadpool.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForDone(int)

/*
Waits up to msecs milliseconds for all threads to exit and removes all threads from the thread pool. Returns true if all threads were removed; otherwise it returns false. If msecs is -1 (the default), the timeout is ignored (waits for the last thread to exit).
*/
func (this *QThreadPool) WaitForDonep() bool {
	// arg: 0, int=Int, =Invalid, , Invalid
	msecs := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool11waitForDoneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qthreadpool.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Removes the runnables that are not yet started from the queue. The runnables for which runnable->autoDelete() returns true are deleted.

This function was introduced in  Qt 5.2.

See also start().
*/
func (this *QThreadPool) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancel(QRunnable *)

/*

 */
func (this *QThreadPool) Cancel(runnable QRunnable_ITF /*777 QRunnable **/) {
	var convArg0 unsafe.Pointer
	if runnable != nil && runnable.QRunnable_PTR() != nil {
		convArg0 = runnable.QRunnable_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool6cancelEP9QRunnable", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthreadpool.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryTake(QRunnable *)

/*
Attempts to remove the specified runnable from the queue if it is not yet started. If the runnable had not been started, returns true, and ownership of runnable is transferred to the caller (even when runnable->autoDelete() == true). Otherwise returns false.

Note: If runnable->autoDelete() == true, this function may remove the wrong runnable. This is known as the ABA problem: the original runnable may already have executed and has since been deleted. The memory is re-used for another runnable, which then gets removed instead of the intended one. For this reason, we recommend calling this function only for runnables that are not auto-deleting.

This function was introduced in  Qt 5.9.

See also start() and QRunnable::autoDelete().
*/
func (this *QThreadPool) TryTake(runnable QRunnable_ITF /*777 QRunnable **/) bool {
	var convArg0 unsafe.Pointer
	if runnable != nil && runnable.QRunnable_PTR() != nil {
		convArg0 = runnable.QRunnable_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QThreadPool7tryTakeEP9QRunnable", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
