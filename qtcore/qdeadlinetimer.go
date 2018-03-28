package qtcore

// /usr/include/qt/QtCore/qdeadlinetimer.h
// #include <qdeadlinetimer.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QDeadlineTimer struct {
	*qtrt.CObject
}
type QDeadlineTimer_ITF interface {
	QDeadlineTimer_PTR() *QDeadlineTimer
}

func (ptr *QDeadlineTimer) QDeadlineTimer_PTR() *QDeadlineTimer { return ptr }

func (this *QDeadlineTimer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDeadlineTimer) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQDeadlineTimerFromPointer(cthis unsafe.Pointer) *QDeadlineTimer {
	return &QDeadlineTimer{&qtrt.CObject{cthis}}
}
func (*QDeadlineTimer) NewFromPointer(cthis unsafe.Pointer) *QDeadlineTimer {
	return NewQDeadlineTimerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QDeadlineTimer(Qt::TimerType)

/*
Constructs an expired QDeadlineTimer object. For this object, remainingTime() will return 0.

The timer type timerType may be ignored, since the timer is already expired. Similarly, for optimization purposes, this function will not attempt to obtain the current time and will use a value known to be in the past. Therefore, deadline() may return an unexpected value and this object cannot be used in calculation of how long it is overdue. If that functionality is required, use QDeadlineTimer::current().

See also hasExpired(), remainingTime(), Qt::TimerType, and current().
*/
func NewQDeadlineTimer(type_ int) *QDeadlineTimer {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimerC2EN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDeadlineTimer)
	return gothis
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QDeadlineTimer(Qt::TimerType)

/*
Constructs an expired QDeadlineTimer object. For this object, remainingTime() will return 0.

The timer type timerType may be ignored, since the timer is already expired. Similarly, for optimization purposes, this function will not attempt to obtain the current time and will use a value known to be in the past. Therefore, deadline() may return an unexpected value and this object cannot be used in calculation of how long it is overdue. If that functionality is required, use QDeadlineTimer::current().

See also hasExpired(), remainingTime(), Qt::TimerType, and current().
*/
func NewQDeadlineTimer__() *QDeadlineTimer {
	// arg: 0, Qt::TimerType=Elaborated, Qt::TimerType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimerC2EN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDeadlineTimer)
	return gothis
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:67
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QDeadlineTimer(QDeadlineTimer::ForeverConstant, Qt::TimerType)

/*
Constructs an expired QDeadlineTimer object. For this object, remainingTime() will return 0.

The timer type timerType may be ignored, since the timer is already expired. Similarly, for optimization purposes, this function will not attempt to obtain the current time and will use a value known to be in the past. Therefore, deadline() may return an unexpected value and this object cannot be used in calculation of how long it is overdue. If that functionality is required, use QDeadlineTimer::current().

See also hasExpired(), remainingTime(), Qt::TimerType, and current().
*/
func NewQDeadlineTimer_1(arg0 int, type_ int) *QDeadlineTimer {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimerC2ENS_15ForeverConstantEN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, arg0, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDeadlineTimer)
	return gothis
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:67
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QDeadlineTimer(QDeadlineTimer::ForeverConstant, Qt::TimerType)

/*
Constructs an expired QDeadlineTimer object. For this object, remainingTime() will return 0.

The timer type timerType may be ignored, since the timer is already expired. Similarly, for optimization purposes, this function will not attempt to obtain the current time and will use a value known to be in the past. Therefore, deadline() may return an unexpected value and this object cannot be used in calculation of how long it is overdue. If that functionality is required, use QDeadlineTimer::current().

See also hasExpired(), remainingTime(), Qt::TimerType, and current().
*/
func NewQDeadlineTimer_1_(arg0 int) *QDeadlineTimer {
	// arg: 1, Qt::TimerType=Elaborated, Qt::TimerType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimerC2ENS_15ForeverConstantEN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, arg0, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDeadlineTimer)
	return gothis
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDeadlineTimer(qint64, Qt::TimerType)

/*
Constructs an expired QDeadlineTimer object. For this object, remainingTime() will return 0.

The timer type timerType may be ignored, since the timer is already expired. Similarly, for optimization purposes, this function will not attempt to obtain the current time and will use a value known to be in the past. Therefore, deadline() may return an unexpected value and this object cannot be used in calculation of how long it is overdue. If that functionality is required, use QDeadlineTimer::current().

See also hasExpired(), remainingTime(), Qt::TimerType, and current().
*/
func NewQDeadlineTimer_2(msecs int64, type_ int) *QDeadlineTimer {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimerC2ExN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, msecs, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDeadlineTimer)
	return gothis
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:69
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDeadlineTimer(qint64, Qt::TimerType)

/*
Constructs an expired QDeadlineTimer object. For this object, remainingTime() will return 0.

The timer type timerType may be ignored, since the timer is already expired. Similarly, for optimization purposes, this function will not attempt to obtain the current time and will use a value known to be in the past. Therefore, deadline() may return an unexpected value and this object cannot be used in calculation of how long it is overdue. If that functionality is required, use QDeadlineTimer::current().

See also hasExpired(), remainingTime(), Qt::TimerType, and current().
*/
func NewQDeadlineTimer_2_(msecs int64) *QDeadlineTimer {
	// arg: 1, Qt::TimerType=Elaborated, Qt::TimerType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimerC2ExN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, msecs, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDeadlineTimer)
	return gothis
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QDeadlineTimer &)

/*
Swaps this deadline timer with the other deadline timer.
*/
func (this *QDeadlineTimer) Swap(other QDeadlineTimer_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QDeadlineTimer_PTR() != nil {
		convArg0 = other.QDeadlineTimer_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isForever() const

/*
Returns true if this QDeadlineTimer object never expires, false otherwise. For timers that never expire, remainingTime() always returns -1 and deadline() returns the maximum value.

See also ForeverConstant, hasExpired(), and remainingTime().
*/
func (this *QDeadlineTimer) IsForever() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer9isForeverEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasExpired() const

/*
Returns true if this QDeadlineTimer object has expired, false if there remains time left. For objects that have expired, remainingTime() will return zero and deadline() will return a time point in the past.

QDeadlineTimer objects created with the ForeverConstant never expire and this function always returns false for them.

See also isForever() and remainingTime().
*/
func (this *QDeadlineTimer) HasExpired() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer10hasExpiredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::TimerType timerType() const

/*
Returns the timer type is active for this object.

See also setTimerType().
*/
func (this *QDeadlineTimer) TimerType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer9timerTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTimerType(Qt::TimerType)

/*
Changes the timer type for this object to timerType.

The behavior for each possible value of timerType is operating-system dependent. Qt::PreciseTimer will use the most precise timer that Qt can find, with resolution of 1 millisecond or better, whereas QDeadlineTimer will try to use a more coarse timer for Qt::CoarseTimer and Qt::VeryCoarseTimer.

See also timerType() and Qt::TimerType.
*/
func (this *QDeadlineTimer) SetTimerType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer12setTimerTypeEN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 remainingTime() const

/*
Returns the remaining time in this QDeadlineTimer object in milliseconds. If the timer has already expired, this function will return zero and it is not possible to obtain the amount of time overdue with this function (to do that, see deadline()). If the timer was set to never expire, this function returns -1.

This function is suitable for use in Qt APIs that take a millisecond timeout, such as the many QIODevice waitFor functions or the timed lock functions in QMutex, QWaitCondition, QSemaphore, or QReadWriteLock. For example:


  mutex.tryLock(deadline.remainingTime());



See also setRemainingTime(), remainingTimeNSecs(), isForever(), and hasExpired().
*/
func (this *QDeadlineTimer) RemainingTime() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer13remainingTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 remainingTimeNSecs() const

/*
Returns the remaining time in this QDeadlineTimer object in nanoseconds. If the timer has already expired, this function will return zero and it is not possible to obtain the amount of time overdue with this function. If the timer was set to never expire, this function returns -1.

See also remainingTime(), isForever(), and hasExpired().
*/
func (this *QDeadlineTimer) RemainingTimeNSecs() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer18remainingTimeNSecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRemainingTime(qint64, Qt::TimerType)

/*
Sets the remaining time for this QDeadlineTimer object to msecs milliseconds from now, if msecs has a positive value. If msecs is zero, this QDeadlineTimer object will be marked as expired, whereas a value of -1 will set it to never expire.

The timer type for this QDeadlineTimer object will be set to the specified timerType.

See also setPreciseRemainingTime(), hasExpired(), isForever(), and remainingTime().
*/
func (this *QDeadlineTimer) SetRemainingTime(msecs int64, type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer16setRemainingTimeExN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs, type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRemainingTime(qint64, Qt::TimerType)

/*
Sets the remaining time for this QDeadlineTimer object to msecs milliseconds from now, if msecs has a positive value. If msecs is zero, this QDeadlineTimer object will be marked as expired, whereas a value of -1 will set it to never expire.

The timer type for this QDeadlineTimer object will be set to the specified timerType.

See also setPreciseRemainingTime(), hasExpired(), isForever(), and remainingTime().
*/
func (this *QDeadlineTimer) SetRemainingTime__(msecs int64) {
	// arg: 1, Qt::TimerType=Elaborated, Qt::TimerType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer16setRemainingTimeExN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs, type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreciseRemainingTime(qint64, qint64, Qt::TimerType)

/*
Sets the remaining time for this QDeadlineTimer object to secs seconds plus nsecs nanoseconds from now, if secs has a positive value. If secs is -1, this QDeadlineTimer will be set it to never expire. If both parameters are zero, this QDeadlineTimer will be marked as expired.

The timer type for this QDeadlineTimer object will be set to the specified timerType.

See also setRemainingTime(), hasExpired(), isForever(), and remainingTime().
*/
func (this *QDeadlineTimer) SetPreciseRemainingTime(secs int64, nsecs int64, type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer23setPreciseRemainingTimeExxN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secs, nsecs, type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreciseRemainingTime(qint64, qint64, Qt::TimerType)

/*
Sets the remaining time for this QDeadlineTimer object to secs seconds plus nsecs nanoseconds from now, if secs has a positive value. If secs is -1, this QDeadlineTimer will be set it to never expire. If both parameters are zero, this QDeadlineTimer will be marked as expired.

The timer type for this QDeadlineTimer object will be set to the specified timerType.

See also setRemainingTime(), hasExpired(), isForever(), and remainingTime().
*/
func (this *QDeadlineTimer) SetPreciseRemainingTime__(secs int64) {
	// arg: 1, qint64=Typedef, qint64=Typedef, long long
	nsecs := int64(0)
	// arg: 2, Qt::TimerType=Elaborated, Qt::TimerType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer23setPreciseRemainingTimeExxN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secs, nsecs, type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreciseRemainingTime(qint64, qint64, Qt::TimerType)

/*
Sets the remaining time for this QDeadlineTimer object to secs seconds plus nsecs nanoseconds from now, if secs has a positive value. If secs is -1, this QDeadlineTimer will be set it to never expire. If both parameters are zero, this QDeadlineTimer will be marked as expired.

The timer type for this QDeadlineTimer object will be set to the specified timerType.

See also setRemainingTime(), hasExpired(), isForever(), and remainingTime().
*/
func (this *QDeadlineTimer) SetPreciseRemainingTime__1(secs int64, nsecs int64) {
	// arg: 2, Qt::TimerType=Elaborated, Qt::TimerType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer23setPreciseRemainingTimeExxN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secs, nsecs, type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 deadline() const

/*
Returns the absolute time point for the deadline stored in QDeadlineTimer object, calculated in milliseconds relative to the reference clock, the same as QElapsedTimer::msecsSinceReference(). The value will be in the past if this QDeadlineTimer has expired.

If this QDeadlineTimer never expires, this function returns std::numeric_limits<qint64>::max().

This function can be used to calculate the amount of time a timer is overdue, by subtracting QDeadlineTimer::current() or QElapsedTimer::msecsSinceReference(), as in the following example:


  qint64 realTimeLeft = deadline.deadline();
  if (realTimeLeft != (std::numeric_limits<qint64>::max)()) {
      realTimeLeft -= QDeadlineTimer::current().deadline();
      // or:
      //QElapsedTimer timer;
      //timer.start();
      //realTimeLeft -= timer.msecsSinceReference();
  }



Note: Timers that were created as expired have an indetermine time point in the past as their deadline, so the above calculation may not work.

See also remainingTime(), deadlineNSecs(), and setDeadline().
*/
func (this *QDeadlineTimer) Deadline() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer8deadlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 deadlineNSecs() const

/*
Returns the absolute time point for the deadline stored in QDeadlineTimer object, calculated in nanoseconds relative to the reference clock, the same as QElapsedTimer::msecsSinceReference(). The value will be in the past if this QDeadlineTimer has expired.

If this QDeadlineTimer never expires, this function returns std::numeric_limits<qint64>::max().

This function can be used to calculate the amount of time a timer is overdue, by subtracting QDeadlineTimer::current(), as in the following example:


  qint64 realTimeLeft = deadline.deadlineNSecs();
  if (realTimeLeft != std::numeric_limits<qint64>::max())
      realTimeLeft -= QDeadlineTimer::current().deadlineNSecs();



Note: Timers that were created as expired have an indetermine time point in the past as their deadline, so the above calculation may not work.

See also remainingTime() and deadlineNSecs().
*/
func (this *QDeadlineTimer) DeadlineNSecs() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer13deadlineNSecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDeadline(qint64, Qt::TimerType)

/*
Sets the deadline for this QDeadlineTimer object to be the msecs absolute time point, counted in milliseconds since the reference clock (the same as QElapsedTimer::msecsSinceReference()), and the timer type to timerType. If the value is in the past, this QDeadlineTimer will be marked as expired.

If msecs is std::numeric_limits<qint64>::max(), this QDeadlineTimer will be set to never expire.

See also setPreciseDeadline(), deadline(), deadlineNSecs(), and setRemainingTime().
*/
func (this *QDeadlineTimer) SetDeadline(msecs int64, timerType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer11setDeadlineExN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs, timerType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDeadline(qint64, Qt::TimerType)

/*
Sets the deadline for this QDeadlineTimer object to be the msecs absolute time point, counted in milliseconds since the reference clock (the same as QElapsedTimer::msecsSinceReference()), and the timer type to timerType. If the value is in the past, this QDeadlineTimer will be marked as expired.

If msecs is std::numeric_limits<qint64>::max(), this QDeadlineTimer will be set to never expire.

See also setPreciseDeadline(), deadline(), deadlineNSecs(), and setRemainingTime().
*/
func (this *QDeadlineTimer) SetDeadline__(msecs int64) {
	// arg: 1, Qt::TimerType=Elaborated, Qt::TimerType=Enum,
	timerType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer11setDeadlineExN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs, timerType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreciseDeadline(qint64, qint64, Qt::TimerType)

/*
Sets the deadline for this QDeadlineTimer object to be secs seconds and nsecs nanoseconds since the reference clock epoch (the same as QElapsedTimer::msecsSinceReference()), and the timer type to timerType. If the value is in the past, this QDeadlineTimer will be marked as expired.

If secs or nsecs is std::numeric_limits<qint64>::max(), this QDeadlineTimer will be set to never expire. If nsecs is more than 1 billion nanoseconds (1 second), then secs will be adjusted accordingly.

See also setDeadline(), deadline(), deadlineNSecs(), and setRemainingTime().
*/
func (this *QDeadlineTimer) SetPreciseDeadline(secs int64, nsecs int64, type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer18setPreciseDeadlineExxN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secs, nsecs, type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreciseDeadline(qint64, qint64, Qt::TimerType)

/*
Sets the deadline for this QDeadlineTimer object to be secs seconds and nsecs nanoseconds since the reference clock epoch (the same as QElapsedTimer::msecsSinceReference()), and the timer type to timerType. If the value is in the past, this QDeadlineTimer will be marked as expired.

If secs or nsecs is std::numeric_limits<qint64>::max(), this QDeadlineTimer will be set to never expire. If nsecs is more than 1 billion nanoseconds (1 second), then secs will be adjusted accordingly.

See also setDeadline(), deadline(), deadlineNSecs(), and setRemainingTime().
*/
func (this *QDeadlineTimer) SetPreciseDeadline__(secs int64) {
	// arg: 1, qint64=Typedef, qint64=Typedef, long long
	nsecs := int64(0)
	// arg: 2, Qt::TimerType=Elaborated, Qt::TimerType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer18setPreciseDeadlineExxN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secs, nsecs, type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreciseDeadline(qint64, qint64, Qt::TimerType)

/*
Sets the deadline for this QDeadlineTimer object to be secs seconds and nsecs nanoseconds since the reference clock epoch (the same as QElapsedTimer::msecsSinceReference()), and the timer type to timerType. If the value is in the past, this QDeadlineTimer will be marked as expired.

If secs or nsecs is std::numeric_limits<qint64>::max(), this QDeadlineTimer will be set to never expire. If nsecs is more than 1 billion nanoseconds (1 second), then secs will be adjusted accordingly.

See also setDeadline(), deadline(), deadlineNSecs(), and setRemainingTime().
*/
func (this *QDeadlineTimer) SetPreciseDeadline__1(secs int64, nsecs int64) {
	// arg: 2, Qt::TimerType=Elaborated, Qt::TimerType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer18setPreciseDeadlineExxN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secs, nsecs, type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [16] QDeadlineTimer addNSecs(QDeadlineTimer, qint64)

/*
Returns a QDeadlineTimer object whose deadline is extended from dt's deadline by nsecs nanoseconds. If dt was set to never expire, this function returns a QDeadlineTimer that will not expire either.

Note: if dt was created as expired, its deadline is indeterminate and adding an amount of time may or may not cause it to become unexpired.
*/
func (this *QDeadlineTimer) AddNSecs(dt QDeadlineTimer_ITF /*123*/, nsecs int64) *QDeadlineTimer /*123*/ {
	var convArg0 unsafe.Pointer
	if dt != nil && dt.QDeadlineTimer_PTR() != nil {
		convArg0 = dt.QDeadlineTimer_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer8addNSecsES_x", qtrt.FFI_TYPE_POINTER, convArg0, nsecs)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDeadlineTimer)
	return rv2
}
func QDeadlineTimer_AddNSecs(dt QDeadlineTimer_ITF /*123*/, nsecs int64) *QDeadlineTimer /*123*/ {
	var nilthis *QDeadlineTimer
	rv := nilthis.AddNSecs(dt, nsecs)
	return rv
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:95
// index:0
// Public static Visibility=Default Availability=Available
// [16] QDeadlineTimer current(Qt::TimerType)

/*
Returns a QDeadlineTimer that is expired but is guaranteed to contain the current time. Objects created by this function can participate in the calculation of how long a timer is overdue, using the deadline() function.

The QDeadlineTimer object will be constructed with the specified timerType.
*/
func (this *QDeadlineTimer) Current(timerType int) *QDeadlineTimer /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer7currentEN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, timerType)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDeadlineTimer)
	return rv2
}
func QDeadlineTimer_Current(timerType int) *QDeadlineTimer /*123*/ {
	var nilthis *QDeadlineTimer
	rv := nilthis.Current(timerType)
	return rv
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:95
// index:0
// Public static Visibility=Default Availability=Available
// [16] QDeadlineTimer current(Qt::TimerType)

/*
Returns a QDeadlineTimer that is expired but is guaranteed to contain the current time. Objects created by this function can participate in the calculation of how long a timer is overdue, using the deadline() function.

The QDeadlineTimer object will be constructed with the specified timerType.
*/
func (this *QDeadlineTimer) Current__() *QDeadlineTimer /*123*/ {
	// arg: 0, Qt::TimerType=Elaborated, Qt::TimerType=Enum,
	timerType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer7currentEN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, timerType)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDeadlineTimer)
	return rv2
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:118
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QDeadlineTimer & operator+=(qint64)

/*

 */
func (this *QDeadlineTimer) Operator_add_equal(msecs int64) *QDeadlineTimer {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimerpLEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDeadlineTimer)
	return rv2
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QDeadlineTimer & operator-=(qint64)

/*

 */
func (this *QDeadlineTimer) Operator_minus_equal(msecs int64) *QDeadlineTimer {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimermIEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDeadlineTimer)
	return rv2
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:162
// index:0
// Public inline Visibility=Default Availability=Available
// [8] std::chrono::nanoseconds remainingTimeAsDuration() const

/*
Returns the time remaining before the deadline.
*/
func (this *QDeadlineTimer) RemainingTimeAsDuration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer23remainingTimeAsDurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

func DeleteQDeadlineTimer(this *QDeadlineTimer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*

 */
type QDeadlineTimer__ForeverConstant = int

// Used when creating a QDeadlineTimer to indicate the deadline should not expire
const QDeadlineTimer__Forever QDeadlineTimer__ForeverConstant = 0

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
