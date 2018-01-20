//  header block begin
// /usr/include/qt/QtCore/qbasictimer.h
// #include <qbasictimer.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QBasicTimer struct {
	*qtrt.CObject
}

func (this *QBasicTimer) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qbasictimer.h:55
// index:0
// inline
// void QBasicTimer()
func NewQBasicTimer() *QBasicTimer {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicTimerC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQBasicTimerFromPointer(cthis)
	return gothis
}
func NewQBasicTimerFromPointer(cthis unsafe.Pointer) *QBasicTimer {
	return &QBasicTimer{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qbasictimer.h:56
// index:0
// inline
// void ~QBasicTimer()
func DeleteQBasicTimer(*QBasicTimer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicTimerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbasictimer.h:58
// index:0
// inline
// bool isActive()
func (this *QBasicTimer) IsActive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QBasicTimer8isActiveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbasictimer.h:59
// index:0
// inline
// int timerId()
func (this *QBasicTimer) TimerId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QBasicTimer7timerIdEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbasictimer.h:61
// index:0
// void start(int, class QObject *)
func (this *QBasicTimer) Start(msec int, obj unsafe.Pointer) {
	// 0: (, msec int, obj QObject *), (&msec, obj)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicTimer5startEiP7QObject", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msec, obj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbasictimer.h:62
// index:1
// void start(int, Qt::TimerType, class QObject *)
func (this *QBasicTimer) Start_1(msec int, timerType int, obj unsafe.Pointer) {
	// 1: (, msec int, timerType Qt::TimerType, obj QObject *), (&msec, &timerType, obj)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicTimer5startEiN2Qt9TimerTypeEP7QObject", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msec, &timerType, obj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbasictimer.h:63
// index:0
// void stop()
func (this *QBasicTimer) Stop() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QBasicTimer4stopEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
