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
// size 112
type QScrollPrepareEvent struct {
	*qtcore.QEvent
}
type QScrollPrepareEvent_ITF interface {
	qtcore.QEvent_ITF
	QScrollPrepareEvent_PTR() *QScrollPrepareEvent
}

func (ptr *QScrollPrepareEvent) QScrollPrepareEvent_PTR() *QScrollPrepareEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QScrollPrepareEventFromptr(cthis Voidptr) *QScrollPrepareEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QScrollPrepareEvent{bcthis0}
}
func (*QScrollPrepareEvent) Fromptr(cthis Voidptr) *QScrollPrepareEvent {
	return QScrollPrepareEventFromptr(cthis)
}

func DeleteQScrollPrepareEvent(this *QScrollPrepareEvent) {
	rv, err := qtrt.Qtcc1(0, "_ZN19QScrollPrepareEventD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10127() {
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
