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
import "gopp"
import "qt.go/qtrt"

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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
// void timerEvent(class QTimerEvent *)
func (this *QTimer) InheritTimerEvent(f func(arg0 *QTimerEvent /*777 QTimerEvent **/)) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

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
func (this *QTimer) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQTimerFromPointer(cthis unsafe.Pointer) *QTimer {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QTimer{bcthis0}
}
func (*QTimer) NewFromPointer(cthis unsafe.Pointer) *QTimer {
	return NewQTimerFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtimer.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTimer) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QTimer10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qtimer.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTimer(QObject *)
func NewQTimer(parent *QObject /*777 QObject **/) *QTimer {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QTimerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimerFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qtimer.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTimer()
func DeleteQTimer(this *QTimer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QTimerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtimer.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isActive()
func (this *QTimer) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QTimer8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtimer.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int timerId()
func (this *QTimer) TimerId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QTimer7timerIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimer.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInterval(int)
func (this *QTimer) SetInterval(msec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QTimer11setIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int interval()
func (this *QTimer) Interval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QTimer8intervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimer.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] int remainingTime()
func (this *QTimer) RemainingTime() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QTimer13remainingTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtimer.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTimerType(Qt::TimerType)
func (this *QTimer) SetTimerType(atype int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QTimer12setTimerTypeEN2Qt9TimerTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), atype)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::TimerType timerType()
func (this *QTimer) TimerType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QTimer9timerTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qtimer.h:80
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setSingleShot(_Bool)
func (this *QTimer) SetSingleShot(singleShot bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QTimer13setSingleShotEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), singleShot)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isSingleShot()
func (this *QTimer) IsSingleShot() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QTimer12isSingleShotEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qtimer.h:83
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void singleShot(int, const QObject *, const char *)
func (this *QTimer) SingleShot(msec int, receiver *QObject /*777 const QObject **/, member string) {
	var convArg1 = receiver.GetCthis()
	var convArg2 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QTimer10singleShotEiPK7QObjectPKc", qtrt.FFI_TYPE_POINTER, msec, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}
func QTimer_SingleShot(msec int, receiver *QObject /*777 const QObject **/, member string) {
	var nilthis *QTimer
	nilthis.SingleShot(msec, receiver, member)
}

// /usr/include/qt/QtCore/qtimer.h:84
// index:1
// Public static Visibility=Default Availability=Available
// [-2] void singleShot(int, Qt::TimerType, const QObject *, const char *)
func (this *QTimer) SingleShot_1(msec int, timerType int, receiver *QObject /*777 const QObject **/, member string) {
	var convArg2 = receiver.GetCthis()
	var convArg3 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg3)
	rv, err := qtrt.InvokeQtFunc6("_ZN6QTimer10singleShotEiN2Qt9TimerTypeEPK7QObjectPKc", qtrt.FFI_TYPE_POINTER, msec, timerType, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
}
func QTimer_SingleShot_1(msec int, timerType int, receiver *QObject /*777 const QObject **/, member string) {
	var nilthis *QTimer
	nilthis.SingleShot_1(msec, timerType, receiver, member)
}

// /usr/include/qt/QtCore/qtimer.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(int)
func (this *QTimer) Start(msec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QTimer5startEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:160
// index:1
// Public Visibility=Default Availability=Available
// [-2] void start()
func (this *QTimer) Start_1() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QTimer5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()
func (this *QTimer) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QTimer4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtimer.h:173
// index:0
// Public inline Visibility=Default Availability=Available
// [8] std::chrono::milliseconds intervalAsDuration()
func (this *QTimer) IntervalAsDuration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QTimer18intervalAsDurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qtimer.h:178
// index:0
// Public inline Visibility=Default Availability=Available
// [8] std::chrono::milliseconds remainingTimeAsDuration()
func (this *QTimer) RemainingTimeAsDuration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QTimer23remainingTimeAsDurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qtimer.h:200
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QTimer) TimerEvent(arg0 *QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN6QTimer10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
