//  header block begin
// /usr/include/qt/QtCore/qabstracteventdispatcher.h
// #include <qabstracteventdispatcher.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
type QAbstractEventDispatcher struct {
	*QObject
}

func (this *QAbstractEventDispatcher) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAbstractEventDispatcher) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAbstractEventDispatcher10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:73
// index:0
// void QAbstractEventDispatcher(class QObject *)
func NewQAbstractEventDispatcher(parent unsafe.Pointer) *QAbstractEventDispatcher {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcherC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractEventDispatcherFromPointer(cthis)
	return gothis
}
func NewQAbstractEventDispatcherFromPointer(cthis unsafe.Pointer) *QAbstractEventDispatcher {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractEventDispatcher{bcthis0}
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:123
// index:1
// void QAbstractEventDispatcher(class QAbstractEventDispatcherPrivate &, class QObject *)
func NewQAbstractEventDispatcher_1(arg0 unsafe.Pointer, parent unsafe.Pointer) *QAbstractEventDispatcher {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcherC1ER31QAbstractEventDispatcherPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, arg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractEventDispatcherFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:74
// index:0
// virtual
// void ~QAbstractEventDispatcher()
func DeleteQAbstractEventDispatcher(*QAbstractEventDispatcher) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcherD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:76
// index:0
// static
// QAbstractEventDispatcher * instance(class QThread *)
func (this *QAbstractEventDispatcher) Instance(thread unsafe.Pointer) {
	// 0: (thread QThread *), (thread)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher8instanceEP7QThread", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAbstractEventDispatcher_Instance(thread unsafe.Pointer) {
	// 0: (thread QThread *), (thread)
	var nilthis *QAbstractEventDispatcher
	nilthis.Instance(thread)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:78
// index:0
// pure virtual
// bool processEvents(class QEventLoop::ProcessEventsFlags)
func (this *QAbstractEventDispatcher) ProcessEvents(flags int) {
	// 0: (, QFlags<QEventLoop::ProcessEventsFlag> flags), (&flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13processEventsE6QFlagsIN10QEventLoop17ProcessEventsFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:79
// index:0
// pure virtual
// bool hasPendingEvents()
func (this *QAbstractEventDispatcher) HasPendingEvents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher16hasPendingEventsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:81
// index:0
// pure virtual
// void registerSocketNotifier(class QSocketNotifier *)
func (this *QAbstractEventDispatcher) RegisterSocketNotifier(notifier unsafe.Pointer) {
	// 0: (, notifier QSocketNotifier *), (notifier)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher22registerSocketNotifierEP15QSocketNotifier", ffiqt.FFI_TYPE_VOID, this.GetCthis(), notifier)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:82
// index:0
// pure virtual
// void unregisterSocketNotifier(class QSocketNotifier *)
func (this *QAbstractEventDispatcher) UnregisterSocketNotifier(notifier unsafe.Pointer) {
	// 0: (, notifier QSocketNotifier *), (notifier)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher24unregisterSocketNotifierEP15QSocketNotifier", ffiqt.FFI_TYPE_VOID, this.GetCthis(), notifier)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:90
// index:0
// int registerTimer(int, Qt::TimerType, class QObject *)
func (this *QAbstractEventDispatcher) RegisterTimer(interval int, timerType int, object unsafe.Pointer) {
	// 0: (, interval int, timerType Qt::TimerType, object QObject *), (&interval, &timerType, object)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13registerTimerEiN2Qt9TimerTypeEP7QObject", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &interval, &timerType, object)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:91
// index:1
// pure virtual
// void registerTimer(int, int, Qt::TimerType, class QObject *)
func (this *QAbstractEventDispatcher) RegisterTimer_1(timerId int, interval int, timerType int, object unsafe.Pointer) {
	// 1: (, timerId int, interval int, timerType Qt::TimerType, object QObject *), (&timerId, &interval, &timerType, object)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13registerTimerEiiN2Qt9TimerTypeEP7QObject", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &timerId, &interval, &timerType, object)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:92
// index:0
// pure virtual
// bool unregisterTimer(int)
func (this *QAbstractEventDispatcher) UnregisterTimer(timerId int) {
	// 0: (, timerId int), (&timerId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher15unregisterTimerEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &timerId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:93
// index:0
// pure virtual
// bool unregisterTimers(class QObject *)
func (this *QAbstractEventDispatcher) UnregisterTimers(object unsafe.Pointer) {
	// 0: (, object QObject *), (object)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher16unregisterTimersEP7QObject", ffiqt.FFI_TYPE_VOID, this.GetCthis(), object)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:94
// index:0
// pure virtual
// QList<QAbstractEventDispatcher::TimerInfo> registeredTimers(class QObject *)
func (this *QAbstractEventDispatcher) RegisteredTimers(object unsafe.Pointer) {
	// 0: (, object QObject *), (object)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAbstractEventDispatcher16registeredTimersEP7QObject", ffiqt.FFI_TYPE_VOID, this.GetCthis(), object)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:96
// index:0
// pure virtual
// int remainingTime(int)
func (this *QAbstractEventDispatcher) RemainingTime(timerId int) {
	// 0: (, timerId int), (&timerId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13remainingTimeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &timerId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:103
// index:0
// pure virtual
// void wakeUp()
func (this *QAbstractEventDispatcher) WakeUp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher6wakeUpEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:104
// index:0
// pure virtual
// void interrupt()
func (this *QAbstractEventDispatcher) Interrupt() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher9interruptEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:105
// index:0
// pure virtual
// void flush()
func (this *QAbstractEventDispatcher) Flush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher5flushEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:107
// index:0
// virtual
// void startingUp()
func (this *QAbstractEventDispatcher) StartingUp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher10startingUpEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:108
// index:0
// virtual
// void closingDown()
func (this *QAbstractEventDispatcher) ClosingDown() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher11closingDownEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:110
// index:0
// void installNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QAbstractEventDispatcher) InstallNativeEventFilter(filterObj unsafe.Pointer) {
	// 0: (, filterObj QAbstractNativeEventFilter *), (filterObj)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher24installNativeEventFilterEP26QAbstractNativeEventFilter", ffiqt.FFI_TYPE_VOID, this.GetCthis(), filterObj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:111
// index:0
// void removeNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QAbstractEventDispatcher) RemoveNativeEventFilter(filterObj unsafe.Pointer) {
	// 0: (, filterObj QAbstractNativeEventFilter *), (filterObj)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher23removeNativeEventFilterEP26QAbstractNativeEventFilter", ffiqt.FFI_TYPE_VOID, this.GetCthis(), filterObj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:112
// index:0
// bool filterNativeEvent(const class QByteArray &, void *, long *)
func (this *QAbstractEventDispatcher) FilterNativeEvent(eventType unsafe.Pointer, message unsafe.Pointer, result unsafe.Pointer) {
	// 0: (, eventType const QByteArray &, message void *, result long *), (eventType, message, result)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher17filterNativeEventERK10QByteArrayPvPl", ffiqt.FFI_TYPE_VOID, this.GetCthis(), eventType, message, result)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:119
// index:0
// void aboutToBlock()
func (this *QAbstractEventDispatcher) AboutToBlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher12aboutToBlockEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:120
// index:0
// void awake()
func (this *QAbstractEventDispatcher) Awake() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher5awakeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
