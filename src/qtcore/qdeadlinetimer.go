package qtcore

// /usr/include/qt/QtCore/qdeadlinetimer.h
// #include <qdeadlinetimer.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QDeadlineTimer struct {
	*qtrt.CObject
}

func (this *QDeadlineTimer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQDeadlineTimerFromPointer(cthis unsafe.Pointer) *QDeadlineTimer {
	return &QDeadlineTimer{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:65
// index:0
// Public inline
// void QDeadlineTimer(Qt::TimerType)
func NewQDeadlineTimer(type_ int) *QDeadlineTimer {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimerC2EN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQDeadlineTimerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:67
// index:1
// Public inline
// void QDeadlineTimer(enum QDeadlineTimer::ForeverConstant, Qt::TimerType)
func NewQDeadlineTimer_1(arg0 int, type_ int) *QDeadlineTimer {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimerC2ENS_15ForeverConstantEN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID, cthis, &arg0, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQDeadlineTimerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:69
// index:2
// Public
// void QDeadlineTimer(qint64, Qt::TimerType)
func NewQDeadlineTimer_2(msecs int64, type_ int) *QDeadlineTimer {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimerC2ExN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID, cthis, &msecs, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQDeadlineTimerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:71
// index:0
// Public inline
// void swap(class QDeadlineTimer &)
func (this *QDeadlineTimer) Swap(other *QDeadlineTimer) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:74
// index:0
// Public inline
// bool isForever()
func (this *QDeadlineTimer) IsForever() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer9isForeverEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:76
// index:0
// Public
// bool hasExpired()
func (this *QDeadlineTimer) HasExpired() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer10hasExpiredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:78
// index:0
// Public inline
// Qt::TimerType timerType()
func (this *QDeadlineTimer) TimerType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer9timerTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:80
// index:0
// Public
// void setTimerType(Qt::TimerType)
func (this *QDeadlineTimer) SetTimerType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer12setTimerTypeEN2Qt9TimerTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:82
// index:0
// Public
// qint64 remainingTime()
func (this *QDeadlineTimer) RemainingTime() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer13remainingTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:83
// index:0
// Public
// qint64 remainingTimeNSecs()
func (this *QDeadlineTimer) RemainingTimeNSecs() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer18remainingTimeNSecsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:84
// index:0
// Public
// void setRemainingTime(qint64, Qt::TimerType)
func (this *QDeadlineTimer) SetRemainingTime(msecs int64, type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer16setRemainingTimeExN2Qt9TimerTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msecs, &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:85
// index:0
// Public
// void setPreciseRemainingTime(qint64, qint64, Qt::TimerType)
func (this *QDeadlineTimer) SetPreciseRemainingTime(secs int64, nsecs int64, type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer23setPreciseRemainingTimeExxN2Qt9TimerTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &secs, &nsecs, &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:88
// index:0
// Public
// qint64 deadline()
func (this *QDeadlineTimer) Deadline() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer8deadlineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:89
// index:0
// Public
// qint64 deadlineNSecs()
func (this *QDeadlineTimer) DeadlineNSecs() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer13deadlineNSecsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:90
// index:0
// Public
// void setDeadline(qint64, Qt::TimerType)
func (this *QDeadlineTimer) SetDeadline(msecs int64, timerType int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer11setDeadlineExN2Qt9TimerTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msecs, &timerType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:91
// index:0
// Public
// void setPreciseDeadline(qint64, qint64, Qt::TimerType)
func (this *QDeadlineTimer) SetPreciseDeadline(secs int64, nsecs int64, type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer18setPreciseDeadlineExxN2Qt9TimerTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &secs, &nsecs, &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:94
// index:0
// Public static
// QDeadlineTimer addNSecs(class QDeadlineTimer, qint64)
func (this *QDeadlineTimer) AddNSecs(dt *QDeadlineTimer /*123*/, nsecs int64) *QDeadlineTimer /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer8addNSecsES_x", ffiqt.FFI_TYPE_POINTER, dt, nsecs)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDeadlineTimer_AddNSecs(dt *QDeadlineTimer /*123*/, nsecs int64) *QDeadlineTimer /*123*/ {
	var nilthis *QDeadlineTimer
	rv := nilthis.AddNSecs(dt, nsecs)
	return rv
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:95
// index:0
// Public static
// QDeadlineTimer current(Qt::TimerType)
func (this *QDeadlineTimer) Current(timerType int) *QDeadlineTimer /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer7currentEN2Qt9TimerTypeE", ffiqt.FFI_TYPE_POINTER, timerType)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QDeadlineTimer_Current(timerType int) *QDeadlineTimer /*123*/ {
	var nilthis *QDeadlineTimer
	rv := nilthis.Current(timerType)
	return rv
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:162
// index:0
// Public inline
// std::chrono::nanoseconds remainingTimeAsDuration()
func (this *QDeadlineTimer) RemainingTimeAsDuration() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer23remainingTimeAsDurationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

//  body block end
