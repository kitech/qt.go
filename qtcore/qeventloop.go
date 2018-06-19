package qtcore

// /usr/include/qt/QtCore/qeventloop.h
// #include <qeventloop.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QEventLoop struct {
	*QObject
}
type QEventLoop_ITF interface {
	QObject_ITF
	QEventLoop_PTR() *QEventLoop
}

func (ptr *QEventLoop) QEventLoop_PTR() *QEventLoop { return ptr }

func (this *QEventLoop) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QEventLoop) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQEventLoopFromPointer(cthis unsafe.Pointer) *QEventLoop {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QEventLoop{bcthis0}
}
func (*QEventLoop) NewFromPointer(cthis unsafe.Pointer) *QEventLoop {
	return NewQEventLoopFromPointer(cthis)
}

// /usr/include/qt/QtCore/qeventloop.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QEventLoop) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QEventLoop10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qeventloop.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QEventLoop(QObject *)

/*
Constructs an event loop object with the given parent.
*/
func NewQEventLoop(parent QObject_ITF /*777 QObject **/) *QEventLoop {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoopC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEventLoopFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QEventLoop")
	return gothis
}

// /usr/include/qt/QtCore/qeventloop.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QEventLoop(QObject *)

/*
Constructs an event loop object with the given parent.
*/
func NewQEventLoop__() *QEventLoop {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoopC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQEventLoopFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QEventLoop")
	return gothis
}

// /usr/include/qt/QtCore/qeventloop.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QEventLoop()

/*

 */
func DeleteQEventLoop(this *QEventLoop) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoopD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qeventloop.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool processEvents(QEventLoop::ProcessEventsFlags)

/*
Processes pending events that match flags until there are no more events to process. Returns true if pending events were handled; otherwise returns false.

This function is especially useful if you have a long running operation and want to show its progress without allowing user input; i.e. by using the ExcludeUserInputEvents flag.

This function is simply a wrapper for QAbstractEventDispatcher::processEvents(). See the documentation for that function for details.
*/
func (this *QEventLoop) ProcessEvents(flags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop13processEventsE6QFlagsINS_17ProcessEventsFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qeventloop.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool processEvents(QEventLoop::ProcessEventsFlags)

/*
Processes pending events that match flags until there are no more events to process. Returns true if pending events were handled; otherwise returns false.

This function is especially useful if you have a long running operation and want to show its progress without allowing user input; i.e. by using the ExcludeUserInputEvents flag.

This function is simply a wrapper for QAbstractEventDispatcher::processEvents(). See the documentation for that function for details.
*/
func (this *QEventLoop) ProcessEvents__() bool {
	// arg: 0, QEventLoop::ProcessEventsFlags=Typedef, QEventLoop::ProcessEventsFlags=Typedef, QFlags<QEventLoop::ProcessEventsFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop13processEventsE6QFlagsINS_17ProcessEventsFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qeventloop.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void processEvents(QEventLoop::ProcessEventsFlags, int)

/*
Processes pending events that match flags until there are no more events to process. Returns true if pending events were handled; otherwise returns false.

This function is especially useful if you have a long running operation and want to show its progress without allowing user input; i.e. by using the ExcludeUserInputEvents flag.

This function is simply a wrapper for QAbstractEventDispatcher::processEvents(). See the documentation for that function for details.
*/
func (this *QEventLoop) ProcessEvents_1(flags int, maximumTime int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop13processEventsE6QFlagsINS_17ProcessEventsFlagEEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, maximumTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventloop.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] int exec(QEventLoop::ProcessEventsFlags)

/*
Enters the main event loop and waits until exit() is called. Returns the value that was passed to exit().

If flags are specified, only events of the types allowed by the flags will be processed.

It is necessary to call this function to start event handling. The main event loop receives events from the window system and dispatches these to the application widgets.

Generally speaking, no user interaction can take place before calling exec(). As a special case, modal widgets like QMessageBox can be used before calling exec(), because modal widgets use their own local event loop.

To make your application perform idle processing (i.e. executing a special function whenever there are no pending events), use a QTimer with 0 timeout. More sophisticated idle processing schemes can be achieved using processEvents().

See also QCoreApplication::quit(), exit(), and processEvents().
*/
func (this *QEventLoop) Exec(flags int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop4execE6QFlagsINS_17ProcessEventsFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qeventloop.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] int exec(QEventLoop::ProcessEventsFlags)

/*
Enters the main event loop and waits until exit() is called. Returns the value that was passed to exit().

If flags are specified, only events of the types allowed by the flags will be processed.

It is necessary to call this function to start event handling. The main event loop receives events from the window system and dispatches these to the application widgets.

Generally speaking, no user interaction can take place before calling exec(). As a special case, modal widgets like QMessageBox can be used before calling exec(), because modal widgets use their own local event loop.

To make your application perform idle processing (i.e. executing a special function whenever there are no pending events), use a QTimer with 0 timeout. More sophisticated idle processing schemes can be achieved using processEvents().

See also QCoreApplication::quit(), exit(), and processEvents().
*/
func (this *QEventLoop) Exec__() int {
	// arg: 0, QEventLoop::ProcessEventsFlags=Typedef, QEventLoop::ProcessEventsFlags=Typedef, QFlags<QEventLoop::ProcessEventsFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop4execE6QFlagsINS_17ProcessEventsFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qeventloop.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void exit(int)

/*
Tells the event loop to exit with a return code.

After this function has been called, the event loop returns from the call to exec(). The exec() function returns returnCode.

By convention, a returnCode of 0 means success, and any non-zero value indicates an error.

Note that unlike the C library function of the same name, this function does return to the caller -- it is event processing that stops.

See also QCoreApplication::quit(), quit(), and exec().
*/
func (this *QEventLoop) Exit(returnCode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop4exitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), returnCode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventloop.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void exit(int)

/*
Tells the event loop to exit with a return code.

After this function has been called, the event loop returns from the call to exec(). The exec() function returns returnCode.

By convention, a returnCode of 0 means success, and any non-zero value indicates an error.

Note that unlike the C library function of the same name, this function does return to the caller -- it is event processing that stops.

See also QCoreApplication::quit(), quit(), and exec().
*/
func (this *QEventLoop) Exit__() {
	// arg: 0, int=Int, =Invalid, , Invalid
	returnCode := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop4exitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), returnCode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventloop.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRunning() const

/*
Returns true if the event loop is running; otherwise returns false. The event loop is considered running from the time when exec() is called until exit() is called.

See also exec() and exit().
*/
func (this *QEventLoop) IsRunning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QEventLoop9isRunningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qeventloop.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void wakeUp()

/*
Wakes up the event loop.

See also QAbstractEventDispatcher::wakeUp().
*/
func (this *QEventLoop) WakeUp() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop6wakeUpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qeventloop.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QEventLoop) Event(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qeventloop.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void quit()

/*
Tells the event loop to exit normally.

Same as exit(0).

See also QCoreApplication::quit() and exit().
*/
func (this *QEventLoop) Quit() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QEventLoop4quitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QEventLoop__ProcessEventsFlag = int

//
const QEventLoop__AllEvents QEventLoop__ProcessEventsFlag = 0

//
const QEventLoop__ExcludeUserInputEvents QEventLoop__ProcessEventsFlag = 1

//
const QEventLoop__ExcludeSocketNotifiers QEventLoop__ProcessEventsFlag = 2

//
const QEventLoop__WaitForMoreEvents QEventLoop__ProcessEventsFlag = 4

//
const QEventLoop__X11ExcludeTimers QEventLoop__ProcessEventsFlag = 8

//
const QEventLoop__EventLoopExec QEventLoop__ProcessEventsFlag = 32

//
const QEventLoop__DialogExec QEventLoop__ProcessEventsFlag = 64

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
