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
// extern C begin: 2
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
type QEnterEvent struct {
	*qtcore.QEvent
}
type QEnterEvent_ITF interface {
	qtcore.QEvent_ITF
	QEnterEvent_PTR() *QEnterEvent
}

func (ptr *QEnterEvent) QEnterEvent_PTR() *QEnterEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QEnterEventFromptr(cthis Voidptr) *QEnterEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QEnterEvent{bcthis0}
}
func (*QEnterEvent) Fromptr(cthis Voidptr) *QEnterEvent {
	return QEnterEventFromptr(cthis)
}

func DeleteQEnterEvent(this *QEnterEvent) {
	rv, err := qtrt.Qtcc3(727210895, "_ZN11QEnterEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10069() {
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
