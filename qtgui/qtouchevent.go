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
// size 72
type QTouchEvent struct {
	*QInputEvent
}
type QTouchEvent_ITF interface {
	QInputEvent_ITF
	QTouchEvent_PTR() *QTouchEvent
}

func (ptr *QTouchEvent) QTouchEvent_PTR() *QTouchEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QTouchEventFromptr(cthis Voidptr) *QTouchEvent {
	bcthis0 := QInputEventFromptr(cthis)
	return &QTouchEvent{bcthis0}
}
func (*QTouchEvent) Fromptr(cthis Voidptr) *QTouchEvent {
	return QTouchEventFromptr(cthis)
}

func DeleteQTouchEvent(this *QTouchEvent) {
	rv, err := qtrt.Qtcc3(1324337181, "_ZN11QTouchEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10125() {
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
