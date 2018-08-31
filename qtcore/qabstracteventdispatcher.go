package qtcore

// /usr/include/qt/QtCore/qabstracteventdispatcher.h
// #include <qabstracteventdispatcher.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
type QAbstractEventDispatcher struct {
	*QObject
}
type QAbstractEventDispatcher_ITF interface {
	QObject_ITF
	QAbstractEventDispatcher_PTR() *QAbstractEventDispatcher
}

func (ptr *QAbstractEventDispatcher) QAbstractEventDispatcher_PTR() *QAbstractEventDispatcher {
	return ptr
}

func (this *QAbstractEventDispatcher) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractEventDispatcher) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQAbstractEventDispatcherFromPointer(cthis unsafe.Pointer) *QAbstractEventDispatcher {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractEventDispatcher{bcthis0}
}
func (*QAbstractEventDispatcher) NewFromPointer(cthis unsafe.Pointer) *QAbstractEventDispatcher {
	return NewQAbstractEventDispatcherFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractEventDispatcher) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAbstractEventDispatcher10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractEventDispatcher(QObject *)

/*
Constructs a new event dispatcher with the given parent.
*/
func (*QAbstractEventDispatcher) NewForInherit(parent QObject_ITF /*777 QObject **/) *QAbstractEventDispatcher {
	return NewQAbstractEventDispatcher(parent)
}
func NewQAbstractEventDispatcher(parent QObject_ITF /*777 QObject **/) *QAbstractEventDispatcher {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcherC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractEventDispatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractEventDispatcher")
	return gothis
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractEventDispatcher(QObject *)

/*
Constructs a new event dispatcher with the given parent.
*/
func (*QAbstractEventDispatcher) NewForInherit__() *QAbstractEventDispatcher {
	return NewQAbstractEventDispatcher__()
}
func NewQAbstractEventDispatcher__() *QAbstractEventDispatcher {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcherC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractEventDispatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractEventDispatcher")
	return gothis
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractEventDispatcher()

/*

 */
func DeleteQAbstractEventDispatcher(this *QAbstractEventDispatcher) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcherD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:76
// index:0
// Public static Visibility=Default Availability=Available
// [8] QAbstractEventDispatcher * instance(QThread *)

/*
Returns a pointer to the event dispatcher object for the specified thread. If thread is zero, the current thread is used. If no event dispatcher exists for the specified thread, this function returns 0.

Note: If Qt is built without thread support, the thread argument is ignored.
*/
func (this *QAbstractEventDispatcher) Instance(thread QThread_ITF /*777 QThread **/) *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/ {
	var convArg0 unsafe.Pointer
	if thread != nil && thread.QThread_PTR() != nil {
		convArg0 = thread.QThread_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher8instanceEP7QThread", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractEventDispatcherFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QAbstractEventDispatcher_Instance(thread QThread_ITF /*777 QThread **/) *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/ {
	var nilthis *QAbstractEventDispatcher
	rv := nilthis.Instance(thread)
	return rv
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:76
// index:0
// Public static Visibility=Default Availability=Available
// [8] QAbstractEventDispatcher * instance(QThread *)

/*
Returns a pointer to the event dispatcher object for the specified thread. If thread is zero, the current thread is used. If no event dispatcher exists for the specified thread, this function returns 0.

Note: If Qt is built without thread support, the thread argument is ignored.
*/
func (this *QAbstractEventDispatcher) Instance__() *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/ {
	// arg: 0, QThread *=Pointer, QThread=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher8instanceEP7QThread", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractEventDispatcherFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:78
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool processEvents(QEventLoop::ProcessEventsFlags)

/*
Processes pending events that match flags until there are no more events to process. Returns true if an event was processed; otherwise returns false.

This function is especially useful if you have a long running operation, and want to show its progress without allowing user input by using the QEventLoop::ExcludeUserInputEvents flag.

If the QEventLoop::WaitForMoreEvents flag is set in flags, the behavior of this function is as follows:


If events are available, this function returns after processing them.
If no events are available, this function will wait until more are available and return after processing newly available events.


If the QEventLoop::WaitForMoreEvents flag is not set in flags, and no events are available, this function will return immediately.

Note: This function does not process events continuously; it returns after all available events are processed.

See also hasPendingEvents().
*/
func (this *QAbstractEventDispatcher) ProcessEvents(flags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13processEventsE6QFlagsIN10QEventLoop17ProcessEventsFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:79
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool hasPendingEvents()

/*

 */
func (this *QAbstractEventDispatcher) HasPendingEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher16hasPendingEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:81
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void registerSocketNotifier(QSocketNotifier *)

/*
Registers notifier with the event loop. Subclasses must implement this method to tie a socket notifier into another event loop.
*/
func (this *QAbstractEventDispatcher) RegisterSocketNotifier(notifier QSocketNotifier_ITF /*777 QSocketNotifier **/) {
	var convArg0 unsafe.Pointer
	if notifier != nil && notifier.QSocketNotifier_PTR() != nil {
		convArg0 = notifier.QSocketNotifier_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher22registerSocketNotifierEP15QSocketNotifier", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void unregisterSocketNotifier(QSocketNotifier *)

/*
Unregisters notifier from the event dispatcher. Subclasses must reimplement this method to tie a socket notifier into another event loop. Reimplementations must call the base implementation.
*/
func (this *QAbstractEventDispatcher) UnregisterSocketNotifier(notifier QSocketNotifier_ITF /*777 QSocketNotifier **/) {
	var convArg0 unsafe.Pointer
	if notifier != nil && notifier.QSocketNotifier_PTR() != nil {
		convArg0 = notifier.QSocketNotifier_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher24unregisterSocketNotifierEP15QSocketNotifier", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] int registerTimer(int, Qt::TimerType, QObject *)

/*
Registers a timer with the specified interval and timerType for the given object and returns the timer id.
*/
func (this *QAbstractEventDispatcher) RegisterTimer(interval int, timerType int, object QObject_ITF /*777 QObject **/) int {
	var convArg2 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg2 = object.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13registerTimerEiN2Qt9TimerTypeEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), interval, timerType, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:91
// index:1
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void registerTimer(int, int, Qt::TimerType, QObject *)

/*
Registers a timer with the specified interval and timerType for the given object and returns the timer id.
*/
func (this *QAbstractEventDispatcher) RegisterTimer_1(timerId int, interval int, timerType int, object QObject_ITF /*777 QObject **/) {
	var convArg3 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg3 = object.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13registerTimerEiiN2Qt9TimerTypeEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timerId, interval, timerType, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:92
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool unregisterTimer(int)

/*
Unregisters the timer with the given timerId. Returns true if successful; otherwise returns false.

See also registerTimer() and unregisterTimers().
*/
func (this *QAbstractEventDispatcher) UnregisterTimer(timerId int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher15unregisterTimerEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timerId)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:93
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool unregisterTimers(QObject *)

/*
Unregisters all the timers associated with the given object. Returns true if all timers were successful removed; otherwise returns false.

See also unregisterTimer() and registeredTimers().
*/
func (this *QAbstractEventDispatcher) UnregisterTimers(object QObject_ITF /*777 QObject **/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher16unregisterTimersEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:96
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int remainingTime(int)

/*
Returns the remaining time in milliseconds with the given timerId. If the timer is inactive, the returned value will be -1. If the timer is overdue, the returned value will be 0.

See also Qt::TimerType.
*/
func (this *QAbstractEventDispatcher) RemainingTime(timerId int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13remainingTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timerId)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:103
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void wakeUp()

/*
Wakes up the event loop.

Note: This function is thread-safe.

See also awake().
*/
func (this *QAbstractEventDispatcher) WakeUp() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher6wakeUpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:104
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void interrupt()

/*
Interrupts event dispatching. The event dispatcher will return from processEvents() as soon as possible.
*/
func (this *QAbstractEventDispatcher) Interrupt() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher9interruptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:105
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void flush()

/*

 */
func (this *QAbstractEventDispatcher) Flush() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:107
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void startingUp()

/*

 */
func (this *QAbstractEventDispatcher) StartingUp() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher10startingUpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:108
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void closingDown()

/*

 */
func (this *QAbstractEventDispatcher) ClosingDown() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher11closingDownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installNativeEventFilter(QAbstractNativeEventFilter *)

/*
Installs an event filter filterObj for all native events received by the application.

The event filter filterObj receives events via its nativeEventFilter() function, which is called for all events received by all threads.

The nativeEventFilter() function should return true if the event should be filtered, (in this case, stopped). It should return false to allow normal Qt processing to continue: the native event can then be translated into a QEvent and handled by the standard Qt event filtering, e.g. QObject::installEventFilter().

If multiple event filters are installed, the filter that was installed last is activated first.

Note: The filter function set here receives native messages, that is, MSG or XEvent structs.

For maximum portability, you should always try to use QEvent objects and QObject::installEventFilter() whenever possible.

This function was introduced in  Qt 5.0.

See also QObject::installEventFilter().
*/
func (this *QAbstractEventDispatcher) InstallNativeEventFilter(filterObj QAbstractNativeEventFilter_ITF /*777 QAbstractNativeEventFilter **/) {
	var convArg0 unsafe.Pointer
	if filterObj != nil && filterObj.QAbstractNativeEventFilter_PTR() != nil {
		convArg0 = filterObj.QAbstractNativeEventFilter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher24installNativeEventFilterEP26QAbstractNativeEventFilter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeNativeEventFilter(QAbstractNativeEventFilter *)

/*
Removes the event filter filter from this object. The request is ignored if such an event filter has not been installed.

All event filters for this object are automatically removed when this object is destroyed.

It is always safe to remove an event filter, even during event filter filter activation (that is, even from within the nativeEventFilter() function).

This function was introduced in  Qt 5.0.

See also installNativeEventFilter() and QAbstractNativeEventFilter.
*/
func (this *QAbstractEventDispatcher) RemoveNativeEventFilter(filterObj QAbstractNativeEventFilter_ITF /*777 QAbstractNativeEventFilter **/) {
	var convArg0 unsafe.Pointer
	if filterObj != nil && filterObj.QAbstractNativeEventFilter_PTR() != nil {
		convArg0 = filterObj.QAbstractNativeEventFilter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher23removeNativeEventFilterEP26QAbstractNativeEventFilter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool filterNativeEvent(const QByteArray &, void *, long *)

/*
Sends message through the event filters that were set by installNativeEventFilter(). This function returns true as soon as an event filter returns true, and false otherwise to indicate that the processing of the event should continue.

Subclasses of QAbstractEventDispatcher must call this function for all messages received from the system to ensure compatibility with any extensions that may be used in the application. The type of event eventType is specific to the platform plugin chosen at run-time, and can be used to cast message to the right type. The result pointer is only used on Windows, and corresponds to the LRESULT pointer.

Note that the type of message is platform dependent. See QAbstractNativeEventFilter for details.

This function was introduced in  Qt 5.0.

See also installNativeEventFilter() and QAbstractNativeEventFilter::nativeEventFilter().
*/
func (this *QAbstractEventDispatcher) FilterNativeEvent(eventType QByteArray_ITF, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool {
	var convArg0 unsafe.Pointer
	if eventType != nil && eventType.QByteArray_PTR() != nil {
		convArg0 = eventType.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher17filterNativeEventERK10QByteArrayPvPl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, message, result)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void aboutToBlock()

/*
This signal is emitted before the event loop calls a function that could block.

See also awake().
*/
func (this *QAbstractEventDispatcher) AboutToBlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher12aboutToBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void awake()

/*
This signal is emitted after the event loop returns from a function that could block.

See also wakeUp() and aboutToBlock().
*/
func (this *QAbstractEventDispatcher) Awake() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher5awakeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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
