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
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

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
// [8] const QMetaObject * metaObject()
func (this *QAbstractEventDispatcher) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAbstractEventDispatcher10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractEventDispatcher(QObject *)
func NewQAbstractEventDispatcher(parent QObject_ITF /*777 QObject **/) *QAbstractEventDispatcher {
	var convArg0 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcherC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractEventDispatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractEventDispatcher()
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
func (this *QAbstractEventDispatcher) Instance(thread QThread_ITF /*777 QThread **/) *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/ {
	var convArg0 = thread.QThread_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher8instanceEP7QThread", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractEventDispatcherFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QAbstractEventDispatcher_Instance(thread QThread_ITF /*777 QThread **/) *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/ {
	var nilthis *QAbstractEventDispatcher
	rv := nilthis.Instance(thread)
	return rv
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:78
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool processEvents(QEventLoop::ProcessEventsFlags)
func (this *QAbstractEventDispatcher) ProcessEvents(flags int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13processEventsE6QFlagsIN10QEventLoop17ProcessEventsFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:79
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool hasPendingEvents()
func (this *QAbstractEventDispatcher) HasPendingEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher16hasPendingEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:81
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void registerSocketNotifier(QSocketNotifier *)
func (this *QAbstractEventDispatcher) RegisterSocketNotifier(notifier QSocketNotifier_ITF /*777 QSocketNotifier **/) {
	var convArg0 = notifier.QSocketNotifier_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher22registerSocketNotifierEP15QSocketNotifier", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void unregisterSocketNotifier(QSocketNotifier *)
func (this *QAbstractEventDispatcher) UnregisterSocketNotifier(notifier QSocketNotifier_ITF /*777 QSocketNotifier **/) {
	var convArg0 = notifier.QSocketNotifier_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher24unregisterSocketNotifierEP15QSocketNotifier", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] int registerTimer(int, Qt::TimerType, QObject *)
func (this *QAbstractEventDispatcher) RegisterTimer(interval int, timerType int, object QObject_ITF /*777 QObject **/) int {
	var convArg2 = object.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13registerTimerEiN2Qt9TimerTypeEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), interval, timerType, convArg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:91
// index:1
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void registerTimer(int, int, Qt::TimerType, QObject *)
func (this *QAbstractEventDispatcher) RegisterTimer_1(timerId int, interval int, timerType int, object QObject_ITF /*777 QObject **/) {
	var convArg3 = object.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13registerTimerEiiN2Qt9TimerTypeEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timerId, interval, timerType, convArg3)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:92
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool unregisterTimer(int)
func (this *QAbstractEventDispatcher) UnregisterTimer(timerId int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher15unregisterTimerEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timerId)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:93
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool unregisterTimers(QObject *)
func (this *QAbstractEventDispatcher) UnregisterTimers(object QObject_ITF /*777 QObject **/) bool {
	var convArg0 = object.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher16unregisterTimersEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:96
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int remainingTime(int)
func (this *QAbstractEventDispatcher) RemainingTime(timerId int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13remainingTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timerId)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:103
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void wakeUp()
func (this *QAbstractEventDispatcher) WakeUp() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher6wakeUpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:104
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void interrupt()
func (this *QAbstractEventDispatcher) Interrupt() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher9interruptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:105
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void flush()
func (this *QAbstractEventDispatcher) Flush() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher5flushEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:107
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void startingUp()
func (this *QAbstractEventDispatcher) StartingUp() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher10startingUpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:108
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void closingDown()
func (this *QAbstractEventDispatcher) ClosingDown() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher11closingDownEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installNativeEventFilter(QAbstractNativeEventFilter *)
func (this *QAbstractEventDispatcher) InstallNativeEventFilter(filterObj QAbstractNativeEventFilter_ITF /*777 QAbstractNativeEventFilter **/) {
	var convArg0 = filterObj.QAbstractNativeEventFilter_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher24installNativeEventFilterEP26QAbstractNativeEventFilter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeNativeEventFilter(QAbstractNativeEventFilter *)
func (this *QAbstractEventDispatcher) RemoveNativeEventFilter(filterObj QAbstractNativeEventFilter_ITF /*777 QAbstractNativeEventFilter **/) {
	var convArg0 = filterObj.QAbstractNativeEventFilter_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher23removeNativeEventFilterEP26QAbstractNativeEventFilter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool filterNativeEvent(const QByteArray &, void *, long *)
func (this *QAbstractEventDispatcher) FilterNativeEvent(eventType QByteArray_ITF, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool {
	var convArg0 = eventType.QByteArray_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher17filterNativeEventERK10QByteArrayPvPl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, message, &result)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void aboutToBlock()
func (this *QAbstractEventDispatcher) AboutToBlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher12aboutToBlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void awake()
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
		qtrt.KeepMe()
	}
}

//  keep block end
