//  header block begin
// /usr/include/qt/QtCore/qtimer.h
// #include <qtimer.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
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
type QTimer struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qtimer.h:59
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTimer) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:66
// index:0
// void QTimer(class QObject *)
func NewQTimer(parent unsafe.Pointer) *QTimer {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimerC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QTimer{cthis}
}

// /usr/include/qt/QtCore/qtimer.h:67
// index:0
// virtual
// void ~QTimer()
func DeleteQTimer(*QTimer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:69
// index:0
// inline
// bool isActive()
func (this *QTimer) IsActive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer8isActiveEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:70
// index:0
// inline
// int timerId()
func (this *QTimer) TimerId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer7timerIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:72
// index:0
// void setInterval(int)
func (this *QTimer) SetInterval(msec int) {
	// 0: (, int msec), (&msec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer11setIntervalEi", ffiqt.FFI_TYPE_VOID, this.cthis, &msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:73
// index:0
// inline
// int interval()
func (this *QTimer) Interval() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer8intervalEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:75
// index:0
// int remainingTime()
func (this *QTimer) RemainingTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer13remainingTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:77
// index:0
// inline
// void setTimerType(Qt::TimerType)
func (this *QTimer) SetTimerType(atype int) {
	// 0: (, Qt::TimerType atype), (&atype)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer12setTimerTypeEN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &atype)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:78
// index:0
// inline
// Qt::TimerType timerType()
func (this *QTimer) TimerType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer9timerTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:80
// index:0
// inline
// void setSingleShot(_Bool)
func (this *QTimer) SetSingleShot(singleShot bool) {
	// 0: (, bool singleShot), (&singleShot)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer13setSingleShotEb", ffiqt.FFI_TYPE_VOID, this.cthis, &singleShot)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:81
// index:0
// inline
// bool isSingleShot()
func (this *QTimer) IsSingleShot() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer12isSingleShotEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:83
// index:0
// static
// void singleShot(int, const class QObject *, const char *)
func (this *QTimer) SingleShot(msec int, receiver unsafe.Pointer, member unsafe.Pointer) {
	// 0: (int msec, const QObject * receiver, const char * member), (msec, receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer10singleShotEiPK7QObjectPKc", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimer_SingleShot(msec int, receiver unsafe.Pointer, member unsafe.Pointer) {
	// 0: (int msec, const QObject * receiver, const char * member), (msec, receiver, member)
	var nilthis *QTimer
	nilthis.SingleShot(msec, receiver, member)
}

// /usr/include/qt/QtCore/qtimer.h:84
// index:1
// static
// void singleShot(int, Qt::TimerType, const class QObject *, const char *)
func (this *QTimer) SingleShot_1(msec int, timerType int, receiver unsafe.Pointer, member unsafe.Pointer) {
	// 1: (int msec, Qt::TimerType timerType, const QObject * receiver, const char * member), (msec, timerType, receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer10singleShotEiN2Qt9TimerTypeEPK7QObjectPKc", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTimer_SingleShot_1(msec int, timerType int, receiver unsafe.Pointer, member unsafe.Pointer) {
	// 1: (int msec, Qt::TimerType timerType, const QObject * receiver, const char * member), (msec, timerType, receiver, member)
	var nilthis *QTimer
	nilthis.SingleShot_1(msec, timerType, receiver, member)
}

// /usr/include/qt/QtCore/qtimer.h:158
// index:0
// void start(int)
func (this *QTimer) Start(msec int) {
	// 0: (, int msec), (&msec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer5startEi", ffiqt.FFI_TYPE_VOID, this.cthis, &msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:160
// index:1
// void start()
func (this *QTimer) Start_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer5startEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:161
// index:0
// void stop()
func (this *QTimer) Stop() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer4stopEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:173
// index:0
// inline
// std::chrono::milliseconds intervalAsDuration()
func (this *QTimer) IntervalAsDuration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer18intervalAsDurationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:178
// index:0
// inline
// std::chrono::milliseconds remainingTimeAsDuration()
func (this *QTimer) RemainingTimeAsDuration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer23remainingTimeAsDurationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
