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
// extern C begin: 0
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
// size 88
type QDragMoveEvent struct {
	*QDropEvent
}
type QDragMoveEvent_ITF interface {
	QDropEvent_ITF
	QDragMoveEvent_PTR() *QDragMoveEvent
}

func (ptr *QDragMoveEvent) QDragMoveEvent_PTR() *QDragMoveEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QDragMoveEventFromptr(cthis Voidptr) *QDragMoveEvent {
	bcthis0 := QDropEventFromptr(cthis)
	return &QDragMoveEvent{bcthis0}
}
func (*QDragMoveEvent) Fromptr(cthis Voidptr) *QDragMoveEvent {
	return QDragMoveEventFromptr(cthis)
}

//  body block end

//  keep block begin

func init_unused_10111() {
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
