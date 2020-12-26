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
// size 64
type QHoverEvent struct {
	*QInputEvent
}
type QHoverEvent_ITF interface {
	QInputEvent_ITF
	QHoverEvent_PTR() *QHoverEvent
}

func (ptr *QHoverEvent) QHoverEvent_PTR() *QHoverEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QHoverEventFromptr(cthis Voidptr) *QHoverEvent {
	bcthis0 := QInputEventFromptr(cthis)
	return &QHoverEvent{bcthis0}
}
func (*QHoverEvent) Fromptr(cthis Voidptr) *QHoverEvent {
	return QHoverEventFromptr(cthis)
}

func DeleteQHoverEvent(this *QHoverEvent) {
	rv, err := qtrt.Qtcc3(2903469262, "_ZN11QHoverEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10065() {
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
