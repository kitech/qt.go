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
type QShortcutEvent struct {
	*qtcore.QEvent
}
type QShortcutEvent_ITF interface {
	qtcore.QEvent_ITF
	QShortcutEvent_PTR() *QShortcutEvent
}

func (ptr *QShortcutEvent) QShortcutEvent_PTR() *QShortcutEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QShortcutEventFromptr(cthis Voidptr) *QShortcutEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QShortcutEvent{bcthis0}
}
func (*QShortcutEvent) Fromptr(cthis Voidptr) *QShortcutEvent {
	return QShortcutEventFromptr(cthis)
}

func DeleteQShortcutEvent(this *QShortcutEvent) {
	rv, err := qtrt.Qtcc3(4005859061, "_ZN14QShortcutEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10121() {
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
