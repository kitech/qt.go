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
type QStatusTipEvent struct {
	*qtcore.QEvent
}
type QStatusTipEvent_ITF interface {
	qtcore.QEvent_ITF
	QStatusTipEvent_PTR() *QStatusTipEvent
}

func (ptr *QStatusTipEvent) QStatusTipEvent_PTR() *QStatusTipEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QStatusTipEventFromptr(cthis Voidptr) *QStatusTipEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QStatusTipEvent{bcthis0}
}
func (*QStatusTipEvent) Fromptr(cthis Voidptr) *QStatusTipEvent {
	return QStatusTipEventFromptr(cthis)
}

func DeleteQStatusTipEvent(this *QStatusTipEvent) {
	rv, err := qtrt.Qtcc3(223208154, "_ZN15QStatusTipEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
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
