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
func NewQDeadlineTimer(type_ int) *QDeadlineTimer {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimerC2EN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDeadlineTimerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDeadlineTimer)
	return gothis
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:67
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QDeadlineTimer(enum QDeadlineTimer::ForeverConstant, Qt::TimerType)
func NewQDeadlineTimer_1(arg0 int, type_ int) *QDeadlineTimer {
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
func NewQDeadlineTimer_2(msecs int64, type_ int) *QDeadlineTimer {
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
// [1] bool isForever()
func (this *QDeadlineTimer) IsForever() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer9isForeverEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasExpired()
func (this *QDeadlineTimer) HasExpired() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer10hasExpiredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::TimerType timerType()
func (this *QDeadlineTimer) TimerType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer9timerTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTimerType(Qt::TimerType)
func (this *QDeadlineTimer) SetTimerType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer12setTimerTypeEN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 remainingTime()
func (this *QDeadlineTimer) RemainingTime() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer13remainingTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 remainingTimeNSecs()
func (this *QDeadlineTimer) RemainingTimeNSecs() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer18remainingTimeNSecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRemainingTime(qint64, Qt::TimerType)
func (this *QDeadlineTimer) SetRemainingTime(msecs int64, type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer16setRemainingTimeExN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs, type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreciseRemainingTime(qint64, qint64, Qt::TimerType)
func (this *QDeadlineTimer) SetPreciseRemainingTime(secs int64, nsecs int64, type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer23setPreciseRemainingTimeExxN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secs, nsecs, type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 deadline()
func (this *QDeadlineTimer) Deadline() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer8deadlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 deadlineNSecs()
func (this *QDeadlineTimer) DeadlineNSecs() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDeadlineTimer13deadlineNSecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDeadline(qint64, Qt::TimerType)
func (this *QDeadlineTimer) SetDeadline(msecs int64, timerType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer11setDeadlineExN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs, timerType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreciseDeadline(qint64, qint64, Qt::TimerType)
func (this *QDeadlineTimer) SetPreciseDeadline(secs int64, nsecs int64, type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDeadlineTimer18setPreciseDeadlineExxN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), secs, nsecs, type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdeadlinetimer.h:94
// index:0
// Public static Visibility=Default Availability=Available
// [16] QDeadlineTimer addNSecs(QDeadlineTimer, qint64)
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

// /usr/include/qt/QtCore/qdeadlinetimer.h:118
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QDeadlineTimer & operator+=(qint64)
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
// [8] std::chrono::nanoseconds remainingTimeAsDuration()
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

type QDeadlineTimer__ForeverConstant = int

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
