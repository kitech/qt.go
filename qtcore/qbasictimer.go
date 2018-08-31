package qtcore

// /usr/include/qt/QtCore/qbasictimer.h
// #include <qbasictimer.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QBasicTimer struct {
	*qtrt.CObject
}
type QBasicTimer_ITF interface {
	QBasicTimer_PTR() *QBasicTimer
}

func (ptr *QBasicTimer) QBasicTimer_PTR() *QBasicTimer { return ptr }

func (this *QBasicTimer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QBasicTimer) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQBasicTimerFromPointer(cthis unsafe.Pointer) *QBasicTimer {
	return &QBasicTimer{&qtrt.CObject{cthis}}
}
func (*QBasicTimer) NewFromPointer(cthis unsafe.Pointer) *QBasicTimer {
	return NewQBasicTimerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qbasictimer.h:55
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QBasicTimer()

/*
Contructs a basic timer.

See also start().
*/
func (*QBasicTimer) NewForInherit() *QBasicTimer {
	return NewQBasicTimer()
}
func NewQBasicTimer() *QBasicTimer {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicTimerC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBasicTimerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQBasicTimer)
	return gothis
}

// /usr/include/qt/QtCore/qbasictimer.h:56
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QBasicTimer()

/*

 */
func DeleteQBasicTimer(this *QBasicTimer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicTimerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 4)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qbasictimer.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isActive() const

/*
Returns true if the timer is running and has not been stopped; otherwise returns false.

See also start() and stop().
*/
func (this *QBasicTimer) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QBasicTimer8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbasictimer.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int timerId() const

/*
Returns the timer's ID.

See also QTimerEvent::timerId().
*/
func (this *QBasicTimer) TimerId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QBasicTimer7timerIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbasictimer.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(int, QObject *)

/*
Starts (or restarts) the timer with a msec milliseconds timeout. The timer will be a Qt::CoarseTimer. See Qt::TimerType for information on the different timer types.

The given object will receive timer events.

See also stop(), isActive(), QObject::timerEvent(), and Qt::CoarseTimer.
*/
func (this *QBasicTimer) Start(msec int, obj QObject_ITF /*777 QObject **/) {
	var convArg1 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg1 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicTimer5startEiP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbasictimer.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void start(int, Qt::TimerType, QObject *)

/*
Starts (or restarts) the timer with a msec milliseconds timeout. The timer will be a Qt::CoarseTimer. See Qt::TimerType for information on the different timer types.

The given object will receive timer events.

See also stop(), isActive(), QObject::timerEvent(), and Qt::CoarseTimer.
*/
func (this *QBasicTimer) Start_1(msec int, timerType int, obj QObject_ITF /*777 QObject **/) {
	var convArg2 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg2 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicTimer5startEiN2Qt9TimerTypeEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec, timerType, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbasictimer.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Stops the timer.

See also start() and isActive().
*/
func (this *QBasicTimer) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QBasicTimer4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

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
