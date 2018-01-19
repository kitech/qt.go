//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QApplicationStateChangeEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:1052
// index:0
// void QApplicationStateChangeEvent(Qt::ApplicationState)
func NewQApplicationStateChangeEvent(state int) *QApplicationStateChangeEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN28QApplicationStateChangeEventC2EN2Qt16ApplicationStateE", ffiqt.FFI_TYPE_VOID, cthis, &state)
	gopp.ErrPrint(err, rv)
	return &QApplicationStateChangeEvent{cthis}
}

// /usr/include/qt/QtGui/qevent.h:1053
// index:0
// Qt::ApplicationState applicationState()
func (this *QApplicationStateChangeEvent) ApplicationState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK28QApplicationStateChangeEvent16applicationStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
