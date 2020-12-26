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
type QMoveEvent struct {
	*qtcore.QEvent
}
type QMoveEvent_ITF interface {
	qtcore.QEvent_ITF
	QMoveEvent_PTR() *QMoveEvent
}

func (ptr *QMoveEvent) QMoveEvent_PTR() *QMoveEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QMoveEventFromptr(cthis Voidptr) *QMoveEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QMoveEvent{bcthis0}
}
func (*QMoveEvent) Fromptr(cthis Voidptr) *QMoveEvent {
	return QMoveEventFromptr(cthis)
}

func DeleteQMoveEvent(this *QMoveEvent) {
	rv, err := qtrt.Qtcc1(0, "_ZN10QMoveEventD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10079() {
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
