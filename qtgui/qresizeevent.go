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
// size 40
type QResizeEvent struct {
	*qtcore.QEvent
}
type QResizeEvent_ITF interface {
	qtcore.QEvent_ITF
	QResizeEvent_PTR() *QResizeEvent
}

func (ptr *QResizeEvent) QResizeEvent_PTR() *QResizeEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QResizeEventFromptr(cthis Voidptr) *QResizeEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QResizeEvent{bcthis0}
}
func (*QResizeEvent) Fromptr(cthis Voidptr) *QResizeEvent {
	return QResizeEventFromptr(cthis)
}

func DeleteQResizeEvent(this *QResizeEvent) {
	rv, err := qtrt.Qtcc3(1460422629, "_ZN12QResizeEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10085() {
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
