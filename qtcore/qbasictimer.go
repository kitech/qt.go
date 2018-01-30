package qtcore

// /usr/include/qt/QtCore/qbasictimer.h
// #include <qbasictimer.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QBasicTimer struct {
	*qtrt.CObject
}

func (this *QBasicTimer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QBasicTimer) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
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
func NewQBasicTimer() *QBasicTimer {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicTimerC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQBasicTimerFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qbasictimer.h:56
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QBasicTimer()
func DeleteQBasicTimer(*QBasicTimer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicTimerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbasictimer.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isActive()
func (this *QBasicTimer) IsActive() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QBasicTimer8isActiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qbasictimer.h:59
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int timerId()
func (this *QBasicTimer) TimerId() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QBasicTimer7timerIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbasictimer.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(int, QObject *)
func (this *QBasicTimer) Start(msec int, obj *QObject /*777 QObject **/) {
	var convArg1 = obj.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicTimer5startEiP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msec, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbasictimer.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void start(int, Qt::TimerType, QObject *)
func (this *QBasicTimer) Start_1(msec int, timerType int, obj *QObject /*777 QObject **/) {
	var convArg2 = obj.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicTimer5startEiN2Qt9TimerTypeEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msec, timerType, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbasictimer.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()
func (this *QBasicTimer) Stop() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicTimer4stopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
