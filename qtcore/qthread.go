package qtcore

// /usr/include/qt/QtCore/qthread.h
// #include <qthread.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// void run()
func (this *QThread) InheritRun(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "run", f)
}

// int exec()
func (this *QThread) InheritExec(f func() int) {
	qtrt.SetAllInheritCallback(this, "exec", f)
}

// void setTerminationEnabled(_Bool)
func (this *QThread) InheritSetTerminationEnabled(f func(enabled bool) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setTerminationEnabled", f)
}

/*

 */
type QThread struct {
	*QObject
}
type QThread_ITF interface {
	QObject_ITF
	QThread_PTR() *QThread
}

func (ptr *QThread) QThread_PTR() *QThread { return ptr }

func (this *QThread) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QThread) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQThreadFromPointer(cthis unsafe.Pointer) *QThread {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QThread{bcthis0}
}
func (*QThread) NewFromPointer(cthis unsafe.Pointer) *QThread {
	return NewQThreadFromPointer(cthis)
}

// /usr/include/qt/QtCore/qthread.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QThread) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qthread.h:74
// index:0
// Public static Visibility=Default Availability=Available
// [8] Qt::HANDLE currentThreadId()

/*
Returns the thread handle of the currently executing thread.

Warning: The handle returned by this function is used for internal purposes and should not be used in any application code.

Warning: On Windows, the returned value is a pseudo-handle for the current thread. It can't be used for numerical comparison. i.e., this function returns the DWORD (Windows-Thread ID) returned by the Win32 function getCurrentThreadId(), not the HANDLE (Windows-Thread HANDLE) returned by the Win32 function getCurrentThread().
*/
func (this *QThread) CurrentThreadId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread15currentThreadIdEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QThread_CurrentThreadId() int {
	var nilthis *QThread
	rv := nilthis.CurrentThreadId()
	return rv
}

// /usr/include/qt/QtCore/qthread.h:75
// index:0
// Public static Visibility=Default Availability=Available
// [8] QThread * currentThread()

/*
Returns a pointer to a QThread which manages the currently executing thread.
*/
func (this *QThread) CurrentThread() *QThread /*777 QThread **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread13currentThreadEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQThreadFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QThread_CurrentThread() *QThread /*777 QThread **/ {
	var nilthis *QThread
	rv := nilthis.CurrentThread()
	return rv
}

// /usr/include/qt/QtCore/qthread.h:76
// index:0
// Public static Visibility=Default Availability=Available
// [4] int idealThreadCount()

/*
Returns the ideal number of threads that can be run on the system. This is done querying the number of processor cores, both real and logical, in the system. This function returns 1 if the number of processor cores could not be detected.
*/
func (this *QThread) IdealThreadCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread16idealThreadCountEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QThread_IdealThreadCount() int {
	var nilthis *QThread
	rv := nilthis.IdealThreadCount()
	return rv
}

// /usr/include/qt/QtCore/qthread.h:77
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void yieldCurrentThread()

/*
Yields execution of the current thread to another runnable thread, if any. Note that the operating system decides to which thread to switch.
*/
func (this *QThread) YieldCurrentThread() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread18yieldCurrentThreadEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QThread_YieldCurrentThread() {
	var nilthis *QThread
	nilthis.YieldCurrentThread()
}

// /usr/include/qt/QtCore/qthread.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QThread(QObject *)

/*
Constructs a new QThread to manage a new thread. The parent takes ownership of the QThread. The thread does not begin executing until start() is called.

See also start().
*/
func NewQThread(parent QObject_ITF /*777 QObject **/) *QThread {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThreadC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQThreadFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QThread")
	return gothis
}

// /usr/include/qt/QtCore/qthread.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QThread(QObject *)

/*
Constructs a new QThread to manage a new thread. The parent takes ownership of the QThread. The thread does not begin executing until start() is called.

See also start().
*/
func NewQThread__() *QThread {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThreadC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQThreadFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QThread")
	return gothis
}

// /usr/include/qt/QtCore/qthread.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QThread()

/*

 */
func DeleteQThread(this *QThread) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThreadD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qthread.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPriority(enum QThread::Priority)

/*
This function sets the priority for a running thread. If the thread is not running, this function does nothing and returns immediately. Use start() to start a thread with a specific priority.

The priority argument can be any value in the QThread::Priority enum except for InheritPriorty.

The effect of the priority parameter is dependent on the operating system's scheduling policy. In particular, the priority will be ignored on systems that do not support thread priorities (such as on Linux, see http://linux.die.net/man/2/sched_setscheduler for more details).

This function was introduced in  Qt 4.1.

See also Priority, priority(), and start().
*/
func (this *QThread) SetPriority(priority int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread11setPriorityENS_8PriorityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), priority)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:97
// index:0
// Public Visibility=Default Availability=Available
// [4] QThread::Priority priority() const

/*
Returns the priority for a running thread. If the thread is not running, this function returns InheritPriority.

This function was introduced in  Qt 4.1.

See also Priority, setPriority(), and start().
*/
func (this *QThread) Priority() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread8priorityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qthread.h:99
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinished() const

/*
Returns true if the thread is finished; otherwise returns false.

See also isRunning().
*/
func (this *QThread) IsFinished() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread10isFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRunning() const

/*
Returns true if the thread is running; otherwise returns false.

See also isFinished().
*/
func (this *QThread) IsRunning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread9isRunningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestInterruption()

/*
Request the interruption of the thread. That request is advisory and it is up to code running on the thread to decide if and how it should act upon such request. This function does not stop any event loop running on the thread and does not terminate it in any way.

This function was introduced in  Qt 5.2.

See also isInterruptionRequested().
*/
func (this *QThread) RequestInterruption() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread19requestInterruptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isInterruptionRequested() const

/*
Return true if the task running on this thread should be stopped. An interruption can be requested by requestInterruption().

This function can be used to make long running tasks cleanly interruptible. Never checking or acting on the value returned by this function is safe, however it is advisable do so regularly in long running functions. Take care not to call it too often, to keep the overhead low.


  void long_task() {
       forever {
          if ( QThread::currentThread()->isInterruptionRequested() ) {
              return;
          }
      }
  }



This function was introduced in  Qt 5.2.

See also currentThread() and requestInterruption().
*/
func (this *QThread) IsInterruptionRequested() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread23isInterruptionRequestedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStackSize(uint)

/*
Sets the maximum stack size for the thread to stackSize. If stackSize is greater than zero, the maximum stack size is set to stackSize bytes, otherwise the maximum stack size is automatically determined by the operating system.

Warning: Most operating systems place minimum and maximum limits on thread stack sizes. The thread will fail to start if the stack size is outside these limits.

See also stackSize().
*/
func (this *QThread) SetStackSize(stackSize uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread12setStackSizeEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stackSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] uint stackSize() const

/*
Returns the maximum stack size for the thread (if set with setStackSize()); otherwise returns zero.

See also setStackSize().
*/
func (this *QThread) StackSize() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread9stackSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
	// unsigned int // 222
}

// /usr/include/qt/QtCore/qthread.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void exit(int)

/*
Tells the thread's event loop to exit with a return code.

After calling this function, the thread leaves the event loop and returns from the call to QEventLoop::exec(). The QEventLoop::exec() function returns returnCode.

By convention, a returnCode of 0 means success, any non-zero value indicates an error.

Note that unlike the C library function of the same name, this function does return to the caller -- it is event processing that stops.

No QEventLoops will be started anymore in this thread until QThread::exec() has been called again. If the eventloop in QThread::exec() is not running then the next call to QThread::exec() will also return immediately.

See also quit() and QEventLoop.
*/
func (this *QThread) Exit(retcode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread4exitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), retcode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void exit(int)

/*
Tells the thread's event loop to exit with a return code.

After calling this function, the thread leaves the event loop and returns from the call to QEventLoop::exec(). The QEventLoop::exec() function returns returnCode.

By convention, a returnCode of 0 means success, any non-zero value indicates an error.

Note that unlike the C library function of the same name, this function does return to the caller -- it is event processing that stops.

No QEventLoops will be started anymore in this thread until QThread::exec() has been called again. If the eventloop in QThread::exec() is not running then the next call to QThread::exec() will also return immediately.

See also quit() and QEventLoop.
*/
func (this *QThread) Exit__() {
	// arg: 0, int=Int, =Invalid,
	retcode := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread4exitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), retcode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractEventDispatcher * eventDispatcher() const

/*
Returns a pointer to the event dispatcher object for the thread. If no event dispatcher exists for the thread, this function returns 0.

This function was introduced in  Qt 5.0.

See also setEventDispatcher().
*/
func (this *QThread) EventDispatcher() *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread15eventDispatcherEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractEventDispatcherFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qthread.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEventDispatcher(QAbstractEventDispatcher *)

/*
Sets the event dispatcher for the thread to eventDispatcher. This is only possible as long as there is no event dispatcher installed for the thread yet. That is, before the thread has been started with start() or, in case of the main thread, before QCoreApplication has been instantiated. This method takes ownership of the object.

This function was introduced in  Qt 5.0.

See also eventDispatcher().
*/
func (this *QThread) SetEventDispatcher(eventDispatcher QAbstractEventDispatcher_ITF /*777 QAbstractEventDispatcher **/) {
	var convArg0 unsafe.Pointer
	if eventDispatcher != nil && eventDispatcher.QAbstractEventDispatcher_PTR() != nil {
		convArg0 = eventDispatcher.QAbstractEventDispatcher_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread18setEventDispatcherEP24QAbstractEventDispatcher", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:113
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QThread) Event(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:114
// index:0
// Public Visibility=Default Availability=Available
// [4] int loopLevel() const

/*
Returns the current event loop level for the thread.

Note: This can only be called within the thread itself, i.e. when it is the current thread.

This function was introduced in  Qt 5.5.
*/
func (this *QThread) LoopLevel() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread9loopLevelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qthread.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(enum QThread::Priority)

/*
Begins execution of the thread by calling run(). The operating system will schedule the thread according to the priority parameter. If the thread is already running, this function does nothing.

The effect of the priority parameter is dependent on the operating system's scheduling policy. In particular, the priority will be ignored on systems that do not support thread priorities (such as on Linux, see the sched_setscheduler documentation for more details).

See also run() and terminate().
*/
func (this *QThread) Start(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread5startENS_8PriorityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(enum QThread::Priority)

/*
Begins execution of the thread by calling run(). The operating system will schedule the thread according to the priority parameter. If the thread is already running, this function does nothing.

The effect of the priority parameter is dependent on the operating system's scheduling policy. In particular, the priority will be ignored on systems that do not support thread priorities (such as on Linux, see the sched_setscheduler documentation for more details).

See also run() and terminate().
*/
func (this *QThread) Start__() {
	// arg: 0, QThread::Priority=Enum, QThread::Priority=Enum,
	arg0 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread5startENS_8PriorityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void terminate()

/*
Terminates the execution of the thread. The thread may or may not be terminated immediately, depending on the operating system's scheduling policies. Use QThread::wait() after terminate(), to be sure.

When the thread is terminated, all threads waiting for the thread to finish will be woken up.

Warning: This function is dangerous and its use is discouraged. The thread can be terminated at any point in its code path. Threads can be terminated while modifying data. There is no chance for the thread to clean up after itself, unlock any held mutexes, etc. In short, use this function only if absolutely necessary.

Termination can be explicitly enabled or disabled by calling QThread::setTerminationEnabled(). Calling this function while termination is disabled results in the termination being deferred, until termination is re-enabled. See the documentation of QThread::setTerminationEnabled() for more information.

See also setTerminationEnabled().
*/
func (this *QThread) Terminate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread9terminateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void quit()

/*
Tells the thread's event loop to exit with return code 0 (success). Equivalent to calling QThread::exit(0).

This function does nothing if the thread does not have an event loop.

See also exit() and QEventLoop.
*/
func (this *QThread) Quit() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread4quitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:140
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wait(unsigned long)

/*
Blocks the thread until either of these conditions is met:


The thread associated with this QThread object has finished execution (i.e. when it returns from run()). This function will return true if the thread has finished. It also returns true if the thread has not been started yet.
time milliseconds has elapsed. If time is ULONG_MAX (the default), then the wait will never timeout (the thread must return from run()). This function will return false if the wait timed out.


This provides similar functionality to the POSIX pthread_join() function.

See also sleep() and terminate().
*/
func (this *QThread) Wait(time uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread4waitEm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), time)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:140
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wait(unsigned long)

/*
Blocks the thread until either of these conditions is met:


The thread associated with this QThread object has finished execution (i.e. when it returns from run()). This function will return true if the thread has finished. It also returns true if the thread has not been started yet.
time milliseconds has elapsed. If time is ULONG_MAX (the default), then the wait will never timeout (the thread must return from run()). This function will return false if the wait timed out.


This provides similar functionality to the POSIX pthread_join() function.

See also sleep() and terminate().
*/
func (this *QThread) Wait__() bool {
	// arg: 0, unsigned long=ULong, =Invalid,
	time := -1
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread4waitEm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), time)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:142
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void sleep(unsigned long)

/*
Forces the current thread to sleep for secs seconds.

See also msleep() and usleep().
*/
func (this *QThread) Sleep(arg0 uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread5sleepEm", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}
func QThread_Sleep(arg0 uint) {
	var nilthis *QThread
	nilthis.Sleep(arg0)
}

// /usr/include/qt/QtCore/qthread.h:143
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void msleep(unsigned long)

/*
Forces the current thread to sleep for msecs milliseconds.

See also sleep() and usleep().
*/
func (this *QThread) Msleep(arg0 uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread6msleepEm", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}
func QThread_Msleep(arg0 uint) {
	var nilthis *QThread
	nilthis.Msleep(arg0)
}

// /usr/include/qt/QtCore/qthread.h:144
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void usleep(unsigned long)

/*
Forces the current thread to sleep for usecs microseconds.

See also sleep() and msleep().
*/
func (this *QThread) Usleep(arg0 uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread6usleepEm", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}
func QThread_Usleep(arg0 uint) {
	var nilthis *QThread
	nilthis.Usleep(arg0)
}

// /usr/include/qt/QtCore/qthread.h:151
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void run()

/*
The starting point for the thread. After calling start(), the newly created thread calls this function. The default implementation simply calls exec().

You can reimplement this function to facilitate advanced thread management. Returning from this method will end the execution of the thread.

See also start() and wait().
*/
func (this *QThread) Run() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread3runEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:152
// index:0
// Protected Visibility=Default Availability=Available
// [4] int exec()

/*
Enters the event loop and waits until exit() is called, returning the value that was passed to exit(). The value returned is 0 if exit() is called via quit().

This function is meant to be called from within run(). It is necessary to call this function to start event handling.

See also quit() and exit().
*/
func (this *QThread) Exec() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread4execEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qthread.h:154
// index:0
// Protected static Visibility=Default Availability=Available
// [-2] void setTerminationEnabled(_Bool)

/*
Enables or disables termination of the current thread based on the enabled parameter. The thread must have been started by QThread.

When enabled is false, termination is disabled. Future calls to QThread::terminate() will return immediately without effect. Instead, the termination is deferred until termination is enabled.

When enabled is true, termination is enabled. Future calls to QThread::terminate() will terminate the thread normally. If termination has been deferred (i.e. QThread::terminate() was called with termination disabled), this function will terminate the calling thread immediately. Note that this function will not return in this case.

See also terminate().
*/
func (this *QThread) SetTerminationEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread21setTerminationEnabledEb", qtrt.FFI_TYPE_POINTER, enabled)
	qtrt.ErrPrint(err, rv)
}
func QThread_SetTerminationEnabled(enabled bool) {
	var nilthis *QThread
	nilthis.SetTerminationEnabled(enabled)
}

// /usr/include/qt/QtCore/qthread.h:154
// index:0
// Protected static Visibility=Default Availability=Available
// [-2] void setTerminationEnabled(_Bool)

/*
Enables or disables termination of the current thread based on the enabled parameter. The thread must have been started by QThread.

When enabled is false, termination is disabled. Future calls to QThread::terminate() will return immediately without effect. Instead, the termination is deferred until termination is enabled.

When enabled is true, termination is enabled. Future calls to QThread::terminate() will terminate the thread normally. If termination has been deferred (i.e. QThread::terminate() was called with termination disabled), this function will terminate the calling thread immediately. Note that this function will not return in this case.

See also terminate().
*/
func (this *QThread) SetTerminationEnabled__() {
	// arg: 0, bool=Bool, =Invalid,
	enabled := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread21setTerminationEnabledEb", qtrt.FFI_TYPE_POINTER, enabled)
	qtrt.ErrPrint(err, rv)
}

/*
This enum type indicates how the operating system should schedule newly created threads.


*/
type QThread__Priority = int

// scheduled only when no other threads are running.
const QThread__IdlePriority QThread__Priority = 0

// scheduled less often than LowPriority.
const QThread__LowestPriority QThread__Priority = 1

// scheduled less often than NormalPriority.
const QThread__LowPriority QThread__Priority = 2

// the default priority of the operating system.
const QThread__NormalPriority QThread__Priority = 3

// scheduled more often than NormalPriority.
const QThread__HighPriority QThread__Priority = 4

// scheduled more often than HighPriority.
const QThread__HighestPriority QThread__Priority = 5

// scheduled as often as possible.
const QThread__TimeCriticalPriority QThread__Priority = 6

// use the same priority as the creating thread. This is the default.
const QThread__InheritPriority QThread__Priority = 7

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
