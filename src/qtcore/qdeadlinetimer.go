//  header block begin
// /usr/include/qt/QtCore/qdeadlinetimer.h
// #include <qdeadlinetimer.h>
// #include <QtCore>
package qtcore

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:65
// index:0
// inline
// void QDeadlineTimer(Qt::TimerType)
func NewQDeadlineTimer(type_ int) *QDeadlineTimer {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimerC2EN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	return &QDeadlineTimer{cthis}
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:67
// index:1
// inline
// void QDeadlineTimer(enum QDeadlineTimer::ForeverConstant, Qt::TimerType)
func NewQDeadlineTimer_1(arg0 int, type_ int) *QDeadlineTimer {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimerC2ENS_15ForeverConstantEN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID, cthis, &arg0, &type_)
	gopp.ErrPrint(err, rv)
	return &QDeadlineTimer{cthis}
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:69
// index:2
// void QDeadlineTimer(qint64, Qt::TimerType)
func NewQDeadlineTimer_2(msecs int64, type_ int) *QDeadlineTimer {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimerC2ExN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID, cthis, &msecs, &type_)
	gopp.ErrPrint(err, rv)
	return &QDeadlineTimer{cthis}
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:71
// index:0
// inline
// void swap(class QDeadlineTimer &)
func (this *QDeadlineTimer) Swap(other unsafe.Pointer) {
	// 0: (, QDeadlineTimer & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:74
// index:0
// inline
// bool isForever()
func (this *QDeadlineTimer) IsForever() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer9isForeverEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:76
// index:0
// bool hasExpired()
func (this *QDeadlineTimer) HasExpired() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer10hasExpiredEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:78
// index:0
// inline
// Qt::TimerType timerType()
func (this *QDeadlineTimer) TimerType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer9timerTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:80
// index:0
// void setTimerType(Qt::TimerType)
func (this *QDeadlineTimer) SetTimerType(type_ int) {
	// 0: (, Qt::TimerType type), (&type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer12setTimerTypeEN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:82
// index:0
// qint64 remainingTime()
func (this *QDeadlineTimer) RemainingTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer13remainingTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:83
// index:0
// qint64 remainingTimeNSecs()
func (this *QDeadlineTimer) RemainingTimeNSecs() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer18remainingTimeNSecsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:84
// index:0
// void setRemainingTime(qint64, Qt::TimerType)
func (this *QDeadlineTimer) SetRemainingTime(msecs int64, type_ int) {
	// 0: (, qint64 msecs, Qt::TimerType type), (&msecs, &type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer16setRemainingTimeExN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &msecs, &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:85
// index:0
// void setPreciseRemainingTime(qint64, qint64, Qt::TimerType)
func (this *QDeadlineTimer) SetPreciseRemainingTime(secs int64, nsecs int64, type_ int) {
	// 0: (, qint64 secs, qint64 nsecs, Qt::TimerType type), (&secs, &nsecs, &type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer23setPreciseRemainingTimeExxN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &secs, &nsecs, &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:88
// index:0
// qint64 deadline()
func (this *QDeadlineTimer) Deadline() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer8deadlineEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:89
// index:0
// qint64 deadlineNSecs()
func (this *QDeadlineTimer) DeadlineNSecs() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer13deadlineNSecsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:90
// index:0
// void setDeadline(qint64, Qt::TimerType)
func (this *QDeadlineTimer) SetDeadline(msecs int64, timerType int) {
	// 0: (, qint64 msecs, Qt::TimerType timerType), (&msecs, &timerType)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer11setDeadlineExN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &msecs, &timerType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:91
// index:0
// void setPreciseDeadline(qint64, qint64, Qt::TimerType)
func (this *QDeadlineTimer) SetPreciseDeadline(secs int64, nsecs int64, type_ int) {
	// 0: (, qint64 secs, qint64 nsecs, Qt::TimerType type), (&secs, &nsecs, &type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer18setPreciseDeadlineExxN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &secs, &nsecs, &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:94
// index:0
// static
// QDeadlineTimer addNSecs(class QDeadlineTimer, qint64)
func (this *QDeadlineTimer) AddNSecs(dt unsafe.Pointer, nsecs int64) {
	// 0: (QDeadlineTimer dt, qint64 nsecs), (dt, nsecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer8addNSecsES_x", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDeadlineTimer_AddNSecs(dt unsafe.Pointer, nsecs int64) {
	// 0: (QDeadlineTimer dt, qint64 nsecs), (dt, nsecs)
	var nilthis *QDeadlineTimer
	nilthis.AddNSecs(dt, nsecs)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:95
// index:0
// static
// QDeadlineTimer current(Qt::TimerType)
func (this *QDeadlineTimer) Current(timerType int) {
	// 0: (Qt::TimerType timerType), (timerType)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDeadlineTimer7currentEN2Qt9TimerTypeE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDeadlineTimer_Current(timerType int) {
	// 0: (Qt::TimerType timerType), (timerType)
	var nilthis *QDeadlineTimer
	nilthis.Current(timerType)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:162
// index:0
// inline
// std::chrono::nanoseconds remainingTimeAsDuration()
func (this *QDeadlineTimer) RemainingTimeAsDuration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDeadlineTimer23remainingTimeAsDurationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
