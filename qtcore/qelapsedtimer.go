package qtcore

// /usr/include/qt/QtCore/qelapsedtimer.h
// #include <qelapsedtimer.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 67
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
type QElapsedTimer struct {
	*qtrt.CObject
}
type QElapsedTimer_ITF interface {
	QElapsedTimer_PTR() *QElapsedTimer
}

func (ptr *QElapsedTimer) QElapsedTimer_PTR() *QElapsedTimer { return ptr }

func (this *QElapsedTimer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QElapsedTimer) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQElapsedTimerFromPointer(cthis unsafe.Pointer) *QElapsedTimer {
	return &QElapsedTimer{&qtrt.CObject{cthis}}
}
func (*QElapsedTimer) NewFromPointer(cthis unsafe.Pointer) *QElapsedTimer {
	return NewQElapsedTimerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QElapsedTimer()

/*
Constructs an invalid QElapsedTimer. A timer becomes valid once it has been started.

This function was introduced in  Qt 5.4.

See also isValid() and start().
*/
func NewQElapsedTimer() *QElapsedTimer {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QElapsedTimerC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQElapsedTimerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQElapsedTimer)
	return gothis
}

// /usr/include/qt/QtCore/qelapsedtimer.h:65
// index:0
// Public static Visibility=Default Availability=Available
// [4] QElapsedTimer::ClockType clockType()

/*
Returns the clock type that this QElapsedTimer implementation uses.

See also isMonotonic().
*/
func (this *QElapsedTimer) ClockType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QElapsedTimer9clockTypeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QElapsedTimer_ClockType() int {
	var nilthis *QElapsedTimer
	rv := nilthis.ClockType()
	return rv
}

// /usr/include/qt/QtCore/qelapsedtimer.h:66
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isMonotonic()

/*
Returns true if this is a monotonic clock, false otherwise. See the information on the different clock types to understand which ones are monotonic.

See also clockType() and QElapsedTimer::ClockType.
*/
func (this *QElapsedTimer) IsMonotonic() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QElapsedTimer11isMonotonicEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QElapsedTimer_IsMonotonic() bool {
	var nilthis *QElapsedTimer
	rv := nilthis.IsMonotonic()
	return rv
}

// /usr/include/qt/QtCore/qelapsedtimer.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start()

/*
Starts this timer. Once started, a timer value can be checked with elapsed() or msecsSinceReference().

Normally, a timer is started just before a lengthy operation, such as:


      QElapsedTimer timer;
      timer.start();

      slowOperation1();

      qDebug() << "The slow operation took" << timer.elapsed() << "milliseconds";



Also, starting a timer makes it valid again.

See also restart(), invalidate(), and elapsed().
*/
func (this *QElapsedTimer) Start() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QElapsedTimer5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 restart()

/*
Restarts the timer and returns the time elapsed since the previous start. This function is equivalent to obtaining the elapsed time with elapsed() and then starting the timer again with start(), but it does so in one single operation, avoiding the need to obtain the clock value twice.

Calling this function on a QElapsedTimer that is invalid results in undefined behavior.

The following example illustrates how to use this function to calibrate a parameter to a slow operation (for example, an iteration count) so that this operation takes at least 250 milliseconds:


      QElapsedTimer timer;

      int count = 1;
      timer.start();
      do {
          count *= 2;
          slowOperation2(count);
      } while (timer.restart() < 250);

      return count;



See also start(), invalidate(), elapsed(), and isValid().
*/
func (this *QElapsedTimer) Restart() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QElapsedTimer7restartEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qelapsedtimer.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invalidate()

/*
Marks this QElapsedTimer object as invalid.

An invalid object can be checked with isValid(). Calculations of timer elapsed since invalid data are undefined and will likely produce bizarre results.

See also isValid(), start(), and restart().
*/
func (this *QElapsedTimer) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QElapsedTimer10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qelapsedtimer.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns false if the timer has never been started or invalidated by a call to invalidate().

See also invalidate(), start(), and restart().
*/
func (this *QElapsedTimer) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QElapsedTimer7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qelapsedtimer.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 nsecsElapsed() const

/*
Returns the number of nanoseconds since this QElapsedTimer was last started.

Calling this function on a QElapsedTimer that is invalid results in undefined behavior.

On platforms that do not provide nanosecond resolution, the value returned will be the best estimate available.

This function was introduced in  Qt 4.8.

See also start(), restart(), hasExpired(), and invalidate().
*/
func (this *QElapsedTimer) NsecsElapsed() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QElapsedTimer12nsecsElapsedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qelapsedtimer.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 elapsed() const

/*
Returns the number of milliseconds since this QElapsedTimer was last started.

Calling this function on a QElapsedTimer that is invalid results in undefined behavior.

See also start(), restart(), hasExpired(), isValid(), and invalidate().
*/
func (this *QElapsedTimer) Elapsed() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QElapsedTimer7elapsedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qelapsedtimer.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasExpired(qint64) const

/*
Returns true if this QElapsedTimer has already expired by timeout milliseconds (that is, more than timeout milliseconds have elapsed). The value of timeout can be -1 to indicate that this timer does not expire, in which case this function will always return false.

See also elapsed() and QDeadlineTimer.
*/
func (this *QElapsedTimer) HasExpired(timeout int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QElapsedTimer10hasExpiredEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qelapsedtimer.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 msecsSinceReference() const

/*
Returns the number of milliseconds between last time this QElapsedTimer object was started and its reference clock's start.

This number is usually arbitrary for all clocks except the QElapsedTimer::SystemTime clock. For that clock type, this number is the number of milliseconds since January 1st, 1970 at 0:00 UTC (that is, it is the Unix time expressed in milliseconds).

On Linux, Windows and Apple platforms, this value is usually the time since the system boot, though it usually does not include the time the system has spent in sleep states.

See also clockType() and elapsed().
*/
func (this *QElapsedTimer) MsecsSinceReference() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QElapsedTimer19msecsSinceReferenceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qelapsedtimer.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 msecsTo(const QElapsedTimer &) const

/*
Returns the number of milliseconds between this QElapsedTimer and other. If other was started before this object, the returned value will be negative. If it was started later, the returned value will be positive.

The return value is undefined if this object or other were invalidated.

See also secsTo() and elapsed().
*/
func (this *QElapsedTimer) MsecsTo(other QElapsedTimer_ITF) int64 {
	var convArg0 unsafe.Pointer
	if other != nil && other.QElapsedTimer_PTR() != nil {
		convArg0 = other.QElapsedTimer_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QElapsedTimer7msecsToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qelapsedtimer.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 secsTo(const QElapsedTimer &) const

/*
Returns the number of seconds between this QElapsedTimer and other. If other was started before this object, the returned value will be negative. If it was started later, the returned value will be positive.

Calling this function on or with a QElapsedTimer that is invalid results in undefined behavior.

See also msecsTo() and elapsed().
*/
func (this *QElapsedTimer) SecsTo(other QElapsedTimer_ITF) int64 {
	var convArg0 unsafe.Pointer
	if other != nil && other.QElapsedTimer_PTR() != nil {
		convArg0 = other.QElapsedTimer_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QElapsedTimer6secsToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qelapsedtimer.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QElapsedTimer &) const

/*

 */
func (this *QElapsedTimer) Operator_equal_equal(other QElapsedTimer_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QElapsedTimer_PTR() != nil {
		convArg0 = other.QElapsedTimer_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QElapsedTimereqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qelapsedtimer.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QElapsedTimer &) const

/*

 */
func (this *QElapsedTimer) Operator_not_equal(other QElapsedTimer_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QElapsedTimer_PTR() != nil {
		convArg0 = other.QElapsedTimer_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QElapsedTimerneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQElapsedTimer(this *QElapsedTimer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QElapsedTimerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum contains the different clock types that QElapsedTimer may use.

QElapsedTimer will always use the same clock type in a particular machine, so this value will not change during the lifetime of a program. It is provided so that QElapsedTimer can be used with other non-Qt implementations, to guarantee that the same reference clock is being used.


*/
type QElapsedTimer__ClockType = int

// The human-readable system time. This clock is not monotonic.
const QElapsedTimer__SystemTime QElapsedTimer__ClockType = 0

// The system's monotonic clock, usually found in Unix systems. This clock is monotonic and does not overflow.
const QElapsedTimer__MonotonicClock QElapsedTimer__ClockType = 1

// The system's tick counter, used on Windows systems. This clock may overflow.
const QElapsedTimer__TickCounter QElapsedTimer__ClockType = 2

// The Mach kernel's absolute time (macOS and iOS). This clock is monotonic and does not overflow.
const QElapsedTimer__MachAbsoluteTime QElapsedTimer__ClockType = 3

// The high-resolution performance counter provided by Windows. This clock is monotonic and does not overflow.
const QElapsedTimer__PerformanceCounter QElapsedTimer__ClockType = 4

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
