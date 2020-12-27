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
type QHideEvent struct {
	*qtcore.QEvent
}
type QHideEvent_ITF interface {
	qtcore.QEvent_ITF
	QHideEvent_PTR() *QHideEvent
}

func (ptr *QHideEvent) QHideEvent_PTR() *QHideEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QHideEventFromptr(cthis Voidptr) *QHideEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QHideEvent{bcthis0}
}
func (*QHideEvent) Fromptr(cthis Voidptr) *QHideEvent {
	return QHideEventFromptr(cthis)
}

func DeleteQHideEvent(this *QHideEvent) {
	rv, err := qtrt.Qtcc3(4019755370, "_ZN10QHideEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10101() {
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
