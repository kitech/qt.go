//  header block begin
// /usr/include/qt/QtCore/qthread.h
// #include <qthread.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QThread struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qthread.h:72
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QThread) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:74
// index:0
// static
// Qt::HANDLE currentThreadId()
func (this *QThread) CurrentThreadId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread15currentThreadIdEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QThread_CurrentThreadId() {
	// 0: (), ()
	var nilthis *QThread
	nilthis.CurrentThreadId()
}

// /usr/include/qt/QtCore/qthread.h:75
// index:0
// static
// QThread * currentThread()
func (this *QThread) CurrentThread() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread13currentThreadEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QThread_CurrentThread() {
	// 0: (), ()
	var nilthis *QThread
	nilthis.CurrentThread()
}

// /usr/include/qt/QtCore/qthread.h:76
// index:0
// static
// int idealThreadCount()
func (this *QThread) IdealThreadCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread16idealThreadCountEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QThread_IdealThreadCount() {
	// 0: (), ()
	var nilthis *QThread
	nilthis.IdealThreadCount()
}

// /usr/include/qt/QtCore/qthread.h:77
// index:0
// static
// void yieldCurrentThread()
func (this *QThread) YieldCurrentThread() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread18yieldCurrentThreadEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QThread_YieldCurrentThread() {
	// 0: (), ()
	var nilthis *QThread
	nilthis.YieldCurrentThread()
}

// /usr/include/qt/QtCore/qthread.h:79
// index:0
// void QThread(class QObject *)
func NewQThread(parent unsafe.Pointer) *QThread {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThreadC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QThread{cthis}
}

// /usr/include/qt/QtCore/qthread.h:80
// index:0
// virtual
// void ~QThread()
func DeleteQThread(*QThread) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThreadD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:96
// index:0
// void setPriority(enum QThread::Priority)
func (this *QThread) SetPriority(priority int) {
	// 0: (, QThread::Priority priority), (&priority)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread11setPriorityENS_8PriorityE", ffiqt.FFI_TYPE_VOID, this.cthis, &priority)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:97
// index:0
// QThread::Priority priority()
func (this *QThread) Priority() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread8priorityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:99
// index:0
// bool isFinished()
func (this *QThread) IsFinished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread10isFinishedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:100
// index:0
// bool isRunning()
func (this *QThread) IsRunning() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread9isRunningEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:102
// index:0
// void requestInterruption()
func (this *QThread) RequestInterruption() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread19requestInterruptionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:103
// index:0
// bool isInterruptionRequested()
func (this *QThread) IsInterruptionRequested() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread23isInterruptionRequestedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:105
// index:0
// void setStackSize(uint)
func (this *QThread) SetStackSize(stackSize uint) {
	// 0: (, uint stackSize), (&stackSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread12setStackSizeEj", ffiqt.FFI_TYPE_VOID, this.cthis, &stackSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:106
// index:0
// uint stackSize()
func (this *QThread) StackSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread9stackSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:108
// index:0
// void exit(int)
func (this *QThread) Exit(retcode int) {
	// 0: (, int retcode), (&retcode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread4exitEi", ffiqt.FFI_TYPE_VOID, this.cthis, &retcode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:110
// index:0
// QAbstractEventDispatcher * eventDispatcher()
func (this *QThread) EventDispatcher() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread15eventDispatcherEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:111
// index:0
// void setEventDispatcher(class QAbstractEventDispatcher *)
func (this *QThread) SetEventDispatcher(eventDispatcher unsafe.Pointer) {
	// 0: (, QAbstractEventDispatcher * eventDispatcher), (eventDispatcher)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread18setEventDispatcherEP24QAbstractEventDispatcher", ffiqt.FFI_TYPE_VOID, this.cthis, eventDispatcher)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:113
// index:0
// virtual
// bool event(class QEvent *)
func (this *QThread) Event(event unsafe.Pointer) {
	// 0: (, QEvent * event), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.cthis, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:114
// index:0
// int loopLevel()
func (this *QThread) LoopLevel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QThread9loopLevelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:134
// index:0
// void start(enum QThread::Priority)
func (this *QThread) Start(arg0 int) {
	// 0: (, QThread::Priority arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread5startENS_8PriorityE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:135
// index:0
// void terminate()
func (this *QThread) Terminate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread9terminateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:136
// index:0
// void quit()
func (this *QThread) Quit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread4quitEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:140
// index:0
// bool wait(unsigned long)
func (this *QThread) Wait(time uint) {
	// 0: (, unsigned long time), (&time)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread4waitEm", ffiqt.FFI_TYPE_VOID, this.cthis, &time)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qthread.h:142
// index:0
// static
// void sleep(unsigned long)
func (this *QThread) Sleep(arg0 uint) {
	// 0: (unsigned long arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread5sleepEm", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QThread_Sleep(arg0 uint) {
	// 0: (unsigned long arg0), (arg0)
	var nilthis *QThread
	nilthis.Sleep(arg0)
}

// /usr/include/qt/QtCore/qthread.h:143
// index:0
// static
// void msleep(unsigned long)
func (this *QThread) Msleep(arg0 uint) {
	// 0: (unsigned long arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread6msleepEm", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QThread_Msleep(arg0 uint) {
	// 0: (unsigned long arg0), (arg0)
	var nilthis *QThread
	nilthis.Msleep(arg0)
}

// /usr/include/qt/QtCore/qthread.h:144
// index:0
// static
// void usleep(unsigned long)
func (this *QThread) Usleep(arg0 uint) {
	// 0: (unsigned long arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QThread6usleepEm", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QThread_Usleep(arg0 uint) {
	// 0: (unsigned long arg0), (arg0)
	var nilthis *QThread
	nilthis.Usleep(arg0)
}

//  body block end
