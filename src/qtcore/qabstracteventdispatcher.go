package qtcore

// /usr/include/qt/QtCore/qabstracteventdispatcher.h
// #include <qabstracteventdispatcher.h>
// #include <QtCore>

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
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func NewQAbstractEventDispatcherFromPointer(cthis unsafe.Pointer) *QAbstractEventDispatcher {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractEventDispatcher{bcthis0}
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractEventDispatcher) MetaObject() *QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QAbstractEventDispatcher10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:73
// index:0
// Public
// void QAbstractEventDispatcher(class QObject *)
func NewQAbstractEventDispatcher(parent *QObject /*444 QObject **/) *QAbstractEventDispatcher {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcherC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractEventDispatcherFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:74
// index:0
// Public virtual
// void ~QAbstractEventDispatcher()
func DeleteQAbstractEventDispatcher(*QAbstractEventDispatcher) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcherD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:76
// index:0
// Public static
// QAbstractEventDispatcher * instance(class QThread *)
func (this *QAbstractEventDispatcher) Instance(thread *QThread /*444 QThread **/) *QAbstractEventDispatcher /*444 QAbstractEventDispatcher **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher8instanceEP7QThread", ffiqt.FFI_TYPE_POINTER, thread)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQAbstractEventDispatcherFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QAbstractEventDispatcher_Instance(thread *QThread /*444 QThread **/) *QAbstractEventDispatcher /*444 QAbstractEventDispatcher **/ {
	var nilthis *QAbstractEventDispatcher
	rv := nilthis.Instance(thread)
	return rv
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:79
// index:0
// Public pure virtual
// bool hasPendingEvents()
func (this *QAbstractEventDispatcher) HasPendingEvents() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher16hasPendingEventsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:81
// index:0
// Public pure virtual
// void registerSocketNotifier(class QSocketNotifier *)
func (this *QAbstractEventDispatcher) RegisterSocketNotifier(notifier *QSocketNotifier /*444 QSocketNotifier **/) {
	var convArg0 = notifier.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher22registerSocketNotifierEP15QSocketNotifier", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:82
// index:0
// Public pure virtual
// void unregisterSocketNotifier(class QSocketNotifier *)
func (this *QAbstractEventDispatcher) UnregisterSocketNotifier(notifier *QSocketNotifier /*444 QSocketNotifier **/) {
	var convArg0 = notifier.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher24unregisterSocketNotifierEP15QSocketNotifier", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:90
// index:0
// Public
// int registerTimer(int, Qt::TimerType, class QObject *)
func (this *QAbstractEventDispatcher) RegisterTimer(interval int, timerType int, object *QObject /*444 QObject **/) int {
	var convArg2 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13registerTimerEiN2Qt9TimerTypeEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &interval, &timerType, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:91
// index:1
// Public pure virtual
// void registerTimer(int, int, Qt::TimerType, class QObject *)
func (this *QAbstractEventDispatcher) RegisterTimer_1(timerId int, interval int, timerType int, object *QObject /*444 QObject **/) {
	var convArg3 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13registerTimerEiiN2Qt9TimerTypeEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &timerId, &interval, &timerType, convArg3)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:92
// index:0
// Public pure virtual
// bool unregisterTimer(int)
func (this *QAbstractEventDispatcher) UnregisterTimer(timerId int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher15unregisterTimerEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &timerId)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:93
// index:0
// Public pure virtual
// bool unregisterTimers(class QObject *)
func (this *QAbstractEventDispatcher) UnregisterTimers(object *QObject /*444 QObject **/) bool {
	var convArg0 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher16unregisterTimersEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:96
// index:0
// Public pure virtual
// int remainingTime(int)
func (this *QAbstractEventDispatcher) RemainingTime(timerId int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher13remainingTimeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &timerId)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:103
// index:0
// Public pure virtual
// void wakeUp()
func (this *QAbstractEventDispatcher) WakeUp() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher6wakeUpEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:104
// index:0
// Public pure virtual
// void interrupt()
func (this *QAbstractEventDispatcher) Interrupt() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher9interruptEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:105
// index:0
// Public pure virtual
// void flush()
func (this *QAbstractEventDispatcher) Flush() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher5flushEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:107
// index:0
// Public virtual
// void startingUp()
func (this *QAbstractEventDispatcher) StartingUp() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher10startingUpEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:108
// index:0
// Public virtual
// void closingDown()
func (this *QAbstractEventDispatcher) ClosingDown() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher11closingDownEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:110
// index:0
// Public
// void installNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QAbstractEventDispatcher) InstallNativeEventFilter(filterObj *QAbstractNativeEventFilter /*444 QAbstractNativeEventFilter **/) {
	var convArg0 = filterObj.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher24installNativeEventFilterEP26QAbstractNativeEventFilter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:111
// index:0
// Public
// void removeNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QAbstractEventDispatcher) RemoveNativeEventFilter(filterObj *QAbstractNativeEventFilter /*444 QAbstractNativeEventFilter **/) {
	var convArg0 = filterObj.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher23removeNativeEventFilterEP26QAbstractNativeEventFilter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:112
// index:0
// Public
// bool filterNativeEvent(const class QByteArray &, void *, long *)
func (this *QAbstractEventDispatcher) FilterNativeEvent(eventType *QByteArray, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool {
	var convArg0 = eventType.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher17filterNativeEventERK10QByteArrayPvPl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, message, result)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:119
// index:0
// Public
// void aboutToBlock()
func (this *QAbstractEventDispatcher) AboutToBlock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher12aboutToBlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstracteventdispatcher.h:120
// index:0
// Public
// void awake()
func (this *QAbstractEventDispatcher) Awake() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QAbstractEventDispatcher5awakeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
