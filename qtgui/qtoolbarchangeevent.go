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
// extern C begin: 1
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
// size 24
type QToolBarChangeEvent struct {
	*qtcore.QEvent
}
type QToolBarChangeEvent_ITF interface {
	qtcore.QEvent_ITF
	QToolBarChangeEvent_PTR() *QToolBarChangeEvent
}

func (ptr *QToolBarChangeEvent) QToolBarChangeEvent_PTR() *QToolBarChangeEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QToolBarChangeEventFromptr(cthis Voidptr) *QToolBarChangeEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QToolBarChangeEvent{bcthis0}
}
func (*QToolBarChangeEvent) Fromptr(cthis Voidptr) *QToolBarChangeEvent {
	return QToolBarChangeEventFromptr(cthis)
}

func DeleteQToolBarChangeEvent(this *QToolBarChangeEvent) {
	rv, err := qtrt.Qtcc3(1334881248, "_ZN19QToolBarChangeEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10119() {
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
