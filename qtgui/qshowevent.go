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
// size 24
type QShowEvent struct {
	*qtcore.QEvent
}
type QShowEvent_ITF interface {
	qtcore.QEvent_ITF
	QShowEvent_PTR() *QShowEvent
}

func (ptr *QShowEvent) QShowEvent_PTR() *QShowEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QShowEventFromptr(cthis Voidptr) *QShowEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QShowEvent{bcthis0}
}
func (*QShowEvent) Fromptr(cthis Voidptr) *QShowEvent {
	return QShowEventFromptr(cthis)
}

func DeleteQShowEvent(this *QShowEvent) {
	rv, err := qtrt.Qtcc3(663296749, "_ZN10QShowEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
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
