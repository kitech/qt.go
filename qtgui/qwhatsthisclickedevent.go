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
type QWhatsThisClickedEvent struct {
	*qtcore.QEvent
}
type QWhatsThisClickedEvent_ITF interface {
	qtcore.QEvent_ITF
	QWhatsThisClickedEvent_PTR() *QWhatsThisClickedEvent
}

func (ptr *QWhatsThisClickedEvent) QWhatsThisClickedEvent_PTR() *QWhatsThisClickedEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QWhatsThisClickedEventFromptr(cthis Voidptr) *QWhatsThisClickedEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QWhatsThisClickedEvent{bcthis0}
}
func (*QWhatsThisClickedEvent) Fromptr(cthis Voidptr) *QWhatsThisClickedEvent {
	return QWhatsThisClickedEventFromptr(cthis)
}

//  body block end

//  keep block begin

func init_unused_10113() {
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
