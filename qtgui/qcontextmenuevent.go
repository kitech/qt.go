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
// size 56
type QContextMenuEvent struct {
	*QInputEvent
}
type QContextMenuEvent_ITF interface {
	QInputEvent_ITF
	QContextMenuEvent_PTR() *QContextMenuEvent
}

func (ptr *QContextMenuEvent) QContextMenuEvent_PTR() *QContextMenuEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QContextMenuEventFromptr(cthis Voidptr) *QContextMenuEvent {
	bcthis0 := QInputEventFromptr(cthis)
	return &QContextMenuEvent{bcthis0}
}
func (*QContextMenuEvent) Fromptr(cthis Voidptr) *QContextMenuEvent {
	return QContextMenuEventFromptr(cthis)
}

func DeleteQContextMenuEvent(this *QContextMenuEvent) {
	rv, err := qtrt.Qtcc3(3486689009, "_ZN17QContextMenuEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QContextMenuEvent__Reason = int

//
const QContextMenuEvent__Mouse QContextMenuEvent__Reason = 0

//
const QContextMenuEvent__Keyboard QContextMenuEvent__Reason = 1

//
const QContextMenuEvent__Other QContextMenuEvent__Reason = 2

func (this *QContextMenuEvent) ReasonItemName(val int) string {
	switch val {
	case QContextMenuEvent__Mouse: // 0
		return "Mouse"
	case QContextMenuEvent__Keyboard: // 1
		return "Keyboard"
	case QContextMenuEvent__Other: // 2
		return "Other"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QContextMenuEvent_ReasonItemName(val int) string {
	var nilthis *QContextMenuEvent
	return nilthis.ReasonItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10095() {
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
