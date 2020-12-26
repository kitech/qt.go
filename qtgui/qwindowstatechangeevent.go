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
type QWindowStateChangeEvent struct {
	*qtcore.QEvent
}
type QWindowStateChangeEvent_ITF interface {
	qtcore.QEvent_ITF
	QWindowStateChangeEvent_PTR() *QWindowStateChangeEvent
}

func (ptr *QWindowStateChangeEvent) QWindowStateChangeEvent_PTR() *QWindowStateChangeEvent {
	return ptr
}

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QWindowStateChangeEventFromptr(cthis Voidptr) *QWindowStateChangeEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QWindowStateChangeEvent{bcthis0}
}
func (*QWindowStateChangeEvent) Fromptr(cthis Voidptr) *QWindowStateChangeEvent {
	return QWindowStateChangeEventFromptr(cthis)
}

func DeleteQWindowStateChangeEvent(this *QWindowStateChangeEvent) {
	rv, err := qtrt.Qtcc3(3837143262, "_ZN23QWindowStateChangeEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10123() {
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
