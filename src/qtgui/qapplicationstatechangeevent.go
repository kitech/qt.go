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
	*qtcore.QEvent
}

func (this *QApplicationStateChangeEvent) GetCthis() unsafe.Pointer {
	return this.QEvent.GetCthis()
}
func NewQApplicationStateChangeEventFromPointer(cthis unsafe.Pointer) *QApplicationStateChangeEvent {
	bcthis0 := qtcore.NewQEventFromPointer(cthis)
	return &QApplicationStateChangeEvent{bcthis0}
}

// /usr/include/qt/QtGui/qevent.h:1052
// index:0
// Public
// void QApplicationStateChangeEvent(Qt::ApplicationState)
func NewQApplicationStateChangeEvent(state int) *QApplicationStateChangeEvent {
	cthis := qtrt.Calloc(1, 256) // 24
	rv, err := ffiqt.InvokeQtFunc6("_ZN28QApplicationStateChangeEventC2EN2Qt16ApplicationStateE", ffiqt.FFI_TYPE_VOID, cthis, &state)
	gopp.ErrPrint(err, rv)
	gothis := NewQApplicationStateChangeEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:1053
// index:0
// Public
// Qt::ApplicationState applicationState()
func (this *QApplicationStateChangeEvent) ApplicationState() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK28QApplicationStateChangeEvent16applicationStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
