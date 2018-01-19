//  header block begin
// /usr/include/qt/QtCore/qelapsedtimer.h
// #include <qelapsedtimer.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 59
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
type QElapsedTimer struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qelapsedtimer.h:59
// index:0
// inline
// void QElapsedTimer()
func NewQElapsedTimer() *QElapsedTimer {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QElapsedTimerC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QElapsedTimer{cthis}
}

// /usr/include/qt/QtCore/qelapsedtimer.h:65
// index:0
// static
// QElapsedTimer::ClockType clockType()
func (this *QElapsedTimer) ClockType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QElapsedTimer9clockTypeEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QElapsedTimer_ClockType() {
	// 0: (), ()
	var nilthis *QElapsedTimer
	nilthis.ClockType()
}

// /usr/include/qt/QtCore/qelapsedtimer.h:66
// index:0
// static
// bool isMonotonic()
func (this *QElapsedTimer) IsMonotonic() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QElapsedTimer11isMonotonicEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QElapsedTimer_IsMonotonic() {
	// 0: (), ()
	var nilthis *QElapsedTimer
	nilthis.IsMonotonic()
}

// /usr/include/qt/QtCore/qelapsedtimer.h:68
// index:0
// void start()
func (this *QElapsedTimer) Start() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QElapsedTimer5startEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:69
// index:0
// qint64 restart()
func (this *QElapsedTimer) Restart() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QElapsedTimer7restartEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:70
// index:0
// void invalidate()
func (this *QElapsedTimer) Invalidate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QElapsedTimer10invalidateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:71
// index:0
// bool isValid()
func (this *QElapsedTimer) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:73
// index:0
// qint64 nsecsElapsed()
func (this *QElapsedTimer) NsecsElapsed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer12nsecsElapsedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:74
// index:0
// qint64 elapsed()
func (this *QElapsedTimer) Elapsed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer7elapsedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:75
// index:0
// bool hasExpired(qint64)
func (this *QElapsedTimer) HasExpired(timeout int64) {
	// 0: (, qint64 timeout), (&timeout)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer10hasExpiredEx", ffiqt.FFI_TYPE_VOID, this.cthis, &timeout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:77
// index:0
// qint64 msecsSinceReference()
func (this *QElapsedTimer) MsecsSinceReference() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer19msecsSinceReferenceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:78
// index:0
// qint64 msecsTo(const class QElapsedTimer &)
func (this *QElapsedTimer) MsecsTo(other unsafe.Pointer) {
	// 0: (, const QElapsedTimer & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer7msecsToERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:79
// index:0
// qint64 secsTo(const class QElapsedTimer &)
func (this *QElapsedTimer) SecsTo(other unsafe.Pointer) {
	// 0: (, const QElapsedTimer & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer6secsToERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

//  body block end
