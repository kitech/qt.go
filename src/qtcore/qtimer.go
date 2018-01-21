package qtcore

// /usr/include/qt/QtCore/qtimer.h
// #include <qtimer.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 33
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
	*QObject
}

func (this *QTimer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func NewQTimerFromPointer(cthis unsafe.Pointer) *QTimer {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QTimer{bcthis0}
}

// /usr/include/qt/QtCore/qtimer.h:59
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTimer) MetaObject() *QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qtimer.h:66
// index:0
// Public
// void QTimer(class QObject *)
func NewQTimer(parent *QObject /*444 QObject **/) *QTimer {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimerC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qtimer.h:67
// index:0
// Public virtual
// void ~QTimer()
func DeleteQTimer(*QTimer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:69
// index:0
// Public inline
// bool isActive()
func (this *QTimer) IsActive() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer8isActiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtimer.h:70
// index:0
// Public inline
// int timerId()
func (this *QTimer) TimerId() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer7timerIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimer.h:72
// index:0
// Public
// void setInterval(int)
func (this *QTimer) SetInterval(msec int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer11setIntervalEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:73
// index:0
// Public inline
// int interval()
func (this *QTimer) Interval() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer8intervalEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimer.h:75
// index:0
// Public
// int remainingTime()
func (this *QTimer) RemainingTime() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer13remainingTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qtimer.h:77
// index:0
// Public inline
// void setTimerType(Qt::TimerType)
func (this *QTimer) SetTimerType(atype int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer12setTimerTypeEN2Qt9TimerTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &atype)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:78
// index:0
// Public inline
// Qt::TimerType timerType()
func (this *QTimer) TimerType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer9timerTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qtimer.h:80
// index:0
// Public inline
// void setSingleShot(_Bool)
func (this *QTimer) SetSingleShot(singleShot bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer13setSingleShotEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &singleShot)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:81
// index:0
// Public inline
// bool isSingleShot()
func (this *QTimer) IsSingleShot() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer12isSingleShotEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtimer.h:83
// index:0
// Public static
// void singleShot(int, const class QObject *, const char *)
func (this *QTimer) SingleShot(msec int, receiver *QObject /*444 const QObject **/, member string) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer10singleShotEiPK7QObjectPKc", ffiqt.FFI_TYPE_POINTER, msec, receiver, member)
	gopp.ErrPrint(err, rv)
}
func QTimer_SingleShot(msec int, receiver *QObject /*444 const QObject **/, member string) {
	var nilthis *QTimer
	nilthis.SingleShot(msec, receiver, member)
}

// /usr/include/qt/QtCore/qtimer.h:84
// index:1
// Public static
// void singleShot(int, Qt::TimerType, const class QObject *, const char *)
func (this *QTimer) SingleShot_1(msec int, timerType int, receiver *QObject /*444 const QObject **/, member string) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer10singleShotEiN2Qt9TimerTypeEPK7QObjectPKc", ffiqt.FFI_TYPE_POINTER, msec, timerType, receiver, member)
	gopp.ErrPrint(err, rv)
}
func QTimer_SingleShot_1(msec int, timerType int, receiver *QObject /*444 const QObject **/, member string) {
	var nilthis *QTimer
	nilthis.SingleShot_1(msec, timerType, receiver, member)
}

// /usr/include/qt/QtCore/qtimer.h:158
// index:0
// Public
// void start(int)
func (this *QTimer) Start(msec int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer5startEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:160
// index:1
// Public
// void start()
func (this *QTimer) Start_1() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer5startEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:161
// index:0
// Public
// void stop()
func (this *QTimer) Stop() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer4stopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:173
// index:0
// Public inline
// std::chrono::milliseconds intervalAsDuration()
func (this *QTimer) IntervalAsDuration() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer18intervalAsDurationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qtimer.h:178
// index:0
// Public inline
// std::chrono::milliseconds remainingTimeAsDuration()
func (this *QTimer) RemainingTimeAsDuration() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QTimer23remainingTimeAsDurationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qtimer.h:200
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QTimer) TimerEvent(arg0 *QTimerEvent /*444 QTimerEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QTimer10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
