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
// size 56
type QInputMethodEvent struct {
	*qtcore.QEvent
}
type QInputMethodEvent_ITF interface {
	qtcore.QEvent_ITF
	QInputMethodEvent_PTR() *QInputMethodEvent
}

func (ptr *QInputMethodEvent) QInputMethodEvent_PTR() *QInputMethodEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QInputMethodEventFromptr(cthis Voidptr) *QInputMethodEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QInputMethodEvent{bcthis0}
}
func (*QInputMethodEvent) Fromptr(cthis Voidptr) *QInputMethodEvent {
	return QInputMethodEventFromptr(cthis)
}

func DeleteQInputMethodEvent(this *QInputMethodEvent) {
	rv, err := qtrt.Qtcc3(2759769397, "_ZN17QInputMethodEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QInputMethodEvent__AttributeType = int

//
const QInputMethodEvent__TextFormat QInputMethodEvent__AttributeType = 0

//
const QInputMethodEvent__Cursor QInputMethodEvent__AttributeType = 1

//
const QInputMethodEvent__Language QInputMethodEvent__AttributeType = 2

//
const QInputMethodEvent__Ruby QInputMethodEvent__AttributeType = 3

//
const QInputMethodEvent__Selection QInputMethodEvent__AttributeType = 4

func (this *QInputMethodEvent) AttributeTypeItemName(val int) string {
	switch val {
	case QInputMethodEvent__TextFormat: // 0
		return "TextFormat"
	case QInputMethodEvent__Cursor: // 1
		return "Cursor"
	case QInputMethodEvent__Language: // 2
		return "Language"
	case QInputMethodEvent__Ruby: // 3
		return "Ruby"
	case QInputMethodEvent__Selection: // 4
		return "Selection"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QInputMethodEvent_AttributeTypeItemName(val int) string {
	var nilthis *QInputMethodEvent
	return nilthis.AttributeTypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10097() {
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
