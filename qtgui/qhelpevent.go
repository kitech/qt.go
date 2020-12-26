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
type QHelpEvent struct {
	*qtcore.QEvent
}
type QHelpEvent_ITF interface {
	qtcore.QEvent_ITF
	QHelpEvent_PTR() *QHelpEvent
}

func (ptr *QHelpEvent) QHelpEvent_PTR() *QHelpEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QHelpEventFromptr(cthis Voidptr) *QHelpEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QHelpEvent{bcthis0}
}
func (*QHelpEvent) Fromptr(cthis Voidptr) *QHelpEvent {
	return QHelpEventFromptr(cthis)
}

func DeleteQHelpEvent(this *QHelpEvent) {
	rv, err := qtrt.Qtcc3(2446753538, "_ZN10QHelpEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10109() {
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
