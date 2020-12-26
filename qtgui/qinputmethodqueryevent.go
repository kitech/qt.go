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
// size 32
type QInputMethodQueryEvent struct {
	*qtcore.QEvent
}
type QInputMethodQueryEvent_ITF interface {
	qtcore.QEvent_ITF
	QInputMethodQueryEvent_PTR() *QInputMethodQueryEvent
}

func (ptr *QInputMethodQueryEvent) QInputMethodQueryEvent_PTR() *QInputMethodQueryEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QInputMethodQueryEventFromptr(cthis Voidptr) *QInputMethodQueryEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QInputMethodQueryEvent{bcthis0}
}
func (*QInputMethodQueryEvent) Fromptr(cthis Voidptr) *QInputMethodQueryEvent {
	return QInputMethodQueryEventFromptr(cthis)
}

func DeleteQInputMethodQueryEvent(this *QInputMethodQueryEvent) {
	rv, err := qtrt.Qtcc1(0, "_ZN22QInputMethodQueryEventD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10099() {
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
