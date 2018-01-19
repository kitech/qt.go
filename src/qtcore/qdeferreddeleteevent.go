//  header block begin
// /usr/include/qt/QtCore/qcoreevent.h
// #include <qcoreevent.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QDeferredDeleteEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qcoreevent.h:377
// index:0
// void QDeferredDeleteEvent()
func NewQDeferredDeleteEvent() *QDeferredDeleteEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QDeferredDeleteEventC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QDeferredDeleteEvent{cthis}
}

// /usr/include/qt/QtCore/qcoreevent.h:378
// index:0
// virtual
// void ~QDeferredDeleteEvent()
func DeleteQDeferredDeleteEvent(*QDeferredDeleteEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QDeferredDeleteEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:379
// index:0
// inline
// int loopLevel()
func (this *QDeferredDeleteEvent) LoopLevel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QDeferredDeleteEvent9loopLevelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
