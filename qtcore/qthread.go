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
func (this *QThread) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qthread.h:74
// index:0
// Public static Visibility=Default Availability=Available
// [8] Qt::HANDLE currentThreadId()
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
func (this *QThread) SetPriority(priority int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread11setPriorityENS_8PriorityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), priority)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:97
// index:0
// Public Visibility=Default Availability=Available
// [4] QThread::Priority priority() const
func (this *QThread) Priority() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread8priorityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qthread.h:99
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinished() const
func (this *QThread) IsFinished() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread10isFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRunning() const
func (this *QThread) IsRunning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread9isRunningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestInterruption()
func (this *QThread) RequestInterruption() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread19requestInterruptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isInterruptionRequested() const
func (this *QThread) IsInterruptionRequested() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread23isInterruptionRequestedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStackSize(uint)
func (this *QThread) SetStackSize(stackSize uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread12setStackSizeEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stackSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] uint stackSize() const
func (this *QThread) StackSize() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread9stackSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qthread.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void exit(int)
func (this *QThread) Exit(retcode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread4exitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), retcode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void exit(int)
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
func (this *QThread) EventDispatcher() *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread15eventDispatcherEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractEventDispatcherFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qthread.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEventDispatcher(QAbstractEventDispatcher *)
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
func (this *QThread) LoopLevel() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QThread9loopLevelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qthread.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(enum QThread::Priority)
func (this *QThread) Start(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread5startENS_8PriorityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(enum QThread::Priority)
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
func (this *QThread) Terminate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread9terminateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void quit()
func (this *QThread) Quit() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread4quitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:140
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wait(unsigned long)
func (this *QThread) Wait(time uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread4waitEm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), time)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qthread.h:140
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wait(unsigned long)
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
func (this *QThread) Run() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread3runEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:152
// index:0
// Protected Visibility=Default Availability=Available
// [4] int exec()
func (this *QThread) Exec() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread4execEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qthread.h:154
// index:0
// Protected static Visibility=Default Availability=Available
// [-2] void setTerminationEnabled(_Bool)
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
func (this *QThread) SetTerminationEnabled__() {
	// arg: 0, bool=Bool, =Invalid,
	enabled := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QThread21setTerminationEnabledEb", qtrt.FFI_TYPE_POINTER, enabled)
	qtrt.ErrPrint(err, rv)
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
