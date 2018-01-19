//  header block begin
// /usr/include/qt/QtCore/qcoreevent.h
// #include <qcoreevent.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QTimerEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qcoreevent.h:340
// index:0
// void QTimerEvent(int)
func NewQTimerEvent(timerId int) *QTimerEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTimerEventC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &timerId)
	gopp.ErrPrint(err, rv)
	return &QTimerEvent{cthis}
}

// /usr/include/qt/QtCore/qcoreevent.h:341
// index:0
// virtual
// void ~QTimerEvent()
func DeleteQTimerEvent(*QTimerEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTimerEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:342
// index:0
// inline
// int timerId()
func (this *QTimerEvent) TimerId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTimerEvent7timerIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
