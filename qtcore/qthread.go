package qtcore

// /usr/include/qt/QtCore/qthread.h
// #include <qthread.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
// void run()
func (this *QThread) InheritRun(f func()) {
	ffiqt.SetAllInheritCallback(this, "run", f)
}

// int exec()
func (this *QThread) InheritExec(f func() int) {
	ffiqt.SetAllInheritCallback(this, "exec", f)
}

// void setTerminationEnabled(_Bool)
func (this *QThread) InheritSetTerminationEnabled(f func(enabled bool)) {
	ffiqt.SetAllInheritCallback(this, "setTerminationEnabled", f)
}

type QThread struct {
	*QObject
}

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
// [8] const QMetaObject * metaObject()
func (this *QThread) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qthread.h:74
// index:0
// Public static Visibility=Default Availability=Available
// [8] Qt::HANDLE currentThreadId()
func (this *QThread) CurrentThreadId() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread15currentThreadIdEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
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
func (this *QThread) CurrentThread() *QThread /*777 QThread **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread13currentThreadEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQThreadFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
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
func (this *QThread) IdealThreadCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread16idealThreadCountEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
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
func (this *QThread) YieldCurrentThread() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread18yieldCurrentThreadEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QThread_YieldCurrentThread() {
	var nilthis *QThread
	nilthis.YieldCurrentThread()
}

// /usr/include/qt/QtCore/qthread.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QThread(QObject *)
func NewQThread(parent *QObject /*777 QObject **/) *QThread {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThreadC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQThreadFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qthread.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QThread()
func DeleteQThread(this *QThread) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThreadD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qthread.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPriority(enum QThread::Priority)
func (this *QThread) SetPriority(priority int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread11setPriorityENS_8PriorityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), priority)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:97
// index:0
// Public Visibility=Default Availability=Available
// [4] QThread::Priority priority()
func (this *QThread) Priority() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread8priorityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qthread.h:99
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinished()
func (this *QThread) IsFinished() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread10isFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRunning()
func (this *QThread) IsRunning() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread9isRunningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestInterruption()
func (this *QThread) RequestInterruption() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread19requestInterruptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isInterruptionRequested()
func (this *QThread) IsInterruptionRequested() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread23isInterruptionRequestedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStackSize(uint)
func (this *QThread) SetStackSize(stackSize uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread12setStackSizeEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), stackSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] uint stackSize()
func (this *QThread) StackSize() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread9stackSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qthread.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void exit(int)
func (this *QThread) Exit(retcode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread4exitEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), retcode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractEventDispatcher * eventDispatcher()
func (this *QThread) EventDispatcher() *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread15eventDispatcherEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractEventDispatcherFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qthread.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEventDispatcher(QAbstractEventDispatcher *)
func (this *QThread) SetEventDispatcher(eventDispatcher *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/) {
	var convArg0 = eventDispatcher.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread18setEventDispatcherEP24QAbstractEventDispatcher", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:113
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QThread) Event(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:114
// index:0
// Public Visibility=Default Availability=Available
// [4] int loopLevel()
func (this *QThread) LoopLevel() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread9loopLevelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qthread.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(enum QThread::Priority)
func (this *QThread) Start(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread5startENS_8PriorityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void terminate()
func (this *QThread) Terminate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread9terminateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void quit()
func (this *QThread) Quit() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread4quitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:140
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wait(unsigned long)
func (this *QThread) Wait(time uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread4waitEm", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), time)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:142
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void sleep(unsigned long)
func (this *QThread) Sleep(arg0 uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread5sleepEm", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QThread_Sleep(arg0 uint) {
	var nilthis *QThread
	nilthis.Sleep(arg0)
}

// /usr/include/qt/QtCore/qthread.h:143
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void msleep(unsigned long)
func (this *QThread) Msleep(arg0 uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread6msleepEm", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QThread_Msleep(arg0 uint) {
	var nilthis *QThread
	nilthis.Msleep(arg0)
}

// /usr/include/qt/QtCore/qthread.h:144
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void usleep(unsigned long)
func (this *QThread) Usleep(arg0 uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread6usleepEm", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QThread_Usleep(arg0 uint) {
	var nilthis *QThread
	nilthis.Usleep(arg0)
}

// /usr/include/qt/QtCore/qthread.h:151
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void run()
func (this *QThread) Run() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread3runEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:152
// index:0
// Protected Visibility=Default Availability=Available
// [4] int exec()
func (this *QThread) Exec() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread4execEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qthread.h:154
// index:0
// Protected static Visibility=Default Availability=Available
// [-2] void setTerminationEnabled(_Bool)
func (this *QThread) SetTerminationEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread21setTerminationEnabledEb", ffiqt.FFI_TYPE_POINTER, enabled)
	gopp.ErrPrint(err, rv)
}
func QThread_SetTerminationEnabled(enabled bool) {
	var nilthis *QThread
	nilthis.SetTerminationEnabled(enabled)
}

type QThread__Priority = int

const QThread__IdlePriority QThread__Priority = 0
const QThread__LowestPriority QThread__Priority = 1
const QThread__LowPriority QThread__Priority = 2
const QThread__NormalPriority QThread__Priority = 3
const QThread__HighPriority QThread__Priority = 4
const QThread__HighestPriority QThread__Priority = 5
const QThread__TimeCriticalPriority QThread__Priority = 6
const QThread__InheritPriority QThread__Priority = 7

//  body block end
