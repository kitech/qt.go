package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*
 */
// size 112
type QNativeGestureEvent struct {
	*QInputEvent
}
type QNativeGestureEvent_ITF interface {
	QInputEvent_ITF
	QNativeGestureEvent_PTR() *QNativeGestureEvent
}

func (ptr *QNativeGestureEvent) QNativeGestureEvent_PTR() *QNativeGestureEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QNativeGestureEventFromptr(cthis Voidptr) *QNativeGestureEvent {
	bcthis0 := QInputEventFromptr(cthis)
	return &QNativeGestureEvent{bcthis0}
}
func (*QNativeGestureEvent) Fromptr(cthis Voidptr) *QNativeGestureEvent {
	return QNativeGestureEventFromptr(cthis)
}

func DeleteQNativeGestureEvent(this *QNativeGestureEvent) {
	rv, err := qtrt.Qtcc1(0, "_ZN19QNativeGestureEventD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10071() {
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
