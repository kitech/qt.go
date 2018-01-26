package qtcore

// /usr/include/qt/QtCore/qelapsedtimer.h
// #include <qelapsedtimer.h>
// #include <QtCore>

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
	*qtrt.CObject
}

func (this *QElapsedTimer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QElapsedTimer) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQElapsedTimerFromPointer(cthis unsafe.Pointer) *QElapsedTimer {
	return &QElapsedTimer{&qtrt.CObject{cthis}}
}
func (*QElapsedTimer) NewFromPointer(cthis unsafe.Pointer) *QElapsedTimer {
	return NewQElapsedTimerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:59
// index:0
// Public inline
// void QElapsedTimer()
func NewQElapsedTimer() *QElapsedTimer {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QElapsedTimerC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQElapsedTimerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qelapsedtimer.h:65
// index:0
// Public static
// QElapsedTimer::ClockType clockType()
func (this *QElapsedTimer) ClockType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QElapsedTimer9clockTypeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QElapsedTimer_ClockType() int {
	var nilthis *QElapsedTimer
	rv := nilthis.ClockType()
	return rv
}

// /usr/include/qt/QtCore/qelapsedtimer.h:66
// index:0
// Public static
// bool isMonotonic()
func (this *QElapsedTimer) IsMonotonic() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QElapsedTimer11isMonotonicEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QElapsedTimer_IsMonotonic() bool {
	var nilthis *QElapsedTimer
	rv := nilthis.IsMonotonic()
	return rv
}

// /usr/include/qt/QtCore/qelapsedtimer.h:68
// index:0
// Public
// void start()
func (this *QElapsedTimer) Start() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QElapsedTimer5startEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:69
// index:0
// Public
// qint64 restart()
func (this *QElapsedTimer) Restart() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QElapsedTimer7restartEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qelapsedtimer.h:70
// index:0
// Public
// void invalidate()
func (this *QElapsedTimer) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QElapsedTimer10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:71
// index:0
// Public
// bool isValid()
func (this *QElapsedTimer) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qelapsedtimer.h:73
// index:0
// Public
// qint64 nsecsElapsed()
func (this *QElapsedTimer) NsecsElapsed() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer12nsecsElapsedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qelapsedtimer.h:74
// index:0
// Public
// qint64 elapsed()
func (this *QElapsedTimer) Elapsed() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer7elapsedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qelapsedtimer.h:75
// index:0
// Public
// bool hasExpired(qint64)
func (this *QElapsedTimer) HasExpired(timeout int64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer10hasExpiredEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qelapsedtimer.h:77
// index:0
// Public
// qint64 msecsSinceReference()
func (this *QElapsedTimer) MsecsSinceReference() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer19msecsSinceReferenceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qelapsedtimer.h:78
// index:0
// Public
// qint64 msecsTo(const class QElapsedTimer &)
func (this *QElapsedTimer) MsecsTo(other *QElapsedTimer) int64 {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer7msecsToERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qelapsedtimer.h:79
// index:0
// Public
// qint64 secsTo(const class QElapsedTimer &)
func (this *QElapsedTimer) SecsTo(other *QElapsedTimer) int64 {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QElapsedTimer6secsToERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

type QElapsedTimer__ClockType = int

const QElapsedTimer__SystemTime QElapsedTimer__ClockType = 0
const QElapsedTimer__MonotonicClock QElapsedTimer__ClockType = 1
const QElapsedTimer__TickCounter QElapsedTimer__ClockType = 2
const QElapsedTimer__MachAbsoluteTime QElapsedTimer__ClockType = 3
const QElapsedTimer__PerformanceCounter QElapsedTimer__ClockType = 4

//  body block end
